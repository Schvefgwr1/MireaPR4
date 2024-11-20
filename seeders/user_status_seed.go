package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedUserStatuses(db *gorm.DB) {
	statuses := []models.UserStatus{
		{ID: 1, Name: "Active"},
		{ID: 2, Name: "Inactive"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.UserStatus{ID: status.ID}).Error; err != nil {
			log.Printf("Failed to seed user status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("UserStatuses seeded successfully")
}
