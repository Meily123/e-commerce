package config

import (
	"WebAPI/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToDatabase() *gorm.DB {

	connection := os.Getenv("DATABASE_URL")

	DB, err := gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatal("cant connect to database", err)
	}

	err = DB.AutoMigrate(model.CartProduct{}, &model.Product{}, &model.User{})
	if err != nil {
		log.Fatal("cant migrate database", err)
	}

	return DB
}
