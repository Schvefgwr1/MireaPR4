package models

type Address struct {
	ID     int    `gorm:"primaryKey;autoIncrement" json:"id"`
	City   string `gorm:"not null" json:"city" binding:"required"`
	Street string `gorm:"not null" json:"street" binding:"required"`
	House  int    `gorm:"not null" json:"house" binding:"required"`
	Index  string `gorm:"not null" json:"index" binding:"required"`
	Flat   int    `gorm:"not null" json:"flat" binding:"required"`
}
