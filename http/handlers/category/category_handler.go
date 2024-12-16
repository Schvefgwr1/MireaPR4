package handlers

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler interface {
	Create(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type categoryHandler struct {
	controller controllers.CategoryController
}

func NewCategoryHandler(controller controllers.CategoryController) CategoryHandler {
	return &categoryHandler{controller}
}

func (h *categoryHandler) RegisterRoutes(router *gin.Engine) {
	categories := router.Group("/categories").Use(middlewares.AuthMiddleware())
	{
		editCategories := categories.Use(middlewares.PermissionsMiddleware("Edit categorical data"))
		{
			// Создание категории с name в URL
			editCategories.POST("/:name", h.Create)

			// Получение категории по ID
			editCategories.GET("/:id", h.GetByID)

			// Обновление категории с ID и name в URL
			editCategories.PUT("/:id", h.Update)

			// Удаление категории по ID
			editCategories.DELETE("/:id", h.Delete)
		}
		// Получение всех категорий
		categories.GET("/", middlewares.PermissionsMiddleware("See all data"), h.GetAll)
	}
}

// Create создает новую категорию
// @Summary Создание категории
// @Description Создает новую категорию на основе имени, переданного в параметрах URL
// @Tags /categories
// @Accept json
// @Produce json
// @Param name path string true "Имя категории"
// @Success 201 {object} map[string]interface{} "Успешно добавлена новая категория"
// @Failure 400 {object} map[string]interface{} "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /categories/{name} [post]
func (h *categoryHandler) Create(ctx *gin.Context) {
	// Получаем name из URL
	name := ctx.Param("name")

	if name == "" {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	createdCategory, err := h.controller.Create(name)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusCreated, createdCategory)
}

// GetAll возвращает все категории
// @Summary Получение всех категорий
// @Description Возвращает список всех существующих категорий
// @Tags /categories
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{} "Список категорий"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /categories [get]
func (h *categoryHandler) GetAll(ctx *gin.Context) {
	categories, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, categories)
}

// GetByID возвращает категорию по ID
// @Summary Получение категории по ID
// @Description Возвращает информацию о категории на основе её ID
// @Tags /categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]interface{} "Информация о категории"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 404 {object} map[string]interface{} "Категория не найдена"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /categories/{id} [get]
func (h *categoryHandler) GetByID(ctx *gin.Context) {
	// Получаем ID из URL
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		return
	}

	category, err := h.controller.GetByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, category)
}

// Update обновляет категорию
// @Summary Обновление категории
// @Description Обновляет информацию о категории на основе её ID и имени, переданного в параметрах
// @Tags /categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Param name query string true "Новое имя категории"
// @Success 200 {object} map[string]interface{} "Обновленная информация о категории"
// @Failure 400 {object} map[string]interface{} "Ошибка валидации входных данных"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /categories/{id} [put]
func (h *categoryHandler) Update(ctx *gin.Context) {
	// Получаем ID и name из URL
	strID := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, ctx)
	if !valid {
		return
	}

	name := ctx.Query("name")
	if name == "" {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "Name is required")
		return
	}

	updatedCategory, err := h.controller.Update(id, name)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	default_functions.RespondWithSuccess(ctx, http.StatusOK, updatedCategory)
}

// Delete удаляет категорию
// @Summary Удаление категории
// @Description Удаляет категорию на основе её ID
// @Tags /categories
// @Accept json
// @Produce json
// @Param id path int true "ID категории"
// @Success 200 {object} map[string]interface{} "Сообщение об успешном удалении"
// @Failure 400 {object} map[string]interface{} "Неверный формат ID"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Ошибка сервера"
// @Security BearerAuth
// @Router /categories/{id} [delete]
func (h *categoryHandler) Delete(ctx *gin.Context) {
	// Получаем ID из URL
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

	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
