package order

import (
	"MireaPR4/controllers"
	"MireaPR4/handlers/order/dto"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	controller controllers.OrderController
}

func NewOrderHandler(controller controllers.OrderController) *OrderHandler {
	return &OrderHandler{controller}
}

// Helper для обработки ошибок
func (h *OrderHandler) respondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"status": "error",
		"error":  message,
	})
}

// Helper для успешных ответов
func (h *OrderHandler) respondWithSuccess(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}

// Helper для проверки Content-Type
func (h *OrderHandler) validateJSON(ctx *gin.Context) bool {
	if ctx.GetHeader("Content-Type") != "application/json" {
		h.respondWithError(ctx, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return false
	}
	return true
}

// Helper для преобразования id из строки в int
func (h *OrderHandler) convertID(ctx *gin.Context) (int, bool) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.respondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return 0, false
	}
	return idInt, true
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
func (h *OrderHandler) RegisterRoutes(router *gin.Engine) {
	orders := router.Group("/orders")
	{
		orders.POST("/", h.Create)
		orders.GET("/", h.GetAll)
		orders.GET("/:id", h.GetByID)
		orders.PUT("/:id", h.Update)
		orders.DELETE("/:id", h.Delete)
	}
}

// Create Хендлеры
func (h *OrderHandler) Create(ctx *gin.Context) {
	if !h.validateJSON(ctx) {
		return
	}

	var requestData dto.CreateOrderDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		h.respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	validDTO, errDTO := validateCreateOrderDTO(&requestData)
	if errDTO != nil {
		h.respondWithError(ctx, http.StatusBadRequest, errDTO.Error())
		return
	}

	response, err := h.controller.Create(validDTO)
	if err != nil {
		h.respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondWithSuccess(ctx, http.StatusCreated, response)
}

func (h *OrderHandler) GetAll(ctx *gin.Context) {
	response, err := h.controller.GetAll()
	if err != nil {
		h.respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondWithSuccess(ctx, http.StatusOK, response)
}

func (h *OrderHandler) GetByID(ctx *gin.Context) {
	id, valid := h.convertID(ctx)
	if !valid {
		return
	}
	response, err := h.controller.GetByID(id)
	if err != nil {
		h.respondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}
	h.respondWithSuccess(ctx, http.StatusOK, response)
}

func (h *OrderHandler) Update(ctx *gin.Context) {
	if !h.validateJSON(ctx) {
		return
	}
	id, valid := h.convertID(ctx)
	if !valid {
		return
	}
	var requestData any
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		h.respondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.controller.Update(id, requestData)
	if err != nil {
		h.respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondWithSuccess(ctx, http.StatusOK, response)
}

func (h *OrderHandler) Delete(ctx *gin.Context) {
	id, valid := h.convertID(ctx)
	if !valid {
		return
	}
	err := h.controller.Delete(id)
	if err != nil {
		h.respondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	h.respondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
