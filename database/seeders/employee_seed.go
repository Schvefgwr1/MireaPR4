package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedEmployees(db *gorm.DB) {
	employees := []models.Employee{
		{UserID: 1, Position: "Manager", Department: "Sales", Phone: "+123456789", Email: "manager@example.com"},
		{UserID: 2, Position: "Support", Department: "Customer Service", Phone: "+987654321", Email: "support@example.com"},
	}
	for _, employee := range employees {
		if err := db.FirstOrCreate(&employee, models.Employee{UserID: employee.UserID}).Error; err != nil {
			log.Printf("Failed to seed employee with ID %d: %v", employee.ID, err)
		}
	}
	log.Println("Employees seeded successfully")
}
