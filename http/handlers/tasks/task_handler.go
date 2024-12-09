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

func GetTask(c *gin.Context) {
	taskID := c.Param("id")
	task := tasks.GetTask(taskID)
	if task == nil {
		c.JSON(404, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(200, task)
}
