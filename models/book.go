package models

type Book struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"not null" json:"title" binding:"required"`
	Author string `gorm:"not null" json:"author" binding:"required"`
}
