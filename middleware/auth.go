package middleware

import (
	"net/http"

	"blog-platform/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("secret")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString, err := utils.ExtractToken(c)

		if err != nil {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})

			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})

			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		userID := claims["user_id"]

		c.Set("user_id", userID)

		c.Next()
	}
}
