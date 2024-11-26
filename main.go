package main

import (
	config "MireaPR4/configs"
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/database/seeders"
	"MireaPR4/http/controllers"
	orderHandlers "MireaPR4/http/handlers/order"
	registerHandlers "MireaPR4/http/handlers/register"
	"MireaPR4/http/jwt"
	"MireaPR4/http/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func main() {
	//Подключаем конфиг
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal("Can't load config " + err.Error())
		return
	}

	//Подключаем ключ jwt
	jwt.InitJWTSecret(cfg.JWT.Secret)

	// Настроим параметры подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port,
	)
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

	// Инициализация репозиториев
	userRepo := repositories.NewUserRepository(db)
	orderRepo := repositories.NewOrderRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderStatusRepo := repositories.NewOrderStatusRepository(db)

	middlewares.InitDB(&userRepo)

	//Инициализация order контроллеров
	orderController := controllers.NewOrderController(
		orderRepo,
		userRepo,
		orderStatusRepo,
		productRepo,
	)
	orderHandler := orderHandlers.NewOrderHandler(orderController)

	//Инициализация register контроллеров
	registerController := controllers.NewRegisterController(userRepo)
	registerHandler := registerHandlers.NewRegisterHandler(registerController)

	// Регистрируем роуты и запускаем сервер
	r := gin.Default()
	orderHandler.RegisterRoutes(r)
	registerHandler.RegisterRoutes(r)

	_ = r.Run(":" + strconv.Itoa(cfg.App.Port))
}
