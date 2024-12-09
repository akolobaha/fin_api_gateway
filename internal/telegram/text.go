package telegram

import "fmt"

func WelcomeText(firstName string) string {
	return fmt.Sprintf(
		`Здравствуйте, %s! 
investBOT умеет рассылать уведомления при достижении заданных инвестором целей по: цене, P/Bv, P/E, P/S
Пример цели:
Цена GAZP (Газпром) по 115 рублей, при достижении этой цены investBOT отправит уведомление.
Для постновки цели достаточно ввести тикер эмитента и значение цели.
`, firstName)
}
