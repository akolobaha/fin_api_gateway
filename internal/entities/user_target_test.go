package entities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUserTarget_Validate(t *testing.T) {
	usf := &UserTarget{
		Ticker: "", // Пустой тикер для проверки валидации
	}

	err := usf.Validate()
	assert.Error(t, err, "Expected validation error for empty ticker")
}

func TestUserTarget_Validate_Valid(t *testing.T) {
	usf := &UserTarget{
		Ticker: "AAPL", // Корректный тикер
	}

	err := usf.Validate()
	assert.NoError(t, err, "Expected no validation error for valid ticker")
}

func TestUserTarget_Save(t *testing.T) {
	// Создаем тестовую базу данных в памяти
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// Автоматически мигрируем структуру
	err = db.AutoMigrate(&UserTarget{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	usf := &UserTarget{
		Ticker: "AAPL",
		PE:     25.0,
		PBv:    5.0,
	}

	// Создаем контекст с userId
	ctx := context.WithValue(context.Background(), "userId", int64(1))

	// Сохраняем структуру
	err = usf.Save(ctx, db)
	assert.NoError(t, err, "Expected no error while saving")

	// Проверяем, что структура сохранена в базе
	var savedUsf UserTarget
	result := db.First(&savedUsf, usf.ID)
	assert.NoError(t, result.Error, "Expected to find the saved UserTarget")
	assert.Equal(t, usf.Ticker, savedUsf.Ticker)
	assert.Equal(t, int64(1), savedUsf.UserId)
}