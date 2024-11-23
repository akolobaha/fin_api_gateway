package entities

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fin_api_gateway/db"
	"gorm.io/gorm"
	"time"
)

const expirationPeriod time.Duration = 1 // days

type UserToken struct {
	Token          string    `gorm:"token not null"`
	UserId         int64     `gorm:"user_id"`
	ExpirationTime time.Time `gorm:"expiration_time"`
}

func FindOrCreateToken(userId int64) (*UserToken, error) {
	dbGorm := new(db.GormDB).Connect()
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	userToken := UserToken{}

	if err := dbGorm.Where("user_id = ? AND expiration_time > ?", userId, formattedTime).First(&userToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			expirationTime := currentTime.Add(time.Duration(expirationPeriod) * 24 * time.Hour)
			token, err := generateRandomToken(32)
			if err != nil {
				return nil, err
			}
			newToken := UserToken{UserId: int64(userId), ExpirationTime: expirationTime, Token: token}
			dbGorm.Create(&newToken)
			return &newToken, nil
		}
	}

	return &userToken, nil
}

func generateRandomToken(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
