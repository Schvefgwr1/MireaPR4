package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedShipmentStatuses(db *gorm.DB) {
	statuses := []models.ShipmentStatus{
		{Name: "In Progress"},
		{Name: "Delivered"},
		{Name: "Returned"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.ShipmentStatus{ID: status.ID}).Error; err != nil {
			log.Printf("Failed to seed shipment status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("ShipmentStatuses seeded successfully")
}
