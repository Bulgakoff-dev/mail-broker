package config

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	// Log config
	LogFilePath string `mapstructure:"log_file_path"`

	// Redis config
	RedisAddr string `mapstructure:"redis_addr"`
	RedisPass string `mapstructure:"redis_pass"`
	RedisDB   int    `mapstructure:"redis_db"`

	// RabbitMQ config
	RMQAddr  string `mapstructure:"rmq_addr"`
	RMQPass  string `mapstructure:"rmq_pass"`
	RMQUser  string `mapstructure:"rmq_user"`
	RMQQueue string `mapstructure:"rmq_queue"`

	// SMTP config
	SMTPAddr string `mapstructure:"smtp_server"`
	SMTPPass string `mapstructure:"smtp_password"`
	SMTPUser string `mapstructure:"smtp_username"`
	SMTPPort string `mapstructure:"smtp_port"`
	SMTPMail string `mapstructure:"smtp_mail"`
}

var (
	instance *Config
	once     sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")

		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file, %s", err)
		}

		config := &Config{}

		if err := viper.Unmarshal(config); err != nil {
			log.Fatalf("Ошибка парсинга конфигурации: %v", err)
		}

		instance = config
	})

	return instance
}

func GetConfig() *Config {
	return instance
}
