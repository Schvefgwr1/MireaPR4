package main

import (
	"MireaPR4/controllers"
	"MireaPR4/handlers"
	"MireaPR4/repositories"
	"MireaPR4/services"
	"github.com/gin-gonic/gin"
)

func main() {
	/*
		// Настроим параметры подключения
		dsn := "host=localhost user=postgres password=postgres dbname=mirea_books port=5432 sslmode=disable"
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}

		// Запускаем миграцию для создания таблицы
		if err := db.AutoMigrate(&models.Book{}); err != nil {
			panic("failed to migrate database")
		}

		//Заполняем БД тестовыми данными
		seeders.SeedData(db)
	*/

	// Инициализация репозиториев, сервисов и контроллеров
	//bookRepo := repositories.NewBookRepository(db)
	bookRepo := repositories.NewBookRepository()
	bookService := services.NewBookService(bookRepo)
	bookController := controllers.NewBookController(bookService)
	bookHandler := handlers.NewBookHandler(bookController)

	// Регистрируем роуты и запускаем сервер
	r := gin.Default()
	bookHandler.RegisterRoutes(r)

	_ = r.Run(":8080")
}
