package models

import "time"

type Order struct {
	ID         int          `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int          `json:"user_id"`
	StatusID   int          `json:"status_id"`
	TotalPrice float64      `json:"total_price"`
	CreatedAt  time.Time    `gorm:"autoCreateTime" json:"created_at"`
	Status     OrderStatus  `gorm:"foreignKey:StatusID" json:"status"`
	OrderItems []*OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
	Shipments  []Shipment   `gorm:"foreignKey:OrderID" json:"shipments"`
	Payments   []Payment    `gorm:"foreignKey:OrderID" json:"payments"`
}
