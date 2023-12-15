package api

import (
	todo "api/backend/Todo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todo.Todos)
}

func FetchTodoById(id string) (*todo.Todo, error) {
	for i, task := range todo.Todos {
		if task.ID == id {
			return &todo.Todos[i], nil
		}
	}
	return nil, fmt.Errorf("could not find a todo item with id %s", id)
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	todo, err := FetchTodoById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func ToggleStatus(c *gin.Context) {
	id := c.Param("id")
	todo, err := FetchTodoById(id)
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

func AddTodo(c *gin.Context) {
	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todo.Todos = append(todo.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	index := -1
	for i, task := range todo.Todos {
		if task.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.IndentedJSON(http.StatusOK, todo.Todos)
}
