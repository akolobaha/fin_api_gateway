package entities

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/log"
	"strings"
)

type Securities []Security

type Security struct {
	Ticker    string `gorm:"ticker"`
	Shortname string `gorm:"shortname"`
	Secname   string `gorm:"secname"`
}

func ValidateTicker(ticker string) (security Security, isValid bool) {
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
	gDB.Model(&Securities{}).Where("ticker = ?", value).Find(&security).Count(&count)

	return security, count > 0
}
