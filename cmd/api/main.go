package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"visit-sidayu-backend/internal/config"
	ctr "visit-sidayu-backend/internal/controllers"
	mdw "visit-sidayu-backend/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
	defer config.DisconnectDB()

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

	V1_api := server.Group("/api/v1")
	{
		// health
		V1_api.GET("/ping", func(ctx *gin.Context) {
			ctx.Header("Content-Type", "application/json")

			ctx.JSON(http.StatusOK, gin.H{
				"status":  200,
				"message": "pong",
			})
		})

		// superadmin
		sa_api := V1_api.Group("/superadmins/auth")
		{
			// auth
			sa_api.POST("/register", ctr.SuperAdminRegister) // ✅
			sa_api.POST("/login", ctr.SuperadminLogin)       // ✅
		}

		img_api := V1_api.Group("/images")
		{
			img_api.GET("/", ctr.GetImages)
			img_api.GET("/id/:id", ctr.GetImageById)
		}

		blog_api := V1_api.Group("/blogs")
		{
			blog_api.GET("/", ctr.GetAllBlogs)
			blog_api.GET("/id/:id", ctr.GetBlogById)
			blog_api.GET("/slug/:slug", ctr.GetBlogBySlug)
		}

		protected := V1_api.Group("/")
		protected.Use(mdw.RequiredAuth()) // ✅
		{
			sa_api := protected.Group("/superadmins")
			{
				sa_api.GET("/me", ctr.SuperadminCurrent)       // ✅
				sa_api.GET("/", ctr.GetAllSuperadmins)         // ✅
				sa_api.GET("/id/:id", ctr.GetSuperadminByID)   // ✅
				sa_api.PATCH("/id/:id", ctr.UpdateSuperadmin)  // ✅
				sa_api.DELETE("/id/:id", ctr.DeleteSuperadmin) // ✅
			}
			img_api := protected.Group("/images")
			{
				img_api.POST("/", ctr.UploadImage)
				img_api.PUT("/id/:id", ctr.UpdateImage)
				img_api.DELETE("/id/:id", ctr.DeleteImage)
			}

			blog_api := protected.Group("/blogs")
			{
				blog_api.POST("/", ctr.CreateBlog)
				blog_api.PATCH("/id/:id", ctr.UpdateBlog)
				blog_api.DELETE("/id/:id", ctr.DeleteBlog)
			}
		}
	}

	server.Run(serverUrl) // set endpoint root
}
