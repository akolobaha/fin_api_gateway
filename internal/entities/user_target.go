package entities

import (
	"context"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserTarget struct {
	ID     uint    `gorm:"primaryKey;column:id" json:"id"`
	Ticker string  `gorm:"not null;column:ticker" json:"ticker" validate:"required"`
	UserId int64   `gorm:"column:user_id" json:"-"`
	PE     float32 `gorm:"column:p_e_msfo_target" json:"pe,omitempty"`
	PBv    float32 `gorm:"column:p_bv_msfo_target" json:"pbv,omitempty"`
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
