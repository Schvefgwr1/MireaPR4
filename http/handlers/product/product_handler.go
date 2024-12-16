package product

import (
	"MireaPR4/database/models"
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/product/dto"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProductHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetAllPaginated(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type productHandler struct {
	controller controllers.ProductController
}

func NewProductHandler(controller controllers.ProductController) ProductHandler {
	return &productHandler{controller}
}

func (h *productHandler) RegisterRoutes(router *gin.Engine) {
	products := router.Group("/products").Use(middlewares.AuthMiddleware())
	{
		products.POST("/", middlewares.PermissionsMiddleware("Edit Products"), h.Create)
		products.GET("/", h.GetAllPaginated)
		products.GET(
			"/all/",
			middlewares.TimeoutMiddleware(2*time.Second),
			middlewares.PermissionsMiddleware("See all data"),
			h.GetAll,
		)
		products.GET("/:id", h.GetByID)
		products.PUT("/:id", middlewares.PermissionsMiddleware("Edit Products"), h.Update)
		products.DELETE("/:id", middlewares.PermissionsMiddleware("Delete some data"), h.Delete)
	}
}

// Create Создание нового продукта
// @Summary Создание нового продукта
// @Description Создает новый продукт, получая данные в формате JSON с деталями продукта
// @Tags /products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductDTO true "DTO для создания продукта"
// @Success 201 {object} map[string]interface{} "Продукт успешно создан"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /products [post]
func (h *productHandler) Create(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	var requestData dto.CreateProductDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.controller.Create(&requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// GetAll Получение всех продуктов
// @Summary Получение всех продуктов
// @Description Получение всех продуктов
// @Tags /products
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Продукты успешно получены"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Failure 504 {object} map[string]interface{} "Ошибка в связи с таймаутом"
// @Security BearerAuth
// @Router /products/all/ [get]
func (h *productHandler) GetAll(ctx *gin.Context) {
	response, err := h.controller.GetAll(ctx.Request.Context())
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetAllPaginated Получение продуктов с пагинацией
// @Summary Получение продуктов с пагинацией
// @Description Получение продуктов с пагинацией
// @Tags /products
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество элементов на странице" default(10)
// @Param categoryID query int false "ID категории для фильтрации продуктов" default(-1)
// @Success 200 {object} map[string]interface{} "Продукты успешно получены"
// @Failure 400 {object} map[string]interface{} "Неверные данные запроса"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /products [get]
func (h *productHandler) GetAllPaginated(ctx *gin.Context) {
	page, err := default_functions.ParseQueryParam(ctx, "page", 1)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "invalid 'page' parameter")
		return
	}

	limit, err := default_functions.ParseQueryParam(ctx, "limit", 10)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "invalid 'limit' parameter")
		return
	}

	var products []models.Product

	var total int64
	categoryID, e := default_functions.ParseQueryParam(ctx, "categoryID", -1)
	if e != nil || categoryID == -1 {
		products, total, err = h.controller.GetAllWithPagination(page, limit, nil)

	} else {
		products, total, err = h.controller.GetAllWithPagination(page, limit, &categoryID)
	}

	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"data":       products,
		"total":      total,
		"page":       page,
		"page_count": (total + int64(limit) - 1) / int64(limit),
	})
}

// GetByID Получение продукта по ID
// @Summary Получение продукта по ID
// @Description Получение продукта по ID
// @Tags /products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]interface{} "Продукт успешно получен"
// @Failure 400 {object} map[string]interface{} "Неверный ID продукта"
// @Failure 404 {object} map[string]interface{} "Продукт не найден"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /products/{id} [get]
func (h *productHandler) GetByID(ctx *gin.Context) {
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "error of param in route")
		return
	}

	response, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// Update Обновление продукта
// @Summary Обновление продукта
// @Description Обновляет информацию о продукте
// @Tags /products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Param product body dto.UpdateProductDTO true "DTO для обновления продукта"
// @Success 200 {object} map[string]interface{} "Продукт успешно обновлен"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 404 {object} map[string]interface{} "Продукт не найден"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /products/{id} [put]
func (h *productHandler) Update(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	var requestData dto.UpdateProductDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.controller.Update(id, &requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// Delete Удаление продукта по ID
// @Summary Удаление продукта по ID
// @Description Удаляет продукт по ID
// @Tags /products
// @Accept json
// @Produce json
// @Param id path int true "ID продукта"
// @Success 200 {object} map[string]interface{} "Продукт успешно удален"
// @Failure 400 {object} map[string]interface{} "Неверный ID продукта"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /products/{id} [delete]
func (h *productHandler) Delete(ctx *gin.Context) {
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

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
