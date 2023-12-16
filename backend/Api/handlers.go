package api

import (
	todo "api/backend/Todo"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context, db *sql.DB) {
	var tasksOutput []todo.Todo

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var task todo.Todo
		if err := rows.Scan(&task.ID, &task.Completed, &task.Item); err != nil {
			log.Fatal(err)
			c.AbortWithStatus(http.StatusBadRequest)
		}
		tasksOutput = append(tasksOutput, task)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.IndentedJSON(http.StatusOK, tasksOutput)
}

func FetchTodoById(id string, db *sql.DB) (*todo.Todo, error) {
	for i, task := range todo.Todos {
		if task.ID == id {
			return &todo.Todos[i], nil
		}
	}
	return nil, fmt.Errorf("could not find a todo item with id %s", id)
}

func GetTodoById(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	todo, err := FetchTodoById(id, db)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func ToggleStatus(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	todo, err := FetchTodoById(id, db)
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

func AddTodo(c *gin.Context, db *sql.DB) {
	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	todo.Todos = append(todo.Todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context, db *sql.DB) {
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
	todo.Todos = append(todo.Todos[:index], todo.Todos[index+1:]...)
	c.IndentedJSON(http.StatusOK, todo.Todos)
}
