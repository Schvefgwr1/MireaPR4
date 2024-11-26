package default_functions

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RespondWithError Helper для обработки ошибок
func RespondWithError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(statusCode, gin.H{
		"status": "error",
		"error":  message,
	})
}

// RespondWithSuccess Helper для успешных ответов
func RespondWithSuccess(ctx *gin.Context, statusCode int, data any) {
	ctx.JSON(statusCode, gin.H{
		"status": "success",
		"data":   data,
	})
}

// ValidateJSON Helper для проверки Content-Type
func ValidateJSON(ctx *gin.Context) bool {
	if ctx.GetHeader("Content-Type") != "application/json" {
		RespondWithError(ctx, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return false
	}
	return true
}

// ConvertStrToIntParam Helper для преобразования id из строки в int
func ConvertStrToIntParam(param string, ctx *gin.Context) (int, bool) {
	idInt, err := strconv.Atoi(param)
	if err != nil {
		RespondWithError(ctx, http.StatusBadRequest, "Invalid ID format")
		return 0, false
	}
	return idInt, true
}
