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
		addresses.GET("/", h.GetAll).
			Use(middlewares.PermissionsMiddleware("See all data"))
		addresses.GET("/:id", h.GetByID)
		addresses.PUT("/:id", h.Update)
		addresses.DELETE("/:id", h.Delete)
	}
}

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

func (h *addressHandler) GetAll(ctx *gin.Context) {
	addresses, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, addresses)
}

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
