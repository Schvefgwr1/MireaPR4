package models

type Employee struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID     int    `json:"user_id" binding:"required"`
	Position   string `json:"position" binding:"required"`
	Department string `json:"department" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
	User       User   `gorm:"foreignKey:UserID" json:"-"`
}
