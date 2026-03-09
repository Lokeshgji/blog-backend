package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

func ExtractToken(c *gin.Context) (string, error) {

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header missing")
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", errors.New("invalid token format")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	if tokenString == "" {
		return "", errors.New("token missing")
	}

	return tokenString, nil
}
