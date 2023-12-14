package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func StartServer() {
	router := gin.Default()
	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowMethods("OPTIONS")
	config.MaxAge = time.Hour
	// Use the config
	router.Use(cors.New(config))
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodoById)
	router.PATCH("/todos/:id", ToggleStatus)
	router.POST("/todos/", AddTodo)

	// Preflight request problem.
	router.OPTIONS("/todos", func(c *gin.Context) {
		c.AbortWithStatus(200)
	})
	router.Run("localhost:1234")
}
