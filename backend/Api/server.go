package api

import (
	"database/sql"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB) {
	router := gin.Default()
	// Define a CORS Config.
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "accept", "origin", "X-Requested-With"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	// Use the middleware.
	router.Use(cors.New(config))
	router.GET("/todos", func(c *gin.Context) { GetTodos(c, db) })
	router.GET("/todos/:id", func(c *gin.Context) { GetTodoById(c, db) })
	router.PATCH("/todos/:id", func(c *gin.Context) { ToggleStatus(c, db) })
	router.POST("/todos", func(c *gin.Context) { AddTodo(c, db) })
	router.DELETE("/todos/:id", func(c *gin.Context) { DeleteTodo(c, db) })
	router.OPTIONS("todos", Options)
	router.OPTIONS("todo/:id", Options)
	router.Run("localhost:1234")
}
