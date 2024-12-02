package product

import (
	"MireaPR4/database/models"
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/product/dto"
	"MireaPR4/http/middlewares"
	"context"
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
		products.POST("/", h.Create)
		products.GET("/", h.GetAllPaginated)
		products.GET("/all/", h.GetAll).
			Use(middlewares.PermissionsMiddleware("See all data"))
		products.GET("/:id", h.GetByID)
		products.PUT("/:id", h.Update)
		products.DELETE("/:id", h.Delete)
	}
}

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

func (h *productHandler) GetAll(ctx *gin.Context) {
	ctxT, cancel := context.WithTimeout(ctx.Request.Context(), 2*time.Second)
	defer cancel()

	response, err := h.controller.GetAll(&ctxT)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

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

	categoryID, e := default_functions.ParseQueryParam(ctx, "categoryID", -1)
	if e != nil || categoryID == -1 {
		products, err = h.controller.GetAllWithPagination(page, limit, nil)
	} else {
		products, err = h.controller.GetAllWithPagination(page, limit, &categoryID)
	}

	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, products)
}

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
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

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
