package entities

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                     int64   `gorm:"primaryKey" json:"id"`
	Password               string  `gorm:"not null" json:"password" validate:"required,min=6"`
	EmailConfirmationToken *string `gorm:"default:null"`
	IsActive               bool    `gorm:"default:false"`
	UserBase
}

type UserResponse struct {
	ID int64 `gorm:"primaryKey" json:"user_id"`
	UserBase
}

type UserBase struct {
	Name  string `gorm:"not null" json:"name" validate:"required,min=3,max=100"`
	Email string `gorm:"not null" json:"email" validate:"required,email"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *User) SetPasswordHash() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) SetEmailConfirmationToken() error {
	byteToken := make([]byte, 32)
	if _, err := rand.Read(byteToken); err != nil {
		return err
	}

	token := base64.RawURLEncoding.EncodeToString(byteToken)
	u.EmailConfirmationToken = &token
	return nil
}
