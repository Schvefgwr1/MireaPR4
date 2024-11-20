package seeders

import (
	"MireaPR4/models"
	"gorm.io/gorm"
	"log"
)

func SeedOrderItems(db *gorm.DB) {
	orderItems := []models.OrderItem{
		{ID: 1, OrderID: 1, ProductID: 1, Quantity: 1, Price: 1000},
		{ID: 2, OrderID: 1, ProductID: 2, Quantity: 1, Price: 20},
		{ID: 3, OrderID: 2, ProductID: 3, Quantity: 2, Price: 30},
	}
	for _, item := range orderItems {
		if err := db.FirstOrCreate(&item, models.OrderItem{ID: item.ID}).Error; err != nil {
			log.Printf("Failed to seed order item with ID %d: %v", item.ID, err)
		}
	}
	log.Println("OrderItems seeded successfully")
}
