package models

import "time"

type Payment struct {
	ID          int           `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID     int           `gorm:"not null" json:"order_id"`
	Amount      float64       `gorm:"not null" json:"amount" binding:"required"`
	PaymentDate time.Time     `gorm:"autoCreateTime" json:"payment_date"`
	StatusID    int           `json:"status_id"`
	Status      PaymentStatus `gorm:"foreignKey:StatusID" json:"status"`
}
