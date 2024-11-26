package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedOrderStatuses(db *gorm.DB) {
	statuses := []models.OrderStatus{
		{ID: 1, Name: "Pending"},
		{ID: 2, Name: "Completed"},
		{ID: 3, Name: "Cancelled"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.OrderStatus{ID: status.ID}).Error; err != nil {
			log.Printf("Failed to seed order status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("OrderStatuses seeded successfully")
}
