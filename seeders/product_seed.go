package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedProducts(db *gorm.DB) {
	products := []models.Product{
		{ID: 1, Name: "Laptop", Description: "A powerful laptop", Price: 1000, Stock: 10, CategoryID: 1},
		{ID: 2, Name: "Book on Go", Description: "Learn Go programming", Price: 20, Stock: 50, CategoryID: 2},
		{ID: 3, Name: "T-Shirt", Description: "Comfortable T-shirt", Price: 15, Stock: 100, CategoryID: 3},
	}
	for _, product := range products {
		if err := db.FirstOrCreate(&product, models.Product{ID: product.ID}).Error; err != nil {
			log.Printf("Failed to seed product with ID %d: %v", product.ID, err)
		}
	}
	log.Println("Products seeded successfully")
}
