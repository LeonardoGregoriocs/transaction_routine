package database

import (
	"github.com/leonardogregoriocs/internal/domain/client"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	url_connection := "host=localhost user=user password=user123 dbname=dataclient port=5432"
	db, err := gorm.Open(postgres.Open(url_connection), &gorm.Config{})

	if err != nil {
		panic("Fail to connect to database")
	}

	db.AutoMigrate(&client.Client{})

	return db

}
