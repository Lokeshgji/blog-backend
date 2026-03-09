package controllers

import (
	"blog-platform/config"
	"blog-platform/models"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {

	var article models.Article

	err := c.BindJSON(&article)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	article.Slug = generateSlug(article.Title)

	// Validation
	if article.Title == "" || article.Content == "" || article.Slug == "" {

		c.JSON(400, gin.H{
			"error": "Title, slug and content are required",
		})

		return
	}

	userID, _ := c.Get("user_id")

	query := `
INSERT INTO articles(title,slug,content,author_id)
VALUES($1,$2,$3,$4)
`

	_, err = config.DB.Exec(
		query,
		article.Title,
		article.Slug,
		article.Content,
		userID,
	)

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Could not create article",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Article created successfully",
	})
}

func GetArticles(c *gin.Context) {

	var articles []models.ArticleWithAuthor

	query := `
	SELECT 
	a.id,
	a.title,
	a.slug,
	a.content,
	u.name AS author_name
	FROM articles a
	LEFT JOIN users u ON a.author_id = u.id
	ORDER BY a.id DESC
	`

	err := config.DB.Select(&articles, query)

	if err != nil {

		fmt.Println("ERROR FETCHING ARTICLES:", err)

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, articles)
}

func GetArticleBySlug(c *gin.Context) {

	slug := c.Param("slug")

	var article models.Article

	query := `
	SELECT id,title,slug,content,author_id
	FROM articles
	WHERE slug=$1
	`

	err := config.DB.Get(&article, query, slug)

	if err != nil {

		fmt.Println("ARTICLE FETCH ERROR:", err)

		c.JSON(404, gin.H{
			"error": "Article not found",
		})

		return
	}

	c.JSON(200, article)
}

func GetArticlesByAuthor(c *gin.Context) {

	authorID := c.Param("id")

	var articles []models.Article

	query := `
SELECT id,title,slug,content,author_id,status
FROM articles
WHERE author_id=$1
`

	err := config.DB.Select(&articles, query, authorID)

	if err != nil {

		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, articles)

}

func DeleteArticle(c *gin.Context) {

	id := c.Param("id")

	userID, _ := c.Get("user_id")

	query := `
	DELETE FROM articles
	WHERE id=$1 AND author_id=$2
	`

	result, err := config.DB.Exec(query, id, userID)

	if err != nil {

		c.JSON(500, gin.H{
			"error": "Could not delete article",
		})

		return
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {

		c.JSON(403, gin.H{
			"error": "You are not allowed to delete this article",
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "Article deleted successfully",
	})
}

func generateSlug(title string) string {

	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.Trim(slug, "-")

	return slug
}
