package main

import (
	"context"
	"fin_api_gateway/cmd/commands"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/log"
	"fin_api_gateway/internal/monitoring"
	"os"
	"os/signal"
	"syscall"
)

const defaultEnvFilePath = ".env"

func init() {
	monitoring.RegisterPrometheus()
}

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)

	if err != nil {
		panic("Ошибка парсинга конфигов")
	}
	config.InitDbDSN(cfg)

	ctx, cancel := context.WithCancel(context.Background())
	monitoring.RunPrometheusServer(cfg.GetPrometheusURL())
	exit := make(chan os.Signal, 1)
	go func() {
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	go commands.RunHttp(ctx, cfg)
	go commands.TelegramBotRun(ctx, cfg)
	go commands.RunGRPCServer(ctx, cfg)

	<-ctx.Done()
	log.Info("server is shutting down...")
}
