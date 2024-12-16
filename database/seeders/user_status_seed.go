package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedUserStatuses(db *gorm.DB) {
	statuses := []models.UserStatus{
		{Name: "Active"},
		{Name: "Inactive"},
	}
	for _, status := range statuses {
		if err := db.FirstOrCreate(&status, models.UserStatus{Name: status.Name}).Error; err != nil {
			log.Printf("Failed to seed user status with ID %d: %v", status.ID, err)
		}
	}
	log.Println("UserStatuses seeded successfully")
}
