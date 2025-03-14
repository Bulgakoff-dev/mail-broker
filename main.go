package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"mail-broker/config"
	"mail-broker/logger"
	"mail-broker/pika"
	"time"
)

const rabbitMQURL = "amqp://admin:p1cky8eimer@149.126.169.135:5672/"

func main() {
	config.LoadConfig()
	cfg := config.GetConfig()
	logger.InitLogger("app.log")
	logg := logger.GetLogger()

	logg.Info().Msg("SMTP broker started...")
	logg.Info().Msg("Ctrl+C for finishing")

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       cfg.RedisDB,
	})

	queueInf := "Queue listening: " + cfg.RMQQueue
	logg.Info().Msg(queueInf)

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		logg.Err(err).Msg("Redis ping failed - application stops.")
		panic(err)
	}

	logg.Info().Msg("RabbitMQ consuming starting")

	go func() {
		for {
			err = pika.ConsumeRabbitMQ(rabbitMQURL, rdb)
			if err != nil {
				logg.Err(err).Msg("RabbitMQ consumer failed.")
				time.Sleep(5 * time.Second)
			}
		}
	}()

	select {}
}
