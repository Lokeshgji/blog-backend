package controllers

import (
	"blog-platform/config"
	"blog-platform/models"

	"github.com/gin-gonic/gin"
)

func GetAuthor(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	err := config.DB.Get(&user,
		"SELECT id,name,email FROM users WHERE id=$1",
		id)

	if err != nil {

		c.JSON(404, gin.H{
			"error": "Author not found",
		})

		return
	}

	response := gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	}

	// Include created_at only if it exists
	if user.CreatedAt != "" {
		response["created_at"] = user.CreatedAt
	}

	c.JSON(200, response)
}
