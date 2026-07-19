# Design Pattern and Architecture Made EZ

The design of this REST API architecture are the simplified form of standard Go REST API nowadays. Don't worry if you hardly understanding go, I cover every single detail and fundamental concept needed as detail as possible. <br />
Here's the brief overview of this project directory

```sh
|--- .air.toml              # for air live reload
|--- .env.example           # .env placeholder
|--- .gitignore             # git ignored file
|--- .prettierignore        # prettier ignored file
|--- .prettierrc            # prettier extension configuration
|--- .vscode                # vscode configuration
|   |--- settings.json      # custom settings.json
|--- cmd                    # main app program
|   |--- api                # api related program
|--- design-pattern.md      # this file
|--- go.mod                 # go module manager
|--- go.sum                 # go package library manager
|--- internal               # main program app architecture development
|   |--- config             # main program app configuration
|   |--- controllers        # api function driver
|   |--- helpers            # function helper and aliasing for reusable code
|   |--- middlewares        # connector guard and validation between controller and repo
|   |--- models             # app data structure
|--- LICENSE                # project license
|--- migrations             # database manual migration
|--- project-tree.ps1       # software to create this tree structure
|--- README.md              # project documentation
|--- scripts                # other developer's script helper
|--- tmp                    # air and go build temporary file
```

Notice that there's no `repository` and only `controller` because controller is already interactimg with database directly making it controller and repository at the same time

## Data Structure

The first thing you want to made is data structure over your project entity and database schema. <br />
In this repo, the database schema and relation is something like this:

```sh
Events -> User # 1 to 1, every event need exactly 1 userId
User -> Events # 1 to many, every user can join other event independently
```

## Entrance Point

The app entrance point is in `cm/api/main.go`. Click [here](./cmd/api/main.go) to examine.

1. Load `.env` data. By doing that, we can access `.env` from `os.Getenv()`

```go
err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
```

2. Setting up database configuration

```go
var DB *gorm.DB

func ConnectDB() {
	log.Println("🚀 Starting database connection")
	dsn := os.Getenv("DATABASE_URI")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URI not found in .env")
	}

    // load database query to gorm, we use postgre so we open the postgre with default gorm config
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Failed to connect the database : ", err)
	}

    // automatically create schema and load schema by using gorm.DB.AutoMigrate
	err = database.AutoMigrate(&models.Event{}, models.User{}, models.Image{})
	if err != nil {
		log.Fatal("❌ Failed to do database migration : ", err)
	}

    // set global variable of DB to database so that we can access DB globally scoped
	DB = database
	log.Println("🎉 Success to connect database")
}
```

Then we run it in our entry point app

```go
config.ConnectDB()
```

3. Setting up router and server host and port

```go
// setup router endpoint with gin
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

	defer server.Run(serverUrl) // set endpoint root, you want the server to run at last place using defer
```

4. The REST API endpoint
   The REST API is setup in two routes, public routes and protected routes. Protected routes have additional middlaware feature to validate the incoming data from controller before interacting with database

```go
// you want to have versioned route so that massive change didn't corrupt, i.e. ehance continous integration
V1_api := server.Group("/api/v1")
	{
		// events

		// /events?search=&page=&limit=
		V1_api.GET("/events", controllers.GetEvents)
		V1_api.GET("/events/:id", controllers.GetEventById) // /events/the_event_id

		// user
		V1_api.POST("/auth/register", controllers.RegisterUser) // register the user
		V1_api.POST("/auth/login", controllers.UserLogin) // handling user login

		protected := V1_api.Group("/") // group the /api/v1/ route with middleware
		protected.Use(middlewares.RequiredAuth())
		{
            // this will be all the protected REST API that only can be acccessed by jwt
			protected.GET("/auth/me", controllers.GetCurrentUser)
			protected.POST("/events", controllers.CreateEvent)
			protected.PUT("/events/:id", controllers.UpdateEvent)
			protected.DELETE("/events/:id", controllers.DeleteEvent)
			protected.POST("/images", controllers.UploadImage)
		}
	}
```

## Entity Modeling and Data Transfer Object (DTO)

In this repo template, I'm not differentiate entity model and data transfer object (DTO) into separate file, both unified in single file for simplicity. All entity and DTO are stored in `./internal/models/` directory. Below are the example of `event` entity and DTO:

