package entities

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Password string `gorm:"not null" json:"password" validate:"required,min=6"`
	UserBase
}

type UserResponse struct {
	ID int64 `gorm:"primaryKey" json:"user_id"`
	UserBase
}

type UserBase struct {
	Name     string `gorm:"not null" json:"name" validate:"required,min=3,max=100"`
	Email    string `gorm:"not null" json:"email" validate:"required,email"`
	Telegram string `gorm:"not null" json:"telegram" validate:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
