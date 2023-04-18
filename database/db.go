package database

import (
	"fmt"
	"log"

	"github.com/haviz000/MyGram-RestAPI/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "jaka"
	DB_NAME     = "db_mygram"
	DB_PORT     = "5432"

	db  *gorm.DB
	err error
)

func ConnectDB() {
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("failed connecting to database", err)
	}

	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.SocialMedia{}, model.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
