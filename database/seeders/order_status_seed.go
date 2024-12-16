package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedOrderStatuses(db *gorm.DB) {
	statuses := []models.OrderStatus{
		{Name: "Pending"},
		{Name: "Completed"},
		{Name: "Cancelled"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.OrderStatus{Name: status.Name}).Error; err != nil {
			log.Printf("Failed to seed order status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("OrderStatuses seeded successfully")
}
