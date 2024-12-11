package telegramhandler

import (
	"fin_api_gateway/internal/service"
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func ButtonPressHandler(bot *telego.Bot, query telego.CallbackQuery) {
	chatId := query.Message.GetChat().ID
	if userInputs[chatId] == nil {
		userInputs[chatId] = &InputItem{}
	}

	// Страница коэффицентов
	indicatorIsSelected := false
	var statusText string

	switch query.Data {
	case "price":
		indicatorIsSelected = true
		userInputs[chatId].Indicator = "price"
	case "pbv":
		indicatorIsSelected = true
		userInputs[chatId].Indicator = "pbv"
	case "ps":
		indicatorIsSelected = true
		userInputs[chatId].Indicator = "ps"
	case "pe":
		indicatorIsSelected = true
		userInputs[chatId].Indicator = "pe"
	}

	if indicatorIsSelected && userInputs[chatId].Ticker == "" {
		userInputs[chatId].Status = "waiting_for_ticker"

		_, _ = bot.SendMessage(tu.Message(tu.ID(chatId), "🚫Эмитент не был выбран!\nВведите тикер 📈 с клавиатуры 🔠").WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🔙 Назад").WithCallbackData("back"),
			))))

	} else if indicatorIsSelected {

		userInputs[chatId].Status = "waiting_for_value"
		statusText = fmt.Sprintf("Для эмитента 📈 %s выбран индикатор 📐%s", userInputs[chatId].Ticker, userInputs[chatId].Indicator)
		_, _ = bot.SendMessage(tu.Message(tu.ID(chatId), statusText))
		_, _ = bot.SendMessage(tu.Message(tu.ID(chatId), fmt.Sprintf("Введите значение %s 🔢", userInputs[chatId].Indicator)).WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🔙 Назад").WithCallbackData("back"),
			))))
	}

	// Страница целей
	switch query.Data {
	case "add_target":
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Введите тикер эмитента 📈 с клавиатуры 🔠").WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🔙 Назад").WithCallbackData("back"),
			),
		)))
		userInputs[query.Message.GetChat().ID].Status = "waiting_for_ticker"
	case "my_targets":
		// Здесь вы можете добавить логику для отображения целей
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "🔍 Цели еще не добавлены").WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("🔙 Назад").WithCallbackData("back"),
			),
		)))
	case "back": // Кнопка назад
		//delete(userInputs, query.Message.GetChat().ID)
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.GetChat().ID), "Выберите действие").WithReplyMarkup(tu.InlineKeyboard(
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Добавить цель").WithCallbackData("add_target"),
			),
			tu.InlineKeyboardRow(
				tu.InlineKeyboardButton("Мои цели").WithCallbackData("my_targets"),
			),
		)))
	}

	// Ответ вверху чата пользователя
	currSelectionText := service.GetTgCurrentSelectionText(userInputs[chatId].Ticker, userInputs[chatId].Indicator)
	_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText(currSelectionText))
}
