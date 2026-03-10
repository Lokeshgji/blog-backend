package routes

import (
	"blog-platform/controllers"
	"blog-platform/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/articles", controllers.GetArticles)
	r.GET("/articles/:slug", controllers.GetArticleBySlug)
	r.GET("/authors/:id", controllers.GetAuthor)
	r.GET("/authors/:id/articles", controllers.GetArticlesByAuthor)
	r.DELETE("/articles/:id", middleware.AuthMiddleware(), controllers.DeleteArticle)
	r.DELETE("/users/me", middleware.AuthMiddleware(), controllers.DeleteUser)
	r.PUT("/articles/:id", middleware.AuthMiddleware(), controllers.UpdateArticle)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/articles", controllers.CreateArticle)
}
