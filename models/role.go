package models

type Role struct {
	ID          int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"not null" json:"name" binding:"required"`
	Permissions []Permission `gorm:"many2many:roles_permissions_rel;" json:"permissions"`
}
