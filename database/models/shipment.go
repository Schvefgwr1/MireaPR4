package models

import "time"

type Shipment struct {
	ID           int            `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID      int            `gorm:"not null" json:"order_id"`
	ShipmentDate time.Time      `gorm:"autoCreateTime" json:"shipment_date"`
	StatusID     int            `json:"status_id"`
	AddressID    int            `json:"address_id"`
	Status       ShipmentStatus `gorm:"foreignKey:StatusID" json:"status"`
	Address      Address        `gorm:"foreignKey:AddressID" json:"address"`
}
