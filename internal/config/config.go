package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

type Config struct {
	ServerAddress         string `env:"SERVER_ADDRESS"`
	PostgresUsername      string `env:"POSTGRES_USERNAME"`
	PostgresPassword      string `env:"POSTGRES_PASSWORD"`
	PostgresHost          string `env:"POSTGRES_HOST"`
	PostgresPort          string `env:"POSTGRES_PORT"`
	PostgresDatabase      string `env:"POSTGRES_DATABASE"`
	GrpcHost              string `env:"GRPC_HOST"`
	GrpcPort              string `env:"GRPC_PORT"`
	TokenExpirationPeriod string `env:"TOKEN_EXPIRATION_PERIOD"`
	LogLevel              string `env:"LOG_LEVEL"`
	RabbitUser            string `env:"RABBIT_USERNAME"`
	RabbitPassword        string `env:"RABBIT_PASSWORD"`
	RabbitHost            string `env:"RABBIT_HOST"`
	RabbitPort            int    `env:"RABBIT_PORT"`
	RabbitQueue           string `env:"RABBIT_QUEUE"`
	PrometheusPort        int    `env:"PROMETHEUS_PORT"`
	PrometheusHost        string `env:"PROMETHEUS_HOST"`
	TelegramBotToken      string `env:"TELEGRAM_BOT_TOKEN"`
}

var DbDsn string

func Parse(s string) (*Config, error) {
	c := &Config{}
	if err := cleanenv.ReadConfig(s, c); err != nil {
		return nil, err
	}

	setLogLevel(c.LogLevel)

	return c, nil
}

func setLogLevel(level string) {
	switch level {
	case "debug":
		slog.SetLogLoggerLevel(-4)
	case "info":
		slog.SetLogLoggerLevel(0)
	case "warn":
		slog.SetLogLoggerLevel(4)
	case "error":
		slog.SetLogLoggerLevel(8)
	default:
		slog.SetLogLoggerLevel(4)
	}
}

func InitDbDSN(c *Config) {
	DbDsn = fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		c.PostgresUsername, c.PostgresPassword, c.PostgresDatabase, c.PostgresHost, c.PostgresPort,
	)
}

func (c *Config) GetRabbitDSN() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%d/", c.RabbitUser, c.RabbitPassword, c.RabbitHost, c.RabbitPort,
	)
}

func (c *Config) GetPrometheusURL() string {
	return fmt.Sprintf(
		"%s:%d", c.PrometheusHost, c.PrometheusPort,
	)
}
