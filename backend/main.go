package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{
		ID:        "1",
		Item:      "Clean room",
		Completed: false,
	},
	{
		ID:        "2",
		Item:      "Read book",
		Completed: false,
	},
	{
		ID:        "3",
		Item:      "Record video",
		Completed: true,
	},
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func fetchTodoById(id string) (*todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			return &todos[i], nil
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
	var newTodo todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
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
