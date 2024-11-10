package seeders

import "gorm.io/gorm"

func SeedData(db *gorm.DB) {
	SeedBooks(db)
}
