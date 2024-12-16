package role

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/role/dto"
	"MireaPR4/http/middlewares"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RoleHandler interface {
	Create(ctx *gin.Context)
	GetByID(ctx *gin.Context)
	GetByName(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type roleHandler struct {
	controller controllers.RoleController
}

func NewRoleHandler(cont controllers.RoleController) RoleHandler {
	return &roleHandler{cont}
}

func (h *roleHandler) RegisterRoutes(router *gin.Engine) {
	roles := router.Group("/roles").Use(middlewares.AuthMiddleware())

	modifyRoles := roles.Use(middlewares.PermissionsMiddleware("Modify roles"))
	{
		modifyRoles.POST("/", h.Create)
		modifyRoles.GET("/", h.GetAll)
		modifyRoles.GET("/:id", h.GetByID)
		modifyRoles.PUT("/:id", h.Update)
		modifyRoles.DELETE("/:id", h.Delete)
	}

	roles.GET("/name/:name", h.GetByName)
}

func validateCreateRoleDTO(data any) (*dto.CreateRoleDTO, error) {
	roleDTO, ok := data.(*dto.CreateRoleDTO)
	if !ok {
		return nil, errors.New("invalid type: expected CreateRoleDTO")
	}

	if roleDTO.Permissions == nil || len(roleDTO.Permissions) == 0 {
		return nil, errors.New("field 'permissions_ids' is required")
	}

	return roleDTO, nil
}

// Create Создание новой роли
// @Summary Создание новой роли
// @Description Создаёт новую роль с указанными правами
// @Tags /roles
// @Accept json
// @Produce json
// @Param role body dto.CreateRoleDTO true "DTO для создания роли"
// @Success 201 {object} map[string]interface{} "Роль успешно создана"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles [post]
func (h *roleHandler) Create(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	var requestData dto.CreateRoleDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	validDTO, errDTO := validateCreateRoleDTO(&requestData)
	if errDTO != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, errDTO.Error())
		return
	}

	response, err := h.controller.Create(validDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusCreated, response)
}

// GetAll Получение всех ролей
// @Summary Получение всех ролей
// @Description Возвращает список всех ролей
// @Tags /roles
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Роли успешно получены"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles [get]
func (h *roleHandler) GetAll(ctx *gin.Context) {
	response, err := h.controller.GetAll()
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetByID Получение роли по ID
// @Summary Получение роли по ID
// @Description Возвращает роль по ID
// @Tags /roles
// @Accept json
// @Produce json
// @Param id path int true "ID роли"
// @Success 200 {object} map[string]interface{} "Роль успешно получена"
// @Failure 400 {object} map[string]interface{} "Неверный ID роли"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 404 {object} map[string]interface{} "Роль не найдена"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles/{id} [get]
func (h *roleHandler) GetByID(ctx *gin.Context) {
	strParam := ctx.Param("id")

	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "error of param in route")
		return
	}
	response, err := h.controller.GetRoleByID(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// GetByName Получение роли по имени
// @Summary Получение роли по имени
// @Description Возвращает роль по имени
// @Tags /roles
// @Accept json
// @Produce json
// @Param name path string true "Имя роли"
// @Success 200 {object} map[string]interface{} "Роль успешно получена"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 404 {object} map[string]interface{} "Роль не найдена"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles/name/{name} [get]
func (h *roleHandler) GetByName(ctx *gin.Context) {
	strParam := ctx.Param("name")

	response, err := h.controller.GetRoleByName(strParam)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

func validateUpdateRoleDTO(data any) (*dto.UpdateRoleDTO, error) {
	roleDTO, ok := data.(*dto.UpdateRoleDTO)
	if !ok {
		return nil, errors.New("invalid type: expected UpdateRoleDTO")
	}

	if roleDTO.Permissions == nil || len(roleDTO.Permissions) == 0 {
		return nil, errors.New("field 'permissions_ids' is required")
	}

	return roleDTO, nil
}

// Update Обновление роли
// @Summary Обновление роли
// @Description Обновляет информацию о роли (например, права доступа)
// @Tags /roles
// @Accept json
// @Produce json
// @Param id path int true "ID роли"
// @Param role body dto.UpdateRoleDTO true "DTO для обновления роли"
// @Success 200 {object} map[string]interface{} "Роль успешно обновлена"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles/{id} [put]
func (h *roleHandler) Update(ctx *gin.Context) {
	if !default_functions.ValidateJSON(ctx) {
		return
	}

	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "error of param in route")
		return
	}
	var requestData dto.UpdateRoleDTO
	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	validDTO, errDTO := validateUpdateRoleDTO(&requestData)
	if errDTO != nil {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, errDTO.Error())
		return
	}

	response, err := h.controller.Update(id, validDTO)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, response)
}

// Delete Удаление роли по ID
// @Summary Удаление роли по ID
// @Description Удаляет роль по ID
// @Tags /roles
// @Accept json
// @Produce json
// @Param id path int true "ID роли"
// @Success 200 {object} map[string]interface{} "Роль успешно удалена"
// @Failure 400 {object} map[string]interface{} "Неверный ID роли"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /roles/{id} [delete]
func (h *roleHandler) Delete(ctx *gin.Context) {
	strParam := ctx.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strParam, ctx)
	if !valid {
		default_functions.RespondWithError(ctx, http.StatusBadRequest, "error of param in route")
		return
	}
	err := h.controller.Delete(id)
	if err != nil {
		default_functions.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	default_functions.RespondWithSuccess(ctx, http.StatusOK, gin.H{
		"message": "Role deleted successfully",
	})
}
