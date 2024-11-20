package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedCategories(db *gorm.DB) {
	categories := []models.Category{
		{ID: 1, Name: "Electronics"},
		{ID: 2, Name: "Books"},
		{ID: 3, Name: "Clothing"},
	}
	for _, category := range categories {
		if err := db.FirstOrCreate(&category, models.Category{ID: category.ID}).Error; err != nil {
			log.Printf("Failed to seed category with ID %d: %v", category.ID, err)
		}
	}
	log.Println("Categories seeded successfully")
}
