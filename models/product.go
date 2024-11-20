package models

type Product struct {
	ID          int      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string   `gorm:"not null" json:"name" binding:"required"`
	Description string   `json:"description"`
	Price       float64  `gorm:"not null" json:"price" binding:"required"`
	Stock       int      `gorm:"not null" json:"stock" binding:"required"`
	CategoryID  int      `json:"category_id"`
	Category    Category `gorm:"foreignKey:CategoryID" json:"category"`
}
