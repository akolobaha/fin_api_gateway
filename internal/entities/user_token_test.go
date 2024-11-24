package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupDatabase() (*gorm.DB, error) {
	// Создание SQLite базы данных в памяти
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Автоматическое создание таблиц
	if err := db.AutoMigrate(&UserToken{}); err != nil {
		return nil, err
	}

	return db, nil
}

func TestFindOrCreateToken(t *testing.T) {
	db, err := setupDatabase()
	if err != nil {
		t.Fatalf("failed to setup database: %v", err)
	}

	userId := int64(1)

	// Тест 1: Создание нового токена
	token, err := FindOrCreateToken(userId, db)
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Token)
	assert.Equal(t, userId, token.UserId)

	// Тест 2: Проверка, что токен был создан и имеет правильное время истечения
	assert.True(t, token.ExpirationTime.After(time.Now()))

	// Тест 3: Повторный вызов для того же пользователя должен возвращать тот же токен
	sameToken, err := FindOrCreateToken(userId, db)
	assert.NoError(t, err)
	assert.Equal(t, token.Token, sameToken.Token)

	// Тест 4: Проверка истечения токена
	time.Sleep(2 * time.Second) // Ждем 2 секунды, чтобы токен не истек (по времени в тестах)
	// Устанавливаем время истечения токена в прошлое
	token.ExpirationTime = time.Now().Add(-24 * time.Hour)
	db.Save(token)
}
