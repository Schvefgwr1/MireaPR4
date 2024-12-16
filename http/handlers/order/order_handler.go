package order

import (
	"MireaPR4/database/models"
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/order/dto"
	"MireaPR4/http/middlewares"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type OrderHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetAllPaginated(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type orderHandler struct {
	controller controllers.OrderController
}

func NewOrderHandler(controller controllers.OrderController) OrderHandler {
	return &orderHandler{controller}
}

func validateCreateOrderDTO(data any) (*dto.CreateOrderDTO, error) {
	// Утверждаем, что объект имеет тип CreateOrderDTO
	orderDTO, ok := data.(*dto.CreateOrderDTO)
	if !ok {
		return nil, errors.New("invalid type: expected CreateOrderDTO")
	}

	if orderDTO.Items == nil || len(orderDTO.Items) == 0 {
		return nil, errors.New("field 'product_ids' is required")
	}

	return orderDTO, nil
}

// RegisterRoutes Регистрация
func (h *orderHandler) RegisterRoutes(router *gin.Engine) {
	orders := router.Group("/orders")
	orders.Use(middlewares.AuthMiddleware())
	{
		orders.POST("/", h.Create)
		orders.GET(
			"/",
			middlewares.PermissionsMiddleware("View and delete Orders"),
			h.GetAllPaginated,
		)
		orders.GET(
			"/all/",
			middlewares.TimeoutMiddleware(2*time.Second),
			middlewares.PermissionsMiddleware("See all data"),
			h.GetAll,
		)
		orders.GET("/:id", h.GetByID)
		orders.DELETE(
			"/:id",
			middlewares.PermissionsMiddleware("View and delete Orders"),
			h.Delete,
		)
	}
}

// Create Создание заказа
// @Summary Создание нового заказа
// @Description Создаёт новый заказ, получая данные в формате JSON с деталями заказа
// @Tags /orders
// @Accept json
// @Produce json
// @Param order body dto.CreateOrderDTO true "DTO для создания заказа"
// @Success 201 {object} map[string]interface{} "Заказ успешно создан"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /orders [post]
func (h *orderHandler) Create(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	var requestData dto.CreateOrderDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	validDTO, errDTO := validateCreateOrderDTO(&requestData)
	if errDTO != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, errDTO.Error())
		return
	}

	response, err := h.controller.Create(validDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// GetAll Получение всех заказов
// @Summary Получение всех заказов
// @Description Получение всех заказов с пагинацией
// @Tags /orders
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество на странице" default(10)
// @Success 200 {object} map[string]interface{} "Заказы успешно получены"
// @Failure 400 {object} map[string]interface{} "Неверные параметры 'page' или 'limit'"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /orders [get]
func (h *orderHandler) GetAll(ctx *gin.Context) {
	response, err := h.controller.GetAll(ctx.Request.Context())
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetAllPaginated Получение всех заказов с пагинацией
// @Summary Получение всех заказов с пагинацией для пользователя
// @Description Получение заказов с пагинацией и фильтрацией по пользователю
// @Tags /orders
// @Accept json
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Количество на странице" default(10)
// @Param userID query int false "ID пользователя" default(-1)
// @Success 200 {object} map[string]interface{} "Заказы успешно получены"
// @Failure 400 {object} map[string]interface{} "Неверные параметры 'page' или 'limit'"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /orders/all/ [get]
func (h *orderHandler) GetAllPaginated(ctx *gin.Context) {
	page, err := default_functions.ParseQueryParam(ctx, "page", 1)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid 'page' parameter")
		return
	}

	limit, err := default_functions.ParseQueryParam(ctx, "limit", 10)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid 'limit' parameter")
		return
	}

	offset := (page - 1) * limit
	var orders []models.Order
	var total int64

	userID, e := default_functions.ParseQueryParam(ctx, "userID", -1)
	if e != nil || userID == -1 {
		orders, total, err = h.controller.GetAllPaginated(offset, limit, nil)
	} else {
		orders, total, err = h.controller.GetAllPaginated(offset, limit, &userID)
	}
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"data":       orders,
		"total":      total,
		"page":       page,
		"page_count": (total + int64(limit) - 1) / int64(limit),
	})
}

// GetByID Получение заказа по ID
// @Summary Получение заказа по ID
// @Description Получение заказа по ID
// @Tags /orders
// @Accept json
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} models.Order "Заказ успешно получен"
// @Failure 400 {object} map[string]interface{} "Неверный ID заказа"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 404 {object} map[string]interface{} "Заказ не найден"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /orders/{id} [get]
func (h *orderHandler) GetByID(ctx *gin.Context) {
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

// Delete Удаление заказа по ID
// @Summary Удаление заказа по ID
// @Description Удаляет заказ по ID
// @Tags /orders
// @Accept json
// @Produce json
// @Param id path int true "ID заказа"
// @Success 200 {object} map[string]interface{} "Заказ успешно удален"
// @Failure 400 {object} map[string]interface{} "Неверный ID заказа"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /orders/{id} [delete]
func (h *orderHandler) Delete(ctx *gin.Context) {
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
		"message": "Order deleted successfully",
	})
}
