package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedPaymentStatuses(db *gorm.DB) {
	statuses := []models.PaymentStatus{
		{ID: 1, Name: "Pending"},
		{ID: 2, Name: "Paid"},
		{ID: 3, Name: "Failed"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.PaymentStatus{ID: status.ID}).Error; err != nil {
			log.Printf("Failed to seed payment status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("PaymentStatuses seeded successfully")
}
