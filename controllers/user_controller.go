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

func DeleteUser(c *gin.Context) {

	userID, _ := c.Get("user_id")

	// delete user's articles first
	_, err := config.DB.Exec(
		"DELETE FROM articles WHERE author_id=$1",
		userID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Could not delete user articles",
		})

		return
	}

	// delete user
	_, err = config.DB.Exec(
		"DELETE FROM users WHERE id=$1",
		userID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Could not delete user",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "User deleted successfully",
	})
}
