package handlers

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	"MireaPR4/http/handlers/register/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterHandler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RegisterRoutes(router *gin.Engine)
}

type registerHandler struct {
	controller controllers.RegisterController
}

func NewRegisterHandler(controller controllers.RegisterController) RegisterHandler {
	return &registerHandler{controller}
}

func (rh *registerHandler) RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/", rh.Login)
		auth.POST("/reg", rh.Register)
	}
}

// Register Регистрация нового пользователя
// @Summary Регистрация нового пользователя
// @Description Создаёт нового пользователя на основе данных из запроса
// @Tags /auth
// @Accept json
// @Produce json
// @Param user body dto.RegisterDTO true "DTO для регистрации пользователя"
// @Success 200 {object} map[string]interface{} "Пользователь успешно зарегистрирован"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 409 {object} map[string]interface{} "Пользователь уже создан"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Router /auth/reg [post]
func (rh *registerHandler) Register(c *gin.Context) {
	var input dto.RegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		default_functions.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := rh.controller.CreateUser(input.Username, input.Password, input.Email, 1)
	if err != nil {
		if err.Error() == "user already exist" {
			default_functions.RespondWithError(c, http.StatusConflict, err.Error())
			return
		}
		default_functions.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(c, http.StatusOK, gin.H{
		"message":  "User registered successfully",
		"login":    newUser.Username,
		"password": input.Password,
	})
}

// Login Аутентификация пользователя
// @Summary Аутентификация пользователя
// @Description Выполняет аутентификацию пользователя и возвращает токен
// @Tags /auth
// @Accept json
// @Produce json
// @Param login body dto.LoginDTO true "DTO для аутентификации пользователя"
// @Success 200 {object} map[string]interface{} "Токен успешно получен"
// @Failure 400 {object} map[string]interface{} "Неверные данные ввода"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Router /auth [post]
func (rh *registerHandler) Login(c *gin.Context) {
	var input dto.LoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := rh.controller.AuthenticateUser(input.Login, input.Password)
	if err != nil {
		default_functions.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	default_functions.RespondWithSuccess(c, http.StatusOK, gin.H{
		"token": token,
	})

}
