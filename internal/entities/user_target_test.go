package entities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUserTarget_Validate_EmptyTicker(t *testing.T) {
	usf := &UserTarget{
		Ticker: "", // Empty ticker for validation check
	}

	err := usf.Validate()
	assert.Error(t, err, "Expected validation error for empty ticker")
}

func TestUserTarget_Validate_Valid(t *testing.T) {
	usf := &UserTarget{
		Ticker:             "AAPL", // Valid ticker
		ValuationRatio:     "pe",   // Valid valuation ratio
		Value:              25.0,   // Valid value
		FinancialReport:    "rsbu", // Valid financial report
		NotificationMethod: "sms",  // Valid notification method
	}

	err := usf.Validate()
	assert.NoError(t, err, "Expected no validation error for valid UserTarget")
}

func TestUserTarget_Save(t *testing.T) {
	// Create an in-memory test database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	// Automatically migrate the structure
	err = db.AutoMigrate(&UserTarget{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	usf := &UserTarget{
		Ticker:             "AAPL",
		ValuationRatio:     "pe",
		Value:              25.0,
		FinancialReport:    "rsbu",
		NotificationMethod: "sms",
	}

	// Create context with userId
	ctx := context.WithValue(context.Background(), "userId", int64(1))

	// Save the structure
	err = usf.Save(ctx, db)
	assert.NoError(t, err, "Expected no error while saving")

	// Check that the structure is saved in the database
	var savedUsf UserTarget
	result := db.First(&savedUsf, usf.ID)
	assert.NoError(t, result.Error, "Expected to find the saved UserTarget")
	assert.Equal(t, usf.Ticker, savedUsf.Ticker)
	assert.Equal(t, int64(1), savedUsf.UserId)
	assert.Equal(t, usf.ValuationRatio, savedUsf.ValuationRatio)
	assert.Equal(t, usf.Value, savedUsf.Value)
	assert.Equal(t, usf.FinancialReport, savedUsf.FinancialReport)
	assert.Equal(t, usf.NotificationMethod, savedUsf.NotificationMethod)
}
