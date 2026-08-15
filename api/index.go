/*
 * For vercel deployment, idk i try using render and i cant
 * MIT License Copyright (c) 2026 Ahmad Fakhrul Bawani
 */
package api

import (
	"net/http"
	"sync"
	"time"

	"visit-sidayu-backend/pkg/config"
	ctr "visit-sidayu-backend/pkg/controllers"
	mdw "visit-sidayu-backend/pkg/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router *gin.Engine
	once   sync.Once
)

func initRouter() {
	_ = godotenv.Load()

	config.ConnectDB()

	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(mdw.CorsConfig()))
	router.Use(gin.Recovery())
	router.Use(mdw.SecurityHeader())
	router.Use(mdw.RequestID())
	router.Use(mdw.BodyLimit(5 << 20))
	router.Use(mdw.Timeout(30 * time.Second))

	V1_api := router.Group("/api/v1")

	V1_api.Use(mdw.RateLimiter(5000.0/60.0, 5000)) // limit to 1000 requests per minute and burst of 1000
	{
		// expose
		V1_api.GET("/expose", ctr.ExposeHandler)

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
			// sa_api.POST("/register", mdw.RateLimiter(10.0/60.0, 10), ctr.SuperAdminRegister) // ✅
			sa_api.POST("/login", mdw.RateLimiter(20.0/60.0, 20), ctr.SuperadminLogin)       // ✅
		}

		img_api := V1_api.Group("/images")
		img_api.Use(mdw.RateLimiter(2500.0/60.0, 1250)) // because its size access, better to limit the rate to 2500 requests per minute and burst of 1250
		{
			img_api.GET("", ctr.GetImages)
			img_api.GET("/id/:id", ctr.GetImageById)
		}

		blog_api := V1_api.Group("/blogs")
		{
			blog_api.GET("", ctr.GetAllBlogs)             // ✅
			blog_api.GET("/id/:id", ctr.GetBlogById)       // ✅
			blog_api.GET("/slug/:slug", ctr.GetBlogBySlug) // ✅
		}

		cb_api := V1_api.Group("/culture-blogs")
		{
			cb_api.GET("", ctr.GetAllCultureBlog)
			cb_api.GET("/id/:id", ctr.GetCultureBlogById)
			cb_api.GET("slug/:slug", ctr.GetCultureBlogBySlug)
		}

		d_api := V1_api.Group("/demographies")
		{
			d_api.GET("", ctr.GetAllDemographies)
			d_api.GET("/id/:id", ctr.GetDemographyById)
			d_api.GET("/district", ctr.GetDistrictDemography)
		}

		g_api := V1_api.Group("/galleries")
		{
			g_api.GET("", ctr.GetAllGalleries)
			g_api.GET("/id/:id", ctr.GetGalleryByID)
			g_api.GET("/slug/:slug", ctr.GetGalleryBySlug)
		}

		geo_api := V1_api.Group("/geographies")
		{
			geo_api.GET("", ctr.GetAllGeographies)
			geo_api.GET("/id/:id", ctr.GetGeographyByID)
			geo_api.GET("/district", ctr.GetDistrictGeographies)
		}

		ind_api := V1_api.Group("/industries-blogs")
		{
			ind_api.GET("", ctr.GetAllIndustriesBlog)
			ind_api.GET("/id/:id", ctr.GetIndustryBlogByID)
			ind_api.GET("/slug/:slug", ctr.GetIndustryBlogBySlug)
		}

		ofc_api := V1_api.Group("/officials")
		{
			ofc_api.GET("", ctr.GetAllOfficials)
			ofc_api.GET("/id/:id", ctr.GetOfficialByID)
		}

		si_api := V1_api.Group("/shops-umkms")
		{
			si_api.GET("", ctr.GetAllShopsAndUmkms)
			si_api.GET("/id/:id", ctr.GetShopAndUmkmByID)
			si_api.GET("/slug/:slug", ctr.GetShopAndUmkmBySlug)
		}

		tm_api := V1_api.Group("/timelines")
		{
			tm_api.GET("", ctr.GetAllTimelines)
			tm_api.GET("/id/:id", ctr.GetTimelineByID)
		}
		rl_api := V1_api.Group("/roles")
		{
			rl_api.GET("", ctr.GetAllRoles)
			rl_api.GET("/id/:id", ctr.GetRoleById)
		}

		protected := V1_api.Group("")
		protected.Use(mdw.RateLimiter(500.0/60.0, 500),mdw.RequiredAuth()) // ✅
		{
			sa_api := protected.Group("/superadmins")
			{
				auth_api := sa_api.Group("/auth")
				{
					auth_api.POST("/register", mdw.RateLimiter(10.0/60.0, 10), ctr.SuperAdminRegister)
					auth_api.POST("/jwt", mdw.RateLimiter(500.0/60.0, 500), ctr.PostGetJwtValidated)
				}
				sa_api.GET("/me", ctr.SuperadminCurrent)       // ✅
				sa_api.GET("", ctr.GetAllSuperadmins)         // ✅
				sa_api.GET("/id/:id", ctr.GetSuperadminByID)   // ✅
				sa_api.PATCH("/id/:id", ctr.UpdateSuperadmin)  // ✅
				sa_api.DELETE("/id/:id", ctr.DeleteSuperadmin) // ✅
			}
			img_api := protected.Group("/images")
			img_api.Use(mdw.RateLimiter(250.0/60.0, 125))
			{
				img_api.POST("", ctr.UploadImage)
				img_api.PUT("/id/:id", ctr.UpdateImage)
				img_api.DELETE("/id/:id", ctr.DeleteImage)
			}

			blog_api := protected.Group("/blogs")
			{
				blog_api.POST("", ctr.CreateBlog)         // ✅
				blog_api.PATCH("/id/:id", ctr.UpdateBlog)  // ✅
				blog_api.DELETE("/id/:id", ctr.DeleteBlog) // ✅
			}

			cb_api := protected.Group("/culture-blogs")
			{
				cb_api.POST("", ctr.CreateCultureBlog)
				cb_api.PATCH("/id/:id", ctr.UpdateCultureBlog)
				cb_api.DELETE("/id/:id", ctr.DeleteCultureBlog)
			}

			d_api := protected.Group("/demographies")
			{
				d_api.POST("", ctr.CreateDemography)
				d_api.PATCH("/id/:id", ctr.UpdateDemography)
				d_api.DELETE("/id/:id", ctr.DeleteDemography)
			}

			g_api := protected.Group("/galleries")
			{
				g_api.POST("", ctr.CreateGallery)
				g_api.PATCH("/id/:id", ctr.UpdateGallery)
				g_api.DELETE("/id/:id", ctr.DeleteGallery)
			}

			geo_api := protected.Group("/geographies")
			{
				geo_api.POST("", ctr.CreateGeography)
				geo_api.PATCH("/id/:id", ctr.UpdateGeography)
				geo_api.DELETE("/id/:id", ctr.DeleteGeograhy)
			}

			ind_api := protected.Group("/industries-blogs")
			{
				ind_api.POST("", ctr.CreateIndustryBlog)
				ind_api.PATCH("/id/:id", ctr.UpdateIndustryBlog)
				ind_api.DELETE("/id/:id", ctr.DeleteIndustryBlog)
			}

			ofc_api := protected.Group("/officials")
			{
				ofc_api.POST("", ctr.CreateOfficial)
				ofc_api.PATCH("/id/:id", ctr.UpdateOfficial)
				ofc_api.DELETE("/id/:id", ctr.DeleteOfficial)
			}

			si_api := protected.Group("/shops-umkms")
			{
				si_api.POST("", ctr.CreateShopAndUmkm)
				si_api.PATCH("/id/:id", ctr.UpdateShopAndUmkm)
				si_api.DELETE("/id/:id", ctr.DeleteShopAndUmkm)
			}

			tm_api := protected.Group("/timelines")
			{
				// tm_api.POST("", ctr.CreateTimeline)
				tm_api.PATCH("/id/:id", ctr.UpdateTimeline)
				tm_api.DELETE("/id/:id", ctr.DeleteTimeline)
			}

			rl_api := protected.Group("/roles")
			{
				rl_api.POST("", ctr.CreateRole)
				rl_api.PATCH("/id/:id", ctr.UpdateRole)
				rl_api.DELETE("/id/:id", ctr.DeleteRole)
			}
		}
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	once.Do(initRouter)

	router.ServeHTTP(w, r)
}
