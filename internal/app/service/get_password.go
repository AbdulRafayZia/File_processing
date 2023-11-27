package service

import (
	"database/sql"
	"fmt"
	"log"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"

)

func GetPassword(username string) (string, error) {
	var hashedPassword string
	db := database.OpenDB()
	defer db.Close()
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		log.Printf("Error retrieving hashed password: %v", err)
		return "", err
	}
	return hashedPassword, nil
}