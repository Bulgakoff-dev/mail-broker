package pika

import (
	"context"
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"testing"
	"time"
)

func TestPublishToQueue(t *testing.T) {

	conn, err := amqp091.Dial("amqp://admin:p1cky8eimer@149.126.169.135:5672/")

	if err != nil {
		t.Fatal("Connect rabbitmq error: ", "Error", err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal("Channel rabbitmq error: ", "Error", err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	body := map[string]string{
		"email":            "lithiumrobot@gmail.com",
		"confirmationLink": "http://creepyHelloWorld.com/someData.burg",
	}

	bodyJson, _ := json.Marshal(body)
	err = ch.PublishWithContext(ctx,
		"",          // exchange
		"queryMail", // routing key
		false,       // mandatory
		false,       // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(bodyJson),
		})

	if err != nil {
		t.Error("Publish rabbitmq error: ", "Error", err.Error())
	}
}
