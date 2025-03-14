package processing

import (
	"github.com/redis/go-redis/v9"
	"log"
	"mail-broker/config"
	"testing"
)

func TestProcessMailFull(t *testing.T) {
	cfg := config.LoadConfig()

	log.Println("Test ProcessMail started")

	testMessage := map[string]string{
		"email":            "lithiumrobot@gmail.com",
		"confirmationLink": "https://hello-world.com/some_data",
		"firstName":        "Jane",
		"lastName":         "Doe",
		"locale":           "RU",
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	err := ProcessMail(testMessage, rdb)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		log.Println("Test ProcessMail finished")
	}
}

func TestProcessMailWithoutLocale(t *testing.T) {
	cfg := config.LoadConfig()

	log.Println("Test ProcessMail started")

	testMessage := map[string]string{
		"email":            "lithiumrobot@gmail.com",
		"confirmationLink": "https://hello-world.com/some_data",
		"firstName":        "Jane",
		"lastName":         "Doe",
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	err := ProcessMail(testMessage, rdb)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		log.Println("Test ProcessMail finished")
	}
}

func TestProcessMailWithoutName(t *testing.T) {
	cfg := config.LoadConfig()

	log.Println("Test ProcessMail started")

	testMessage := map[string]string{
		"email":            "lithiumrobot@gmail.com",
		"confirmationLink": "https://hello-world.com/some_data",
		"locale":           "RU",
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	err := ProcessMail(testMessage, rdb)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		log.Println("Test ProcessMail finished")
	}
}
