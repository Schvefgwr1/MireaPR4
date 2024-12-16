package seeders

import (
	"MireaPR4/database/models"
	"gorm.io/gorm"
	"log"
)

func SeedOrderItems(db *gorm.DB) {
	orderItems := []models.OrderItem{
		{OrderID: 1, ProductID: 1, Quantity: 1, Price: 1000},
		{OrderID: 1, ProductID: 2, Quantity: 1, Price: 20},
		{OrderID: 2, ProductID: 3, Quantity: 2, Price: 30},
	}
	for _, item := range orderItems {
		if err := db.FirstOrCreate(&item, models.OrderItem{
			OrderID: item.OrderID, ProductID: item.ProductID,
		}).Error; err != nil {
			log.Printf("Failed to seed order item with ID %d: %v", item.ID, err)
		}
	}
	log.Println("OrderItems seeded successfully")
}
