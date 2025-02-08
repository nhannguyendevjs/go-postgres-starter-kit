package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"hoainhaannguyen/go-postgres-starter-kit/pkg/logger"
	"hoainhaannguyen/go-postgres-starter-kit/pkg/middleware"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		logger.Info("No .env file found, using default values", nil)
	}
}

func main() {
	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()

	// Apply CORS middleware 3rd party library
	// AllowOrigins: Specifies allowed origins (use "*" to allow all origins, but it’s not recommended for production).
	// AllowMethods: Defines HTTP methods that can be used in requests.
	// AllowHeaders: Specifies allowed request headers.
	// ExposeHeaders: Allows the client to access specific response headers.
	// AllowCredentials: Enables cookies or authentication credentials to be sent.
	// MaxAge: Specifies how long the preflight response can be cached.
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://example.com", "http://localhost:3000"}, // Change as needed
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour, // Preflight request cache duration
	// }))

	// Apply CORS middleware custom
	r.Use(middleware.CORSMiddleware())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	address := fmt.Sprintf("%s:%s", host, port)

	r.Run(address)
	logger.Info("Starting server", map[string]interface{}{"address": address})
}
