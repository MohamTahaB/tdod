package backend

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// Use the config
	router.Use(cors.New(config))
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodoById)
	router.PATCH("/todos/:id", toggleStatus)
	router.POST("/todos/", addTodo)
	router.Run("localhost:1234")
}
