package handlers

import (
	"MireaPR4/http/controllers"
	"MireaPR4/http/default_functions"
	dto2 "MireaPR4/http/handlers/register/dto"
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

func (rh *registerHandler) Register(c *gin.Context) {
	var input dto2.RegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		default_functions.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	newUser, err := rh.controller.CreateUser(input.Username, input.Password, input.Email, 1)
	if err != nil {
		default_functions.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	default_functions.RespondWithSuccess(c, http.StatusOK, gin.H{
		"message":  "User registered successfully",
		"login":    newUser.Username,
		"password": input.Password,
	})
}

func (rh *registerHandler) Login(c *gin.Context) {
	var input dto2.LoginDTO
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
