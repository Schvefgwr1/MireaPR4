package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedPermissions(db *gorm.DB) {
	permissions := []models.Permission{
		{ID: 1, Name: "View and delete Orders"},
		{ID: 2, Name: "Manage Users"},
		{ID: 3, Name: "Edit Products"},
		{ID: 4, Name: "Edit categorical data"},
		{ID: 5, Name: "See all data"},
		{ID: 6, Name: "Run tasks"},
		{ID: 7, Name: "Delete some data"},
		{ID: 8, Name: "Edit payments"},
		{ID: 9, Name: "Modify roles"},
	}
	for _, permission := range permissions {
		if err := db.FirstOrCreate(&permission, models.Permission{ID: permission.ID}).Error; err != nil {
			log.Printf("Failed to seed permission with ID %d: %v", permission.ID, err)
		}
	}
	log.Println("Permissions seeded successfully")
}
