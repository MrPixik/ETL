package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const dsn = "YOUR_DSN"

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
		return nil, err
	}

	log.Println("Connected to DB!")
	return db, nil

}
