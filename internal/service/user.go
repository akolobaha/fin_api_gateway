package service

import (
	"errors"
	"fin_api_gateway/internal/entities"
	"fin_api_gateway/internal/log"
	"gorm.io/gorm"
)

type UserAuth struct {
	Email    string `gorm:"primaryKey" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
}

func Authenticate(conn *gorm.DB, ua *UserAuth) (entities.User, error) {
	// Проверка на существование в базе
	var user entities.User
	if err := conn.Where("email = ? AND is_active = TRUE", ua.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Info("Запись не найдена")
			return user, gorm.ErrRecordNotFound
		}
	}

	// Полверка пароля
	err := user.CheckPassword(ua.Password)
	if err != nil {
		log.Info("Пароль не верный")
		return user, errors.New("incorrect password")
	}

	return user, nil
}
