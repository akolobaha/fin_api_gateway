package telegramhandler

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/service"
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

				securityModel, isValidTicker := entities.ValidateTicker(ticker)

				if isValidTicker {
					userInputs[message.Chat.ID].Status = "waiting_for_indicator"
					userInputs[message.Chat.ID].Ticker = securityModel.Ticker
					text := fmt.Sprintf("Выберите коэффициент для цели по %s", ticker)
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), text).WithReplyMarkup(tu.InlineKeyboard(
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("price").WithCallbackData("price"),
							tu.InlineKeyboardButton("p / bv").WithCallbackData("pbv"),
						),
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("p / e").WithCallbackData("pe"),
							tu.InlineKeyboardButton("p / s").WithCallbackData("ps"),
						),
						tu.InlineKeyboardRow(
							tu.InlineKeyboardButton("🔙 Назад").WithCallbackData("back"),
						),
					)))
				} else {
					hintMessage := service.GetSecuritiesHint()
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), fmt.Sprintf("🚫Неверный тикер. Попробуйте: %s", hintMessage)))
				}
			} else if input.Status == "waiting_for_value" {
				// value -> _
				value, err := strconv.ParseFloat(message.Text, 32)

				if err == nil {
					// Здесь вы можете сохранить цель пользователя

					tgUser := entities.FirstOrCreateTgUser(message.From.ID, message.From.Username)

					curr := userInputs[message.Chat.ID]

					userTarget := entities.NewUserTarget(
						nil, &tgUser.ID, curr.Ticker, curr.Indicator, float32(value), "rsbu", "telegram")

					err = userTarget.Save(db.DB)
					if err != nil {
						fmt.Println(err)
					} else {
						successMsgText := fmt.Sprintf("Цель 📐 %s по 📈 %s успешно сохранена 💾", userTarget.ValuationRatio, userTarget.Ticker)
						_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), successMsgText))
					}

					//delete(userInputs, message.Chat.ID) // Очистка статуса
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
					_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "😢Неверное значение. \nВведите число с плавающей точкой. 🔢\n Пример: GAZP"))
				}
			}
		}

	})
}
