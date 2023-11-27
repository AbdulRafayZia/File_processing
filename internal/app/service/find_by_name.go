package service

import (
	"database/sql"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"

	

)

func FindByName(username string) (*utils.Credentials, error) {
	var user utils.Credentials
	db := database.OpenDB()
	defer db.Close()
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err // Other error
	}
	return &user, nil
}