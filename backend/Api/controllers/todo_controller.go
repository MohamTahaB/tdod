package controllers

import (
	"api/backend/Api/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateToDo is a function that creates a new todo item
// It takes in a gin.Context and returns a JSON response
// It requires a valid token to be passed in the request header
// It requires a valid JSON body with the todo item details

// TODO! Add errors to the json resp in case it is not nil.
func (server *Server) CreateToDo(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
		})
		return
	}
	// Unmarshal the request body into a todo item
	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
		})
		return
	}
	// Save the todo item
	todoCreated, err := todo.SaveToDo(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	// Return the created todo item
	c.JSON(http.StatusCreated, gin.H{
		"status":   http.StatusCreated,
		"response": todoCreated,
	})
}

// GetToDo is a function that gets all the tasks from the database.
// It takes in a gin.Context object as an argument and returns a JSON response.
func (server *Server) GetToDo(c *gin.Context) {
	db := server.DB.Debug().Model(models.Todo{})
	if db.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
}

// UpdateToDo is a function that updates a ToDo item in the database.
// It takes in a gin.Context object as an argument and returns a JSON response.
func (server *Server) UpdateToDo(c *gin.Context) {
	// Get the ToDo ID from the URL parameter
	todoID := c.Param("id")
	// Check if the todo id is valid
	pid, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
	//Check if the todo exist
	origToDo := models.Todo{}
	err = server.DB.Debug().Model(models.Todo{}).Where("id = ?", pid).Take(&origToDo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
		})
		return
	}
	// Read the data from the request body
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
		})
		return
	}
	// Unmarshal the data into a ToDo object
	todo := models.Todo{}
	err = json.Unmarshal(body, &todo)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
		})
		return
	}
	// Set the ToDo ID and Author ID
	todo.ID = origToDo.ID //this is important to tell the model the todo id to update, the other update field are set above
	// Update the ToDo in the database
	todoUpdated, err := todo.UpdateAToDo(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	// Return the updated ToDo
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": todoUpdated,
	})
}

// DeleteToDo is a function that deletes a ToDo item from the database.
// It takes in a gin.Context object as an argument and returns a JSON response.
func (server *Server) DeleteToDo(c *gin.Context) {
	todoID := c.Param("id")
	// Check if the ToDo ID is valid
	pid, err := strconv.ParseUint(todoID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
	fmt.Println("this is delete todo ")
	// Check if the todo exist
	todo := models.Todo{}
	err = server.DB.Debug().Model(models.Todo{}).Where("id = ?", pid).Take(&todo).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
		})
		return
	}
	// If all the conditions are met, delete the ToDo
	_, err = todo.DeleteAToDo(server.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
		})
		return
	}
	// Return a success message
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": "todo deleted",
	})
}
