package order

import (
	"MireaPR4/database/models"
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/order/dto"
	"MireaPR4/http/middlewares"
	"context"
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
	Update(ctx *gin.Context)
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
		orders.GET("/all/", h.GetAll).
			Use(middlewares.PermissionsMiddleware("See all data"))
		orders.GET("/:id", h.GetByID)
		orders.PUT("/:id", h.Update)
		orders.DELETE(
			"/:id",
			middlewares.PermissionsMiddleware("View and delete Orders"),
			h.Delete,
		)
	}
}

// Create Хендлеры
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

func (h *orderHandler) GetAll(ctx *gin.Context) {
	ctxT, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	response, err := h.controller.GetAll(&ctxT)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

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

func (h *orderHandler) Update(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}
	var requestData any
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	response, err := h.controller.Update(id, requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

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
