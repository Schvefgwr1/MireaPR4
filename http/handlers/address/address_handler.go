package address

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/address/dto"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AddressHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type addressHandler struct {
	controller controllers.AddressController
}

func NewAddressHandler(controller controllers.AddressController) AddressHandler {
	return &addressHandler{controller}
}

func (h *addressHandler) RegisterRoutes(router *gin.Engine) {
	addresses := router.Group("/addresses").Use(middlewares.AuthMiddleware())
	{
		addresses.POST("/", h.Create)
		addresses.GET("/", middlewares.PermissionsMiddleware("See all data"), h.GetAll)
		addresses.GET("/:id", h.GetByID)
		addresses.PUT("/:id", h.Update)
		addresses.DELETE("/:id", middlewares.PermissionsMiddleware("Delete some data"), h.Delete)
	}
}

// Create CreateAddress создает новый адрес
// @Summary Создание адреса
// @Description Создает новый адрес на основе предоставленных данных
// @Tags /addresses
// @Accept json
// @Produce json
// @Param address body dto.CreateAddressDTO true "Данные для создания адреса"
// @Success 201 {object} map[string]interface{} "Успешно добавлен новый адрес"
// @Failure 400 {object} map[string]interface{} "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /addresses [post]
func (h *addressHandler) Create(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	var requestData dto.CreateAddressDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	address, err := h.controller.Create(&requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, address)
}

// GetAll возвращает все адреса
// @Summary Получение всех адресов
// @Description Возвращает список всех адресов
// @Tags /addresses
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Список адресов"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /addresses [get]
func (h *addressHandler) GetAll(ctx *gin.Context) {
	addresses, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, addresses)
}

// GetByID возвращает адрес по идентификатору
// @Summary Получение адреса по ID
// @Description Возвращает адрес на основе его идентификатора
// @Tags /addresses
// @Accept json
// @Produce json
// @Param id path int true "ID адреса"
// @Success 200 {object} map[string]interface{} "Информация об адресе"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 404 {object} map[string]interface{} "Адрес не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /addresses/{id} [get]
func (h *addressHandler) GetByID(ctx *gin.Context) {
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return
	}

	address, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, address)
}

// Update обновляет информацию об адресе
// @Summary Обновление адреса
// @Description Обновляет данные существующего адреса
// @Tags /addresses
// @Accept json
// @Produce json
// @Param id path int true "ID адреса"
// @Param address body dto.UpdateAddressDTO true "Данные для обновления адреса"
// @Success 200 {object} map[string]interface{} "Обновленные данные адреса"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID или данные для обновления"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /addresses/{id} [put]
func (h *addressHandler) Update(ctx *gin.Context) {
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	var requestData dto.UpdateAddressDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	address, err := h.controller.Update(id, &requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, address)
}

// Delete удаляет адрес по идентификатору
// @Summary Удаление адреса
// @Description Удаляет адрес на основе его идентификатора
// @Tags /addresses
// @Accept json
// @Produce json
// @Param id path int true "ID адреса"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /addresses/{id} [delete]
func (h *addressHandler) Delete(ctx *gin.Context) {
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

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
