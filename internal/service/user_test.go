package service

import (
	"errors"
	"fin_api_gateway/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockUser struct {
	Email    string
	Password string
}

func (u *MockUser) CheckPassword(password string) error {
	if password != u.Password {
		return errors.New("incorrect password")
	}
	return nil
}

func TestAuthenticate_UserFoundAndPasswordCorrect(t *testing.T) {
	// Создаем тестовую базу данных SQLite
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Автоматически мигрируем схемы
	err = db.AutoMigrate(&entities.User{})
	assert.NoError(t, err)

	// Создаем пользователя
	testUser := entities.User{
		Password: "password123",
		UserBase: entities.UserBase{
			Email: "test@example.com",
		},
	}

	testUser.IsActive = true

	// Валидация
	testUser.Validate()

	db.Create(&testUser)

	// Создаем объект UserAuth для аутентификации
	ua := &UserAuth{Email: "test@example.com", Password: "password123"}

	// Выполняем аутентификацию
	result, err := Authenticate(db, ua)

	// Проверяем результат
	assert.NoError(t, err)
	assert.Equal(t, testUser.Email, result.Email)
}

func TestAuthenticate_UserNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&entities.User{})
	assert.NoError(t, err)

	ua := &UserAuth{Email: "notfound@example.com", Password: "password123"}

	result, err := Authenticate(db, ua)

	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
	assert.Empty(t, result.Email)
}

func TestAuthenticate_IncorrectPassword(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&MockUser{})
	assert.NoError(t, err)

	testUser := entities.User{
		Password: "password123",
		UserBase: entities.UserBase{
			Email: "test@example.com",
		},
	}
	testUser.IsActive = true
	db.Create(&testUser)

	ua := &UserAuth{Email: "test@example.com", Password: "wrongpassword"}

	_, err = Authenticate(db, ua)

	assert.Error(t, err)
	assert.Equal(t, "incorrect password", err.Error())
}
