package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedPermissions(db *gorm.DB) {
	permissions := []models.Permission{
		{Name: "View and delete Orders"},
		{Name: "Manage Users"},
		{Name: "Edit Products"},
		{Name: "Edit categorical data"},
		{Name: "See all data"},
		{Name: "Run tasks"},
		{Name: "Delete some data"},
		{Name: "Edit payments"},
		{Name: "Modify roles"},
	}
	for _, permission := range permissions {
		if err := db.FirstOrCreate(&permission, models.Permission{Name: permission.Name}).Error; err != nil {
			log.Printf("Failed to seed permission with ID %d: %v", permission.ID, err)
		}
	}
	log.Println("Permissions seeded successfully")
}
