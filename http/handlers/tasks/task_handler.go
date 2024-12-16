package taskspackage

import (
	"MireaPR4/http/default_functions"
	"MireaPR4/http/middlewares"
	"MireaPR4/tasks"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	products := router.Group("/tasks").Use(
		middlewares.AuthMiddleware(),
		middlewares.PermissionsMiddleware("Run tasks"),
	)
	{
		products.POST("/:id", CreateTask)
		products.GET("/:id", GetTask)
	}
}

// CreateTask Создание задачи
// @Summary Создание новой задачи
// @Description Создаёт задачу на основе переданного ID и запускает её
// @Tags /tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 201 {object} map[string]interface{} "Задача успешно создана"
// @Failure 400 {object} map[string]interface{} "Неверный запрос"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /tasks/{id} [post]
func CreateTask(c *gin.Context) {
	strID := c.Param("id")
	id, valid := default_functions.ConvertStrToIntParam(strID, c)
	if !valid {
		return
	}
	taskID, err := tasks.CreateTask(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tasks.RunTask(taskID)
	c.JSON(201, gin.H{"task_id": taskID})
}

// GetTask Получение задачи по ID
// @Summary Получение информации о задаче по ID
// @Description Возвращает задачу с указанным ID
// @Tags /tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Success 200 {object} map[string]interface{} "Задача успешно найдена"
// @Failure 401 {object} map[string]interface{} "Ошибка аутентификации"
// @Failure 403 {object} map[string]interface{} "Ошибка прав доступа"
// @Failure 404 {object} map[string]interface{} "Задача не найдена"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Security BearerAuth
// @Router /tasks/{id} [get]
func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task := tasks.GetTask(taskID)
	if task == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}
