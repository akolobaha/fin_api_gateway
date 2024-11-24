package service

import (
	"errors"
	"fin_api_gateway/internal/entities"
	"gorm.io/gorm"
	"log/slog"
)

type UserAuth struct {
	Email    string `gorm:"primaryKey" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
}

func Authenticate(conn *gorm.DB, ua *UserAuth) (entities.User, error) {
	// Проверка на существование в базе
	var user entities.User
	if err := conn.Where("email = ?", ua.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			slog.Info("Запись не найдена")
			return user, gorm.ErrRecordNotFound
		}
	}

	// Полверка пароля
	err := user.CheckPassword(ua.Password)
	if err != nil {
		slog.Info("Пароль не верный")
		return user, errors.New("incorrect password")
	}

	return user, nil
}
