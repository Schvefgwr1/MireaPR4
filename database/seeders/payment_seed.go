package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedPayments(db *gorm.DB) {
	payments := []models.Payment{
		{ID: 1, OrderID: 1, Amount: 1020, StatusID: 2},
		{ID: 2, OrderID: 2, Amount: 30, StatusID: 1},
	}
	for _, payment := range payments {
		if err := db.FirstOrCreate(&payment, models.Payment{ID: payment.ID}).Error; err != nil {
			log.Printf("Failed to seed payment with ID %d: %v", payment.ID, err)
		}
	}
	log.Println("Payments seeded successfully")
}
