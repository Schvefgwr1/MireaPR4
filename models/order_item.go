package models

type OrderItem struct {
	ID        int     `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   int     `gorm:"not null" json:"order_id"`
	ProductID int     `gorm:"not null" json:"product_id"`
	Quantity  int     `gorm:"not null" json:"quantity" binding:"required"`
	Price     float64 `gorm:"not null" json:"price" binding:"required"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
