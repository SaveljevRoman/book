package db

import (
	"book/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal(err)
	}

	return db
}
