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
		employees.GET("/", middlewares.PermissionsMiddleware("See all data"), h.GetAll)
		employees.GET("/:id", h.GetByID)
		employees.PUT("/:id", h.Update)
		employees.DELETE("/:id", middlewares.PermissionsMiddleware("Delete some data"), h.Delete)
	}
}

// Create создает нового сотрудника
// @Summary Создание сотрудника
// @Description Создает нового сотрудника на основе предоставленных данных
// @Tags /employees
// @Accept json
// @Produce json
// @Param employee body dto.CreateEmployeeDTO true "Данные для создания сотрудника"
// @Success 201 {object} map[string]interface{} "Успешно добавлен новый сотрудник"
// @Failure 400 {object} map[string]interface{} "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /employees [post]
func (h *employeeHandler) Create(ctx *gin.Context) {
	var employeeDTO dto.CreateEmployeeDTO
	if err := ctx.ShouldBindJSON(&employeeDTO); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Invalid input data")
		return
	}

	createdEmployee, err := h.controller.Create(&employeeDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, createdEmployee)
}

// GetAll возвращает всех сотрудников
// @Summary Получение всех сотрудников
// @Description Возвращает список всех существующих сотрудников
// @Tags /employees
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Список сотрудников"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /employees [get]
func (h *employeeHandler) GetAll(ctx *gin.Context) {
	employees, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, employees)
}

// GetByID возвращает сотрудника по ID
// @Summary Получение сотрудника по ID
// @Description Возвращает информацию о сотруднике на основе его ID
// @Tags /employees
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} map[string]interface{} "Информация о сотруднике"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 404 {object} map[string]interface{} "Сотрудник не найден"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /employees/{id} [get]
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

// Update обновляет данные сотрудника
// @Summary Обновление сотрудника
// @Description Обновляет информацию о сотруднике на основе его ID
// @Tags /employees
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Param employee body dto.UpdateEmployeeDTO true "Обновленные данные сотрудника"
// @Success 200 {object} map[string]interface{} "Обновленная информация о сотруднике"
// @Failure 400 {object} map[string]interface{} "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /employees/{id} [put]
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

// Delete удаляет сотрудника
// @Summary Удаление сотрудника
// @Description Удаляет сотрудника на основе его ID
// @Tags /employees
// @Accept json
// @Produce json
// @Param id path int true "ID сотрудника"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /employees/{id} [delete]
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
