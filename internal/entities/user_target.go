package entities

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

const NOTIFICATION_TELEGRAM = "telegram"
const NOTIFICATION_EMAIL = "email"
const NOTIFICATION_SMS = "sms"

const REPORT_RSBU = "rsbu"
const REPORT_MSFO = "msfo"

const VALUATION_RATIO_PE = "pe"
const VALUATION_RATIO_PBV = "pbv"

type UserTarget struct {
	ID                 uint    `gorm:"primaryKey;column:id" json:"id"`
	Ticker             string  `gorm:"not null;column:ticker" json:"ticker" validate:"required"`
	UserId             int64   `gorm:"column:user_id" json:"-"`
	ValuationRatio     string  `gorm:"not null;column:valuation_ratio" json:"valuation_ratio" validate:"required,oneof=pe pbv"`
	Value              float32 `gorm:"not null;column:value" json:"value"`
	FinancialReport    string  `gorm:"not null;column:financial_report;default:rsbu" json:"financial_report" validate:"required,oneof=rsbu msfo"`
	Achieved           bool    `gorm:"not null;column:achieved;default:false" json:"-"`
	NotificationMethod string  `gorm:"not null;column:notification_method" json:"notification_method" validate:"required,oneof=sms email telegram"`
}

func (usf *UserTarget) Save(ctx context.Context, conn *gorm.DB) error {
	userId := ctx.Value("userId").(int64)
	usf.UserId = userId

	result := conn.Save(usf)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func (usf *UserTarget) Validate() error {
	validate := validator.New()
	return validate.Struct(usf)
}
