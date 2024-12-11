package entities

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/log"
	"fmt"
)

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

func FirstOrCreateTgUser(fromId int64, username string) *TgUser {
	// Инициализация БД
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()

	user := NewTgUser(fromId, username)

	result := gDB.Where(user).Attrs(user).FirstOrCreate(&user)

	if result.RowsAffected > 0 {
		fmt.Println("Inserted new user:", user)
	} else {
		fmt.Println("User already exists:", user)
	}

	return user
}
