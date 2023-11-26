package api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func openDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

var db *sql.DB

func init() {
	// Open a connection to the database when the application starts
	db = openDB()
}
