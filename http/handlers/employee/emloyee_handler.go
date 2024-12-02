package employee

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/employee/dto"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type employeeHandler struct {
	controller controllers.EmployeeController
}

func NewEmployeeHandler(controller controllers.EmployeeController) EmployeeHandler {
	return &employeeHandler{controller}
}

func (h *employeeHandler) RegisterRoutes(router *gin.Engine) {
	employees := router.Group("/employees").Use(middlewares.AuthMiddleware())
	{
		employees.POST("/", h.Create)
		employees.GET("/", h.GetAll).
			Use(middlewares.PermissionsMiddleware("See all data"))
		employees.GET("/:id", h.GetByID)
		employees.PUT("/:id", h.Update)
		employees.DELETE("/:id", h.Delete).
			Use(middlewares.PermissionsMiddleware("Edit categorical data"))
	}
}

func (h *employeeHandler) Create(ctx *gin.Context) {
	var employeeDTO dto.CreateEmployeeDTO
	if err := ctx.ShouldBindJSON(&employeeDTO); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	createdEmployee, err := h.controller.Create(&employeeDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, createdEmployee)
}

func (h *employeeHandler) GetAll(ctx *gin.Context) {
	employees, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, employees)
}

func (h *employeeHandler) GetByID(ctx *gin.Context) {
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	employee, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, employee)
}

func (h *employeeHandler) Update(ctx *gin.Context) {
	// Получаем ID из параметров маршрута
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	// Привязываем данные из тела запроса к DTO
	var updateDTO dto.UpdateEmployeeDTO
	if err := ctx.ShouldBindJSON(&updateDTO); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	// Передаем данные в контроллер
	updatedEmployee, err := h.controller.Update(id, &updateDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, updatedEmployee)
}

func (h *employeeHandler) Delete(ctx *gin.Context) {
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

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Employee deleted successfully"})
}
