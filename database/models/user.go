package models

type User struct {
	ID       int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string     `gorm:"not null" json:"username" binding:"required"`
	Password string     `gorm:"not null" json:"password" binding:"required"`
	Email    string     `gorm:"not null" json:"email" binding:"required,email"`
	RoleID   int        `json:"role_id"`
	StatusID int        `json:"status_id"`
	Role     Role       `gorm:"foreignKey:RoleID" json:"role"`
	Status   UserStatus `gorm:"foreignKey:StatusID" json:"status"`
	Orders   []Order    `gorm:"foreignKey:UserID" json:"user_id"`
}
