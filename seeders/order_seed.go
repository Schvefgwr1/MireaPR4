package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedOrders(db *gorm.DB) {
	orders := []models.Order{
		{ID: 1, UserID: 1, StatusID: 1, TotalPrice: 1020},
		{ID: 2, UserID: 2, StatusID: 2, TotalPrice: 30},
	}
	for _, order := range orders {
		if err := db.FirstOrCreate(&order, models.Order{ID: order.ID}).Error; err != nil {
			log.Printf("Failed to seed order with ID %d: %v", order.ID, err)
		}
	}
	log.Println("Orders seeded successfully")
}
