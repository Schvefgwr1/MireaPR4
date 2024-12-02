package handlers

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type categoryHandler struct {
	controller controllers.CategoryController
}

func NewCategoryHandler(controller controllers.CategoryController) CategoryHandler {
	return &categoryHandler{controller}
}

func (h *categoryHandler) RegisterRoutes(router *gin.Engine) {
	categories := router.Group("/categories").Use(middlewares.AuthMiddleware())
	{
		editCategories := categories.Use(middlewares.PermissionsMiddleware("Edit categorical data"))
		{
			// Создание категории с name в URL
			editCategories.POST("/:name", h.Create)

			// Получение категории по ID
			editCategories.GET("/:id", h.GetByID)

			// Обновление категории с ID и name в URL
			editCategories.PUT("/:id/:name", h.Update)

			// Удаление категории по ID
			editCategories.DELETE("/:id", h.Delete)
		}
		// Получение всех категорий
		categories.GET("/", h.GetAll).
			Use(middlewares.PermissionsMiddleware("See all data"))
	}
}

func (h *categoryHandler) Create(ctx *gin.Context) {
	// Получаем name из URL
	name := ctx.Param("name")

	if name == "" {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	createdCategory, err := h.controller.Create(name)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, createdCategory)
}

func (h *categoryHandler) GetAll(ctx *gin.Context) {
	categories, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, categories)
}

func (h *categoryHandler) GetByID(ctx *gin.Context) {
	// Получаем ID из URL
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	category, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, category)
}

func (h *categoryHandler) Update(ctx *gin.Context) {
	// Получаем ID и name из URL
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	name := ctx.Param("name")
	if name == "" {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	updatedCategory, err := h.controller.Update(id, name)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, updatedCategory)
}

func (h *categoryHandler) Delete(ctx *gin.Context) {
	// Получаем ID из URL
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	err := h.controller.Delete(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
