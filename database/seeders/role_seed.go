package seeders

import (
	models2 "MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedRoles(db *gorm.DB) {
	roles := []models2.Role{
		{
			ID:   1,
			Name: "Admin",
			Permissions: []models2.Permission{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
		{
			ID:   2,
			Name: "User",
			Permissions: []models2.Permission{
				{ID: 1},
			},
		},
	}
	for _, role := range roles {
		if err := db.FirstOrCreate(&role).Error; err != nil {
			log.Printf("Failed to seed role with ID %d: %v", role.ID, err)
		}
	}
	log.Println("Roles seeded successfully")
}
