package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	// Use the config
	router.Use(cors.New(config))
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodoById)
	router.PATCH("/todos/:id", ToggleStatus)
	router.POST("/todos/", AddTodo)
	router.Run("localhost:1234")
}
