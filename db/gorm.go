package db

import (
	"fin_api_gateway/internal/config"
	"fin_api_gateway/internal/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	*gorm.DB
}

func ConnectToDB() (*Connection, error) {
	gDB := &Connection{}
	if err := gDB.Connect(); err != nil {
		return nil, err
	}
	return gDB, nil
}

func (g *Connection) Connect() error {
	var err error
	g.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DbDsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Error("Failed to connect to database: ", err)
		return err
	}

	return nil
}

func (g *Connection) Close() error {
	if g.DB == nil {
		return nil // Если DB уже nil, ничего не делать
	}

	dbInstance, err := g.DB.DB()
	if err != nil {
		log.Error("Failed to get DB instance: ", err)
		return err
	}

	err = dbInstance.Close()
	if err != nil {
		log.Error("Failed to close database connection: ", err)
		return err
	}

	g.DB = nil // Устанавливаем DB в nil после закрытия
	return nil
}
