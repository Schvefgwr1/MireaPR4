package payment

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/payment/dto"
	"MireaPR4/http/middlewares"
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
	payments := router.Group("/payments").Use(middlewares.AuthMiddleware())
	{
		// Создание нового платежа
		payments.POST("/", h.Create)

		// Получение всех платежей
		payments.GET("/", middlewares.PermissionsMiddleware("See all data"), h.GetAll)

		// Получение платежа по ID
		payments.GET("/:id", middlewares.PermissionsMiddleware("Edit payments"), h.GetByID)

		// Обновление платежа (сумма или статус)
		payments.PUT("/:id", middlewares.PermissionsMiddleware("Edit payments"), h.Update)

		// Удаление платежа по ID
		payments.DELETE("/:id", middlewares.PermissionsMiddleware("Delete some data"), h.Delete)
	}
}

// Create Создание нового платежа
// @Summary Создание нового платежа
// @Description Создаёт новый платеж, получая данные в формате JSON с деталями платежа
// @Tags /payments
// @Accept json
// @Produce json
// @Param payment body dto.CreatePaymentDTO true "DTO для создания платежа"
// @Success 201 {object} map[string]interface{} "Платеж успешно создан"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /payments [post]
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

// GetAll Получение всех платежей
// @Summary Получение всех платежей
// @Description Получение всех платежей
// @Tags /payments
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Платежи успешно получены"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /payments [get]
func (h *paymentHandler) GetAll(ctx *gin.Context) {
	payments, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, payments)
}

// GetByID Получение платежа по ID
// @Summary Получение платежа по ID
// @Description Получение платежа по ID
// @Tags /payments
// @Accept json
// @Produce json
// @Param id path int true "ID платежа"
// @Success 200 {object} map[string]interface{} "Платеж успешно получен"
// @Failure 400 {object} map[string]interface{} "Неверный ID платежа"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 404 {object} map[string]interface{} "Платеж не найден"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /payments/{id} [get]
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

// Update Обновление платежа
// @Summary Обновление платежа
// @Description Обновляет информацию о платеже (сумма, статус)
// @Tags /payments
// @Accept json
// @Produce json
// @Param id path int true "ID платежа"
// @Param payment body dto.UpdatePaymentDTO true "DTO для обновления платежа"
// @Success 200 {object} map[string]interface{} "Платеж успешно обновлен"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 404 {object} map[string]interface{} "Платеж не найден"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /payments/{id} [put]
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

// Delete Удаление платежа по ID
// @Summary Удаление платежа по ID
// @Description Удаляет платеж по ID
// @Tags /payments
// @Accept json
// @Produce json
// @Param id path int true "ID платежа"
// @Success 200 {object} map[string]interface{} "Платеж успешно удален"
// @Failure 400 {object} map[string]interface{} "Неверный ID платежа"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /payments/{id} [delete]
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
