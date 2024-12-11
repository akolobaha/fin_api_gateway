package telegramhandler

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"strconv"
)

func KeyboardInputHandler(bot *telego.Bot, message telego.Message) {
	handleDbMessageRequest(bot, message, func(db *db.Connection) {
		if input, exists := userInputs[message.Chat.ID]; exists {
			if input.Status == "waiting_for_ticker" {
				ticker := message.Text

				securityModel, isValidTicker := ValidateTicker(ticker)

				if isValidTicker {
					userInputs[message.Chat.ID].Status = "waiting_for_coefficient"
					userInputs[message.Chat.ID].Ticker = securityModel.Ticker
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Выберите коэффициент").WithReplyMarkup(tu.InlineKeyboard(
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton(coefficients[0]).WithCallbackData(coefficients[0]),
							tu.InlineKeyboardButton(coefficients[1]).WithCallbackData(coefficients[1]),
						),
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton(coefficients[2]).WithCallbackData(coefficients[2]),
							tu.InlineKeyboardButton(coefficients[3]).WithCallbackData(coefficients[3]),
						),
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("Назад").WithCallbackData("back"),
						),
					)))
				} else {
					hintMessage := GetSecuritiesHint()
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), fmt.Sprintf("Неверный тикер. Попробуйте: %s", hintMessage)))
				}
			} else if input.Status == "waiting_for_coefficient" {
				// Здесь мы обрабатываем выбор коэффициента через кнопки
				coefficient := message.Text

				if IsValidCoefficient(coefficient) {
					userInputs[message.Chat.ID].Status = "waiting_for_value"
					userInputs[message.Chat.ID].Indicator = coefficient
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Введите значение (число с плавающей точкой)"))
				} else {
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Неверный коэффициент. Выберите из предложенных."))
				}
			} else if input.Status == "waiting_for_value" {
				// value -> _
				value, err := strconv.ParseFloat(message.Text, 32)

				if err == nil {
					// Здесь вы можете сохранить цель пользователя

					FirstOrCreateTgUser(message.From.ID, message.From.Username)

					curr := userInputs[message.Chat.ID]

					userTarget := entities.NewUserTarget(
						nil, &curr.TgUser.ID, curr.Ticker, curr.Indicator, float32(value), "rsbu", "telegram")

					err = userTarget.Save(db.DB)
					if err != nil {
						fmt.Println(err)
					} else {
						_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Значение успешно добавлено"))
					}

					delete(userInputs, message.Chat.ID) // Очистка статуса
					// Возврат в главное меню
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Выберите действие").WithReplyMarkup(tu.InlineKeyboard(
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("Добавить цель").WithCallbackData("add_target"),
						),
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("Мои цели").WithCallbackData("my_targets"),
						),
					)))
				} else {
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Неверное значение. Введите число с плавающей точкой."))
				}
			}
		}
	})

}

func TargetButtonsHandler(bot *telego.Bot, query telego.CallbackQuery) {
	if userInputs[query.Message.GetChat().ID] == nil {
		userInputs[query.Message.GetChat().ID] = &InputItem{}
	}

	switch query.Data {
	case "add_target":
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Введите тикер"))
		userInputs[query.Message.GetChat().ID].Status = "waiting_for_ticker"
	case "my_targets":
		// Здесь вы можете добавить логику для отображения целей
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Пока нет добавленных целей."))
	case "back": // Кнопка назад
		delete(userInputs, query.Message.GetChat().ID)
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Выберите действие").WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Добавить цель").WithCallbackData("add_target"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Мои цели").WithCallbackData("my_targets"),
			),
		)))
	}

	// Обработка нажатий на кнопки
	_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Выбор сделан"))
}

func TargetCoefficentButtonHandler(bot *telego.Bot, query telego.CallbackQuery) {
	if input, exists := userInputs[query.Message.GetChat().ID]; exists && input.Status == "waiting_for_coefficient" {
		coefficient := query.Data
		if IsValidCoefficient(coefficient) {
			userInputs[query.Message.GetChat().ID].Status = "waiting_for_value"
			_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Введите значение (число с плавающей точкой)"))
		}
	}
}
