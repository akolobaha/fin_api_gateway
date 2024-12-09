package commands

import (
	"fin_api_gateway/internal/telegram"
	"fmt"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
)

func TelegramBotRun() {
	// Get Bot token from environment variables
	botToken := ""

	// Note: Please keep in mind that default logger may expose sensitive information,
	// use in development only
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Call method getMe
	botUser, _ := bot.GetMe()
	fmt.Printf("Bot User: %+v\n", botUser)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	// Create bot handler and specify from where to get updates
	bh, _ := th.NewBotHandler(bot, updates)

	defer bh.Stop()
	defer bot.StopLongPolling()

	// Register new handler with match on command /start
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		// Зарегаем пользователя по telegram_id, если такого раньше не было

		// Send a message with inline keyboard
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(message.Chat.ID),
			telegram.WelcomeText(message.From.FirstName),
		).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("\U00002705 Зарегистрироваться!").WithCallbackData("go"),
			)),
		))
	}, th.CommandEqual("start"))

	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {

		// Будем ловить и парсить текст

	})

	// Register new handler for command "/"
	bh.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		// Send a message with available commands
		commandsList := "Available commands:\n/start - Start the bot\n/help - Show this help message\n/go - Trigger the go action"
		_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), commandsList))
	}, th.CommandEqual("/"))

	// Register new handler with match on the callback query
	// with data equal to go and non-nil message
	bh.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "GO GO GO"))

		// Answer callback query
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual("go"))

	bh.Start()
}
