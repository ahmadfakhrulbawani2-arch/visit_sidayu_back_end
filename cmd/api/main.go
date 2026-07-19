package main

import (
	"fmt"
	"log"
	"os"
	"simple_go_gin_gorm_postgres_be_template/internal/config"
	"simple_go_gin_gorm_postgres_be_template/internal/controllers"
	"simple_go_gin_gorm_postgres_be_template/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()

	server := gin.Default()

	// pick host and port from env
	var host, port string
	switch mode := os.Getenv("RUN_MODE"); mode {
	case "production":
		host = os.Getenv("PROD_HOST")
		port = os.Getenv("PROD_PORT")

	case "development":
		host = os.Getenv("DEV_HOST")
		port = os.Getenv("DEV_PORT")

	default:
		// Fallback if RUN_MODE is empty, "development", or anything else
		host = os.Getenv("DEV_HOST")
		port = os.Getenv("DEV_PORT")
	}

	// Fallback to localhost if host is still empty to prevent crashes
	if host == "" {
		host = "127.0.0.1"
	}

	server.SetTrustedProxies([]string{host})
	serverUrl := fmt.Sprintf("%s:%s", host, port)

	defer server.Run(serverUrl) // set endpoint root

	V1_api := server.Group("/api/v1")
	{
		// events

		// /events?search=&page=&limit=
		V1_api.GET("/events", controllers.GetEvents)
		V1_api.GET("/events/:id", controllers.GetEventById)

		// user
		V1_api.POST("/auth/register", controllers.RegisterUser)
		V1_api.POST("/auth/login", controllers.UserLogin)

		protected := V1_api.Group("/")
		protected.Use(middlewares.RequiredAuth())
		{
			protected.GET("/auth/me", controllers.GetCurrentUser)
			protected.POST("/events", controllers.CreateEvent)
			protected.PUT("/events/:id", controllers.UpdateEvent)
			protected.DELETE("/events/:id", controllers.DeleteEvent)
			protected.POST("/images", controllers.UploadImage)
		}
	}
}
