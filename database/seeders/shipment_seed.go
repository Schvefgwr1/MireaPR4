package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedShipments(db *gorm.DB) {
	shipments := []models.Shipment{
		{ID: 1, OrderID: 1, AddressID: 1, StatusID: 1},
		{ID: 2, OrderID: 2, AddressID: 2, StatusID: 2},
	}
	for _, shipment := range shipments {
		if err := db.FirstOrCreate(&shipment, models.Shipment{ID: shipment.ID}).Error; err != nil {
			log.Printf("Failed to seed shipment with ID %d: %v", shipment.ID, err)
		}
	}
	log.Println("Shipments seeded successfully")
}
