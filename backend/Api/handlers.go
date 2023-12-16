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
			return
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
	var task todo.Todo

	row := db.QueryRow("SELECT * FROM tasks WHERE id = ?", id)
	if err := row.Scan(&task.ID, &task.Completed, &task.Item); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("could not find a todo item with id %s", id)
	}
	return &task, nil
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
	task, err := FetchTodoById(id, db)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	toggleValue := 0
	if !task.Completed {
		toggleValue = 1
	}
	row := db.QueryRow(fmt.Sprintf("UPDATE tasks SET completed = %d WHERE id = ?", toggleValue), id)
	if err := row.Scan(&task.ID, &task.Completed, &task.Item); err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	task, err = FetchTodoById(id, db)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo item not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, task)
}

func AddTodo(c *gin.Context, db *sql.DB) {
	var newTodo todo.Todo
	if err := c.BindJSON(&newTodo); err != nil {
		return
	}
	completedVal := 0
	if newTodo.Completed {
		completedVal = 1
	}

	result, err := db.Exec("INSERT INTO tasks (id, completed, item) VALUES (?, ?, ?)", newTodo.ID, completedVal, newTodo.Item)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func DeleteTodo(c *gin.Context, db *sql.DB) {
	id := c.Param("id")
	result, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if _, err = result.RowsAffected(); err != nil {
		log.Fatal(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Retrieve the new list of elements.
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
			return
		}
		tasksOutput = append(tasksOutput, task)
	}
	c.IndentedJSON(http.StatusOK, tasksOutput)
}
