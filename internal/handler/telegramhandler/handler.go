package telegramhandler

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	"fin_api_gateway/internal/telegram"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type InputItem struct {
	Status    string
	Ticker    string
	Indicator string
	TgUser    *entities.TgUser
}

var userInputs = make(map[int64]*InputItem)

var coefficients = []string{"price", "p / bv", "p / e", "p / s"}

func handleDbMessageRequest(bot *telego.Bot, message telego.Message, handler func(conn *db.Connection)) {
	gDB, err := db.ConnectToDB()
	if err != nil {
		log.Error("failed to connect to db", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("failed to close db", err)
		}
	}()

	handler(gDB)
}

func handleDbQueryRequest(bot *telego.Bot, query telego.CallbackQuery, handler func(conn *db.Connection)) {
	gDB, err := db.ConnectToDB()
	if err != nil {
		log.Error("failed to connect to db", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("failed to close db", err)
		}
	}()

	handler(gDB)
}

func StartButtonHandler(bot *telego.Bot, message telego.Message) {
	if message.Text == "/start" {

		tgUser := FirstOrCreateTgUser(message.From.ID, message.From.Username)

		if userInputs[message.GetChat().ID] == nil {
			userInputs[message.GetChat().ID] = &InputItem{
				TgUser: tgUser,
			}
		}

		// Проверим существование пользователя и запомним, если такого не было

		_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), telegram.WelcomeText(message.From.FirstName)).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Добавить цель").WithCallbackData("add_target"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Мои цели").WithCallbackData("my_targets"),
			),
		)))
	}
}

func IsValidCoefficient(coefficient string) bool {
	for _, validCoefficient := range coefficients {
		if coefficient == validCoefficient {
			return true
		}
	}
	return false
}
