package payment

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/payment/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PaymentHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type paymentHandler struct {
	controller controllers.PaymentController
}

func NewPaymentHandler(controller controllers.PaymentController) PaymentHandler {
	return &paymentHandler{controller}
}

func (h *paymentHandler) RegisterRoutes(router *gin.Engine) {
	payments := router.Group("/payments")
	{
		// Создание нового платежа
		payments.POST("/", h.Create)

		// Получение всех платежей
		payments.GET("/", h.GetAll)

		// Получение платежа по ID
		payments.GET("/:id", h.GetByID)

		// Обновление платежа (сумма или статус)
		payments.PUT("/:id", h.Update)

		// Удаление платежа по ID
		payments.DELETE("/:id", h.Delete)
	}
}

func (h *paymentHandler) Create(ctx *gin.Context) {
	var paymentDTO dto.CreatePaymentDTO
	if err := ctx.ShouldBindJSON(&paymentDTO); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	createdPayment, err := h.controller.Create(&paymentDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, createdPayment)
}

func (h *paymentHandler) GetAll(ctx *gin.Context) {
	payments, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, payments)
}

func (h *paymentHandler) GetByID(ctx *gin.Context) {
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	payment, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, payment)
}

func (h *paymentHandler) Update(ctx *gin.Context) {
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	var updateDTO dto.UpdatePaymentDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	updatedPayment, err := h.controller.Update(id, &updateDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, updatedPayment)
}

func (h *paymentHandler) Delete(ctx *gin.Context) {
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	err := h.controller.Delete(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
