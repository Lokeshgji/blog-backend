package controllers

import (
	"blog-platform/config"
	"blog-platform/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, err)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	query := `
INSERT INTO users (name,email,password,role)
VALUES ($1,$2,$3,'author')
`

	config.DB.Exec(query,
		user.Name,
		user.Email,
		string(hash))

	c.JSON(200, gin.H{"message": "User created"})
}

func Login(c *gin.Context) {

	var input models.User
	var user models.User

	c.BindJSON(&input)

	err := config.DB.Get(&user,
		"SELECT * FROM users WHERE email=$1",
		input.Email)

	if err != nil {
		c.JSON(401, "Invalid credentials")
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(401, "Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.Email,
		})

	tokenString, _ := token.SignedString([]byte("secret"))

	c.JSON(200, gin.H{
		"token": tokenString,
	})

}
