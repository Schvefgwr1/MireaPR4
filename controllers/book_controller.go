package controllers

import (
	"MireaPR4/models"
	"MireaPR4/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookController struct {
	service services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{service}
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if book.Title == "" || book.Author == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title and Author are required fields"})
		return
	}

	createdBook, err := c.service.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	ctx.JSON(http.StatusCreated, createdBook)
}

func (c *BookController) GetBooks(ctx *gin.Context) {
	books, err := c.service.GetBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

func (c *BookController) GetBookByID(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := c.service.GetBookByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id
	if err := c.service.UpdateBook(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.service.DeleteBook(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
