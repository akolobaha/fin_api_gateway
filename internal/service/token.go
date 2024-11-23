package service

import (
	"errors"
	"fin_api_gateway/db"
	"fin_api_gateway/internal/entities"
	"gorm.io/gorm"
	"log/slog"
	"time"
)

func GetTokenEntityByToken(token string) (entities.UserToken, error) {
	gDB := new(db.GormDB).Connect()
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	userToken := entities.UserToken{}

	if err := gDB.Where("token = ? AND expiration_time > ?", token, formattedTime).First(&userToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Токен не найден или истек")
			return userToken, err
		}
	}
	return userToken, nil
}
