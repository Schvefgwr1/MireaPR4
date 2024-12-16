package shipment

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/shipment/dto"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShipmentHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type shipmentHandler struct {
	controller controllers.ShipmentController
}

func NewShipmentHandler(controller controllers.ShipmentController) ShipmentHandler {
	return &shipmentHandler{controller}
}

func (h *shipmentHandler) RegisterRoutes(router *gin.Engine) {
	shipments := router.Group("/shipments").Use(middlewares.AuthMiddleware())
	{
		shipments.POST("/", h.Create)
		shipments.GET("/", h.GetAll)
		shipments.GET("/:id", h.GetByID)
		shipments.PUT("/:id", h.Update)
		shipments.DELETE("/:id", h.Delete)
	}
}

func (h *shipmentHandler) Create(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	var requestData dto.CreateShipmentDTO
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

func (h *shipmentHandler) GetAll(ctx *gin.Context) {
	response, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

func (h *shipmentHandler) GetByID(ctx *gin.Context) {
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

func (h *shipmentHandler) Update(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	var requestData dto.UpdateShipmentDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.controller.Update(id, &requestData)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

func (h *shipmentHandler) Delete(ctx *gin.Context) {
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
		"message": "Shipment deleted successfully",
	})
}
