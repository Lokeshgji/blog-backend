package middleware

import (
	"net/http"

	"blog-platform/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("9f94493f-945b-4326-9801-829696eda26e")

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenString, err := utils.ExtractToken(c)

		if err != nil {

			if err.Error() == "Token is expired" {

				c.JSON(401, gin.H{
					"error": "Token expired. Please login again",
				})

			} else {

				c.JSON(401, gin.H{
					"error": "Invalid token",
				})

			}

			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if !token.Valid {

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
