package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	// Define a CORS Config.
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	// Use the middleware.
	router.Use(cors.New(config))
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodoById)
	router.PATCH("/todos/:id", ToggleStatus)
	router.POST("/todos", AddTodo)
	router.Run("localhost:1234")
}
