package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedAddresses(db *gorm.DB) {
	addresses := []models.Address{
		{ID: 1, City: "New York", Street: "5th Ave", House: 123, Index: "10001", Flat: 10},
		{ID: 2, City: "San Francisco", Street: "Market Street", House: 45, Index: "94103", Flat: 5},
	}
	for _, address := range addresses {
		if err := db.FirstOrCreate(&address, models.Address{ID: address.ID}).Error; err != nil {
			log.Printf("Failed to seed address with ID %d: %v", address.ID, err)
		}
	}
	log.Println("Addresses seeded successfully")
}
