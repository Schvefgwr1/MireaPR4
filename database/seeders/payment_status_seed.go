package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedPaymentStatuses(db *gorm.DB) {
	statuses := []models.PaymentStatus{
		{Name: "Pending"},
		{Name: "Paid"},
		{Name: "Failed"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.PaymentStatus{Name: status.Name}).Error; err != nil {
			log.Printf("Failed to seed payment status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("PaymentStatuses seeded successfully")
}
