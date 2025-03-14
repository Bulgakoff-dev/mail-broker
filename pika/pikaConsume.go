package pika

import (
	"encoding/json"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"mail-broker/config"
	"mail-broker/logger"
	"mail-broker/processing"
	"time"
)

func ConsumeRabbitMQ(rmqAddress string, client *redis.Client) error {
	cfg := config.GetConfig()
	logg := logger.GetLogger()

	conn, err := amqp091.Dial(rmqAddress)
	if err != nil {
		// Logging and returning error
		logg.Error().Msg("Failed to connect to RabbitMQ")
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		// Logging and returning error
		logg.Error().Msg("Failed to open a channel")
		return err
	}
	defer ch.Close()

	q, err := DeclareQueueRabbitMQ(ch, cfg.RMQQueue)
	if err != nil {
		logg.Error().Msg("Failed to declare a queue")
		return err
	}

	//uuid_consumer := uuid.New().String()

	// Poll until messages are availble
	for {
		msg, ok, err := ch.Get(q.Name, true)
		if err != nil {
			logg.Error().Msg("Failed to get message")
			return err
		}

		if !ok {
			// Logging
			time.Sleep(1500 * time.Millisecond)
			logg.Info().Msg("No new messages")
			continue
		}

		// Message unpack
		data := map[string]string{}
		err = json.Unmarshal(msg.Body, &data)
		if err != nil {
			// Logging and returning error
			logg.Error().Msg("Failed to unmarshal message")
			return err
		}

		// Message processing
		logg.Info().Msg(string(msg.Body))
		err = processing.ProcessMail(data, client)

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func DeclareQueueRabbitMQ(channel *amqp091.Channel, queueName string) (amqp091.Queue, error) {
	q, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil)

	if err != nil {
		return q, err
	}

	return q, nil
}
