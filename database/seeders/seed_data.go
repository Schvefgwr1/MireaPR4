package seeders

import "gorm.io/gorm"

func SeedData(db *gorm.DB) {
	SeedPermissions(db)
	SeedRoles(db)

	SeedCategories(db)
	SeedProducts(db)

	SeedUserStatuses(db)
	SeedUsers(db)

	SeedOrderStatuses(db)
	SeedOrders(db)
	SeedOrderItems(db)

	SeedAddresses(db)
	SeedPaymentStatuses(db)
	SeedPayments(db)
	SeedShipmentStatuses(db)
	SeedShipments(db)
	SeedEmployees(db)
}
