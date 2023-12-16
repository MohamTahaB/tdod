package main

import (
	api "api/backend/Api"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	var db *sql.DB
	//Capture connection properties.
	cfg := mysql.Config{
		User:   "root",
		Passwd: "mypassword",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "todo",
	}
	// Handle database error
	var err error
	if db, err = sql.Open("mysql", cfg.FormatDSN()); err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	// Start the server! Happy hacking !
	api.StartServer(db)
}
