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

func (g *GormDB) Connect() error {
	var err error
	g.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DbDsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		slog.Error("Failed to connect to database: ", err.Error())
		return err
	}

	return nil
}

func (g *GormDB) Close() error {
	if g.DB == nil {
		return nil // Если DB уже nil, ничего не делать
	}

	dbInstance, err := g.DB.DB()
	if err != nil {
		slog.Error("Failed to get DB instance: ", err.Error())
		return err
	}

	err = dbInstance.Close()
	if err != nil {
		slog.Error("Failed to close database connection: ", err.Error())
		return err
	}

	g.DB = nil // Устанавливаем DB в nil после закрытия
	return nil
}
