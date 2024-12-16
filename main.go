package main

import (
	config "MireaPR4/configs"
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/database/seeders"
	"MireaPR4/docs"
	"MireaPR4/http/controllers"
	addressHandlers "MireaPR4/http/handlers/address"
	categoryHandlers "MireaPR4/http/handlers/category"
	employeeHandlers "MireaPR4/http/handlers/employee"
	orderHandlers "MireaPR4/http/handlers/order"
	paymentHandlers "MireaPR4/http/handlers/payment"
	productHandlers "MireaPR4/http/handlers/product"
	registerHandlers "MireaPR4/http/handlers/register"
	roleHandlers "MireaPR4/http/handlers/role"
	shipmentHandlers "MireaPR4/http/handlers/shipment"
	taskspackage "MireaPR4/http/handlers/tasks"
	"MireaPR4/http/jwt"
	"MireaPR4/http/middlewares"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// @title           Market API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @tokenUrl /auth
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//Подключаем конфиг
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatal("Can't load config " + err.Error())
		return
	}

	//Подключаем хост в документацию
	docs.SwaggerInfo.Host = config.GetActualHost() + ":" + strconv.Itoa(cfg.App.Port)

	//Подключаем ключ jwt
	jwt.InitJWTSecret(cfg.JWT.Secret)

	// Настроим параметры подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Database.Host, cfg.Database.User, cfg.Database.Password, cfg.Database.Name, cfg.Database.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//db = db.Debug()

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
	roleRepo := repositories.NewRoleRepository(db)
	permissionRepo := repositories.NewPermissionRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)
	addressRepo := repositories.NewAddressRepository(db)
	employeeRepo := repositories.NewEmployeeRepository(db)
	paymentRepo := repositories.NewPaymentRepository(db)
	shipmentRepo := repositories.NewShipmentRepository(db)
	shipmentStatusRepo := repositories.NewShipmentStatusRepository(db)
	orderItemRepo := repositories.NewOrderItemRepository(db)
	paymentStatusRepo := repositories.NewPaymentStatusRepository(db)

	//Инициализация middlewares
	middlewares.InitDB(&userRepo)

	//Инициализация order контроллеров
	orderController := controllers.NewOrderController(
		orderRepo,
		userRepo,
		orderStatusRepo,
		productRepo,
		orderItemRepo,
	)
	orderHandler := orderHandlers.NewOrderHandler(orderController)

	//Инициализация register контроллеров
	registerController := controllers.NewRegisterController(userRepo)
	registerHandler := registerHandlers.NewRegisterHandler(registerController)

	//Инициализация roles контроллеров
	roleController := controllers.NewRoleController(roleRepo, permissionRepo)
	roleHandler := roleHandlers.NewRoleHandler(roleController)

	//Инициализация category контроллеров
	categoryController := controllers.NewCategoryController(categoryRepo)
	categoryHandler := categoryHandlers.NewCategoryHandler(categoryController)

	//Инициализация address контроллеров
	addressController := controllers.NewAddressController(addressRepo)
	addressHandler := addressHandlers.NewAddressHandler(addressController)

	//Инициализация employee контроллеров
	employeeController := controllers.NewEmployeeController(employeeRepo)
	employeeHandler := employeeHandlers.NewEmployeeHandler(employeeController)

	//Инициализация payment контроллеров
	paymentController := controllers.NewPaymentController(paymentRepo, paymentStatusRepo)
	paymentHandler := paymentHandlers.NewPaymentHandler(paymentController)

	//Инициализация shipment контроллеров
	shipmentController := controllers.NewShipmentController(shipmentRepo, orderRepo, shipmentStatusRepo)
	shipmentHandler := shipmentHandlers.NewShipmentHandler(shipmentController)

	//Инициализация product контроллеров
	productController := controllers.NewProductController(productRepo, categoryRepo)
	productHandler := productHandlers.NewProductHandler(productController)

	// Регистрируем роуты и запускаем сервер
	r := gin.Default()

	// Настройка маршрута для Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderHandler.RegisterRoutes(r)
	registerHandler.RegisterRoutes(r)
	roleHandler.RegisterRoutes(r)
	categoryHandler.RegisterRoutes(r)
	addressHandler.RegisterRoutes(r)
	employeeHandler.RegisterRoutes(r)
	paymentHandler.RegisterRoutes(r)
	shipmentHandler.RegisterRoutes(r)
	productHandler.RegisterRoutes(r)
	taskspackage.RegisterRoutes(r)

	_ = r.Run(":" + strconv.Itoa(cfg.App.Port))
}
