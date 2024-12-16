package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedCategories(db *gorm.DB) {
	categories := []models.Category{
		{Name: "Electronics"},
		{Name: "Books"},
		{Name: "Clothing"},
	}
	for _, category := range categories {
		if err := db.FirstOrCreate(&category, models.Category{Name: category.Name}).Error; err != nil {
			log.Printf("Failed to seed category with ID %d: %v", category.ID, err)
		}
	}
	log.Println("Categories seeded successfully")
}
