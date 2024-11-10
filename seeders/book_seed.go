package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedBooks(db *gorm.DB) {
	books := []models.Book{
		{ID: "1", Title: "Go Programming", Author: "John Doe"},
		{ID: "2", Title: "Learning PostgreSQL", Author: "Jane Smith"},
		{ID: "3", Title: "Microservices in Go", Author: "Alex Johnson"},
	}

	for _, book := range books {
		if err := db.FirstOrCreate(&book, models.Book{ID: book.ID}).Error; err != nil {
			log.Printf("failed to seed book with ID %s: %v", book.ID, err)
		}
	}
	log.Println("Books seeded successfully")
}
