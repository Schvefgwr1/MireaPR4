package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{ID: 1, Username: "admin", Password: "adminpass", Email: "admin@example.com", RoleID: 1, StatusID: 1},
		{ID: 2, Username: "john", Password: "userpass", Email: "john@example.com", RoleID: 2, StatusID: 1},
	}
	for _, user := range users {
		if err := db.FirstOrCreate(&user, models.User{ID: user.ID}).Error; err != nil {
			log.Printf("Failed to seed user with ID %d: %v", user.ID, err)
		}
	}
	log.Println("Users seeded successfully")
}
