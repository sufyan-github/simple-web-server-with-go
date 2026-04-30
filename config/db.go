package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	connStr := "user=postgres password=1234 dbname=go_db sslmode=disable"
	
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}

	DB = db
	fmt.Println("Successfully connected to the database!")
}