package controllers

import (
	"blog-platform/config"
	"blog-platform/models"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	// Field specific validation

	if strings.TrimSpace(user.Name) == "" {
		c.JSON(400, gin.H{
			"error": "Name is required",
		})
		return
	}

	if strings.TrimSpace(user.Email) == "" {
		c.JSON(400, gin.H{
			"error": "Email is required",
		})
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

	if !emailRegex.MatchString(user.Email) {

		c.JSON(400, gin.H{
			"error": "Invalid email format",
		})

		return
	}

	var existingUser models.User

	err = config.DB.Get(&existingUser,
		"SELECT id FROM users WHERE email=$1",
		user.Email)

	if err == nil {

		c.JSON(400, gin.H{
			"error": "Email already registered",
		})

		return
	}

	if strings.TrimSpace(user.Password) == "" {
		c.JSON(400, gin.H{
			"error": "Password is required",
		})
		return
	}

	// Password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error hashing password",
		})
		return
	}

	query := `
	INSERT INTO users(name,email,password)
	VALUES($1,$2,$3)
	`

	_, err = config.DB.Exec(
		query,
		user.Name,
		user.Email,
		string(hashedPassword),
	)

	if err != nil {
		c.JSON(500, gin.H{
			"error": "User could not be created",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Signup successful",
	})
}

func Login(c *gin.Context) {

	var loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.BindJSON(&loginRequest)

	var user models.User

	err := config.DB.Get(
		&user,
		"SELECT id,email,password FROM users WHERE email=$1",
		loginRequest.Email,
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(loginRequest.Password),
	)

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Could not generate token",
		})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenString,
	})
}
