package db

import (
	"ETL/internal/models"
	"encoding/json"
	"gorm.io/gorm"
)

func SaveCategoriesForEvent(db *gorm.DB, eventID uint, categoriesJSON []byte) error {
	var rawCategories models.RawCategories
	if err := json.Unmarshal(categoriesJSON, &rawCategories); err != nil {
		return err
	}

	for _, categoryName := range rawCategories.Categories {
		category := models.Category{
			EventID: eventID,
			Name:    categoryName,
		}
		if err := db.Create(&category).Error; err != nil {
			return err
		}
	}
	return nil
}
