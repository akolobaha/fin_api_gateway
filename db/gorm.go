package db

import (
	"fin_api_gateway/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type GormDB struct {
	*gorm.DB
}

func (g *GormDB) Connect() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DbDsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		slog.Error(err.Error())
	}

	return db
}
