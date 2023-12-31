package controllers

import (
	"api/backend/Api/middlewares"
	"api/backend/api/models"
	"api/backend/api/seed"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Server is a struct that holds the database connection and router
type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

// Initialize is a function that initializes the server
// It takes in the database driver, user, password, port, host and name
// It returns an error if there is one
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	} else {
		fmt.Println("Unknown Driver")
	}
	//seed the db
	seed.Load(server.DB)
	//database migration
	server.DB.Debug().AutoMigrate(
		&models.Todo{},
	)
	// Initialize the router
	server.Router = gin.Default()
	// Use the CORS middleware
	server.Router.Use(middlewares.CORSMiddleware())
	server.initializeRoutes()
}

// Run is a function that runs the server
// It takes in the address of the server
// It returns an error if there is one
func (server *Server) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
