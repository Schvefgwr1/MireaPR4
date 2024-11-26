package middlewares

import (
	"MireaPR4/database/models"
	"MireaPR4/database/repositories"
	"MireaPR4/http/default_functions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var userRepo *repositories.UserRepository

func InitDB(UserRepo *repositories.UserRepository) {
	userRepo = UserRepo
}

func PermissionsMiddleware(permissions string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, haveFind := c.Get("userID")
		if !haveFind {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Error of permission check"})
			c.Abort()
			return
		}

		userIDConvert, valid := default_functions.ConvertStrToIntParam(fmt.Sprint(userID), c)
		if !valid {
			default_functions.RespondWithError(c, http.StatusInternalServerError, "Error of permission check")
			c.Abort()
		}

		var user *models.User
		user, err := (*userRepo).GetByID(userIDConvert)
		if err != nil || user == nil {
			default_functions.RespondWithError(c, http.StatusBadRequest, "Error of permission check")
			c.Abort()
		}

		isAdmin := false
		for _, permission := range user.Role.Permissions {
			if permission.Name == permissions {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}
