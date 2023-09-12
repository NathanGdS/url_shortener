package db

import (
	"log"

	"url_shortener/helpers"
	"url_shortener/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := helpers.GetEnvVariable("DB_URL")

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.UrlShortener{})

	return db
}
