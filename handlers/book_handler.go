package handlers

import (
	"MireaPR4/controllers"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	controller *controllers.BookController
}

func NewBookHandler(controller *controllers.BookController) *BookHandler {
	return &BookHandler{controller}
}

func (h *BookHandler) RegisterRoutes(router *gin.Engine) {
	books := router.Group("/books")
	{
		books.POST("/", h.controller.CreateBook)
		books.GET("/", h.controller.GetBooks)
		books.GET("/:id", h.controller.GetBookByID)
		books.PUT("/:id", h.controller.UpdateBook)
		books.DELETE("/:id", h.controller.DeleteBook)
	}
}
