package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{
			ID:   1,
			Name: "Admin",
			Permissions: []models.Permission{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
		},
		{
			ID:   2,
			Name: "User",
			Permissions: []models.Permission{
				{ID: 1},
			},
		},
	}
	for _, role := range roles {
		if err := db.Create(&role).Error; err != nil {
			log.Printf("Failed to seed role with ID %d: %v", role.ID, err)
		}
	}
	log.Println("Roles seeded successfully")
}