```go
// ./internal/models/event.go

package models

import (
	"time"

	"github.com/google/uuid"
)

// entity
type Event struct {
	BaseModel
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      uuid.UUID `json:"user_id" gorm:"not null"`
	User        User      `gorm:"foreignKey:UserId" json:"user"` // Event to User is many to one
	ImageID     uuid.UUID `json:"image_id" gorm:"not null"`
	Image       Image     `json:"image" gorm:"foreignkey:ImageID"` // Event to Image is one to one
	DateTime    time.Time `json:"datetime" binding:"required"`
}

// dto
type CreateEventRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	ImageID     uuid.UUID `json:"image_id" binding:"required"` // client will receive new uploaded image url
}
```

## Middleware Details

In this repo template, the middleware is currently only one that is user authentication for protected endpoints. I will add more security feature middleware in the future. Below is the implementation:

```go
// ./internal/middlewares/auth_middleware.go

func RequiredAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
        // get jwt from header value of key 'Authorization'
		tokenStr := ctx.GetHeader("Authorization")

        // check that jwt is Bearer and not nil/null
		if tokenStr == "" || !strings.HasPrefix(tokenStr, "Bearer ") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":     false,
				"status_code": http.StatusUnauthorized,
				"message":     "Access denied, token is not found or invalid format",
				"error":       "Unauthorized action detected, Authorization token with 'Bearer' prefix is not found",
			})
			return
		}

        //get the token string by trimming 'Bearer' prefix
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

        // parse the jwt to make sure jwt is valid format
		token, errTokenStr := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if errTokenStr != nil {
			fmt.Println(errTokenStr)
			fmt.Println(os.Getenv("JWT_SECRET"))
		}

        // if the jwt is correct and valid, set the gin context to user id get from jwt in sub key
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("user_id", claims["sub"].(string))
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"success":     false,
				"status_code": http.StatusUnauthorized,
				"message":     "Access denied, token is valid by format but failed in JWT validation, the token is probably expired, try to relogin",
				"error":       "Unauthorized action detected, JWT validation failed",
			})
		}
	}
}
```

## Controller Logic

Controller logic are completely freedom as how you want to design your REST API structure. For this repo, I only implement basic CRUD controller for each endpoints, especially `event`. To create controller, we need a handler/helper to simplify the logic.

### Helper or pkg

In this repo, I provide the response handler or helper in `response.go` for reusable response function. Then there's also `get_user_id_from_ctx.go` to get parsed UUID of user id accessing the endpoint from gin context that is from user authorization jwt. All other helper must be stored in `./internal/helper`. I'm not name it to pkg as it can cause interference naming in GO Workspace package for some devices.

```go
// ./internal/helpers/reponse.go
package helpers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// succes response [2xx] dto
type SuccessResponseHeader struct {
	Success    bool        `json:"succes"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	JwtToken   string      `json:"jwt_token,omitempty"`
	Meta       interface{} `json:"meta"`
}

// error response [4xx, 5xx] dto
type ErrorResponseHeader struct {
	Success    bool   `json:"succes"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

// reusable function for success [2xx]
func RespSuccess(ctx *gin.Context, statusCode int, message string, data interface{}, jwtToken string, meta interface{}) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(statusCode, SuccessResponseHeader{
		Success:    true,
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		JwtToken:   jwtToken,
		Meta:       meta,
	})
}

// reusable function for error [4xx, 5xx]
func RespError(ctx *gin.Context, statusCode int, message string, err error, errStr ...string) {
	finalMessage := message
	if finalMessage == "" {
		finalMessage = http.StatusText(statusCode)
	}

	var finalErrString string

	if len(errStr) > 0 && errStr[0] != "" {
		finalErrString = errStr[0]
	} else if err != nil {
		finalErrString = err.Error()
	}

	ctx.JSON(statusCode, ErrorResponseHeader{
		Success:    false,
		StatusCode: statusCode,
		Message:    finalMessage,
		Error:      finalErrString,
	})
}

```

```go
// ./internal/helpers/get_user_id_from_ctx.go

package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetUserIDFromCtx mengambil user_id dari Gin Context dan mem-parsing-nya menjadi uuid.UUID.
func GetUserIDFromCtx(ctx *gin.Context) (uuid.UUID, error) {
    // check 'user_id' key value in gin context
	userIdInterface, exists := ctx.Get("user_id")
	if !exists {
		return uuid.Nil, fmt.Errorf("user_id not found in context")
	}

    // parse to string
	userIdStr, ok := userIdInterface.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user_id in context is not a string")
	}

    // parse to uuid
	parsedUUID, err := uuid.Parse(userIdStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID format: %w", err)
	}

    // return the parsed user_id uuid and nil error
	return parsedUUID, nil
}

```

### Main Logic

The main logic is like I said completely freedom as how you want to design it. I can't explain it one by one as it was a lot of word to describe. I will try to add it later if I have time. For now, try to read and understand the code I write in `./internal/controllers/` directory.
