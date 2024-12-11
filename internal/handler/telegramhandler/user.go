package telegramhandler

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	"fmt"
	"strings"
)

func FirstOrCreateTgUser(fromId int64, username string) *entities.TgUser {
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

	user := entities.NewTgUser(fromId, username)

	result := gDB.Where(user).Attrs(user).FirstOrCreate(&user)

	if result.RowsAffected > 0 {
		fmt.Println("Inserted new user:", user)
	} else {
		fmt.Println("User already exists:", user)
	}

	return user
}

func ValidateTicker(ticker string) (security entities.Security, isValid bool) {
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()

	value := strings.ToUpper(strings.TrimSpace(ticker))
	var count int64
	gDB.Model(&entities.Securities{}).Where("ticker = ?", value).Find(&security).Count(&count)

	return security, count > 0
}

func GetSecuritiesHint() string {
	gDB := &db.Connection{}
	if err := gDB.Connect(); err != nil {
		log.Error("Could not connect to database: ", err)
	}
	defer func() {
		if err := gDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		}
	}()

	securities := entities.Securities{}
	gDB.Find(&securities)

	var sb strings.Builder
	for _, sec := range securities {
		sb.WriteString(fmt.Sprintf("%s (%s),", sec.Ticker, sec.Shortname))
	}

	result := sb.String()

	return result
}
