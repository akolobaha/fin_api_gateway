package service

import (
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	"fmt"
	"strings"
)

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

func GetTgCurrentSelectionText(ticker string, indicator string) string {
	var sb strings.Builder
	sb.WriteString("✅ Текущий выбор: ")
	if ticker != "" {
		sb.WriteString(fmt.Sprintf("📈 %s ", ticker))
	}
	if indicator != "" {
		sb.WriteString(fmt.Sprintf("📐 %s", indicator))
	}
	return sb.String()
}
