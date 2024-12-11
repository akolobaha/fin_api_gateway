package entities

type TgUser struct {
	ID             int64 `gorm:"primaryKey"`
	TelegramUserID int64
	Username       string
}

func NewTgUser(tgUserId int64, username string) *TgUser {
	return &TgUser{
		TelegramUserID: tgUserId,
		Username:       username,
	}
}
