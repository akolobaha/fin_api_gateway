package main

import (
	"context"
	"fin_api_gateway/cmd/commands"
	"fin_api_gateway/internal/config"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const defaultEnvFilePath = ".env"

//const defaultEnvFilePath = "/usr/local/bin/.env"

//const defaultEnvFilePath = ".env"

func main() {
	cfg, err := config.Parse(defaultEnvFilePath)

	if err != nil {
		panic("Ошибка парсинга конфигов")
	}
	config.InitDbDSN(cfg)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		<-exit
		cancel()
	}()

	commands.RunHttp(ctx, cfg)

	if err != nil {
		fmt.Println(err)
		return
	}
}
