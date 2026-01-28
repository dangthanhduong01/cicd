package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
	User    string `json:"user,omitempty"`
}

func main() {
	// Tạo router Gin
	r := gin.Default()

	// Middleware CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Route health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API is running!",
			"version": "1.0.0",
		})
	})

	// Route health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Route login
	r.POST("/api/login", handleLogin)

	// Route để test authentication
	r.GET("/api/profile", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No authorization token provided",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Profile data",
			"user":    "john_doe",
			"email":   "john@example.com",
		})
	})

	// Lấy port từ environment variable hoặc dùng 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

func handleLogin(c *gin.Context) {
	var loginReq LoginRequest

	// Bind JSON request
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, LoginResponse{
			Message: "Invalid request format",
		})
		return
	}

	// Validation đơn giản (trong thực tế nên dùng database)
	if loginReq.Username == "admin" && loginReq.Password == "password123" {
		// Login thành công
		c.JSON(http.StatusOK, LoginResponse{
			Message: "Login successful",
			Token:   "fake-jwt-token-" + loginReq.Username,
			User:    loginReq.Username,
		})
		return
	}

	// Login thất bại
	c.JSON(http.StatusUnauthorized, LoginResponse{
		Message: "Invalid username or password",
	})
}
