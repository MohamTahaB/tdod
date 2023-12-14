package main

import (
	todo "api/backend/Todo"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo.Todos)
}

func fetchTodoById(id string) (*todo.Todo, error) {
	for i, task := range todo.Todos {
		if task.ID == id {
			return &todo.Todos[i], nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Could not find a todo item with id %s", id))
}

func getTodoById(c *gin.Context) {
	id := c.Param("id")
	todo, err := fetchTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func toggleStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := fetchTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	toggleVal := true
	if todo.Completed {
		toggleVal = false
	}
	todo.Completed = toggleVal
	c.IndentedJSON(http.StatusOK, todo)
}

func addTodo(c *gin.Context) {
	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todo.Todos = append(todo.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

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
