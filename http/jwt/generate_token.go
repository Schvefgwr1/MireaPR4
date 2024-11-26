package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

// InitJWTSecret устанавливает секретный ключ из конфигурации
func InitJWTSecret(secret string) {
	jwtSecret = []byte(secret)
}

type TokenClaims struct {
	UserID int   `json:"user_id"`
	Exp    int64 `json:"exp"`
}

func GenerateToken(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Проверка валидности токена и извлечение claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, okUserID := claims["user_id"].(float64)
		exp, okExp := claims["exp"].(float64)

		if !okUserID || !okExp {
			return nil, errors.New("invalid claims structure")
		}

		return &TokenClaims{
			UserID: int(userID),
			Exp:    int64(exp),
		}, nil
	}
	return nil, errors.New("invalid token")
}
