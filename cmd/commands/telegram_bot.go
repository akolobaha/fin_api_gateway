package commands

import (
	"context"
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/handler/telegramhandler"
	"fin_api_gateway/internal/log"
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

func TelegramBotRun(ctx context.Context, cfg *config.Config) error {
	bot, err := telego.NewBot(cfg.TelegramBotToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Error("Ошибка инициализации телеграм-бота: ", err)
		return err
	}

	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	// Start
	bh.HandleMessage(telegramhandler.StartButtonHandler, th.CommandEqual("start"))

	// Ловим нажатия кнопок кнопок Добавление цели, мои цели
	bh.HandleCallbackQuery(telegramhandler.TargetButtonsHandler, th.AnyCallbackQueryWithMessage())

	// Обработка сообщений с клавиатуры: тикеры, коэффиценты, значения
	bh.HandleMessage(telegramhandler.KeyboardInputHandler)

	// Обработка выбора коэффициента через кнопки
	bh.HandleCallbackQuery(telegramhandler.TargetCoefficentButtonHandler, th.AnyCallbackQueryWithMessage())

	go func() {
		bh.Start()
	}()

	<-ctx.Done()

	log.Info("Received shutdown signal, stopping the bot...")

	return nil
}

//func isValidTicker(ticker string) bool {
//	for _, validTicker := range validTickers {
//		if ticker == validTicker {
//			return true
//		}
//	}
//	return false
//}
