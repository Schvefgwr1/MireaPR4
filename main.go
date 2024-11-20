package main

import (
	"MireaPR4/controllers"
	"MireaPR4/handlers/order"
	"MireaPR4/models"
	"MireaPR4/repositories"
	"MireaPR4/seeders"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Настроим параметры подключения
	dsn := "host=localhost user=postgres password=postgres dbname=mirea_books port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Запускаем миграцию для создания таблиц
	if err := db.AutoMigrate(
		&models.Role{},
		&models.Permission{},
		&models.UserStatus{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.OrderStatus{},
		&models.Order{},
		&models.OrderItem{},
		&models.PaymentStatus{},
		&models.Payment{},
		&models.ShipmentStatus{},
		&models.Address{},
		&models.Shipment{},
		&models.Employee{},
	); err != nil {
		panic("failed to migrate database")
	}
	//Заполняем БД тестовыми данными
	seeders.SeedData(db)

	// Инициализация order контроллеров
	userRepo := repositories.NewUserRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderStatusRepo := repositories.NewOrderStatusRepository(db)
	orderController := controllers.NewOrderController(
		orderRepo,
		userRepo,
		orderStatusRepo,
		productRepo,
	)
	orderHandler := order.NewOrderHandler(orderController)

	// Регистрируем роуты и запускаем сервер
	r := gin.Default()
	orderHandler.RegisterRoutes(r)

	_ = r.Run(":8080")
}
