package models

type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name" binding:"required"`
}
