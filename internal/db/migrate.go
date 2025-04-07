package db

import (
	"ETL/internal/models"
	"gorm.io/gorm"
	"log"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Event{}, &models.Place{}, &models.DateRange{}, &models.Participant{}, &models.Category{})
	if err != nil {
		log.Fatal("migration failed:", err)
		return err
	}
	log.Println("Migration successful!")
	return nil
}
