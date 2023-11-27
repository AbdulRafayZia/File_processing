package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"
)

// type Users struct {
// 	Name     string `json:"name"`
// 	Password string `json: "password"`
// 	Role string `json: "role"`

// }

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	err := json.NewDecoder(r.Body).Decode(&request)
	if err!=nil{
		http.Error(w, "Unable to get data", http.StatusBadRequest)
	}
	fmt.Printf("The user request value %v", request)

	// Extract user information from the form

	// Insert the user into the database
	db:=database.OpenDB()
	defer db.Close()
	_, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", request.Username, request.Password , request.Role)
	if err != nil {
		http.Error(w, "Unable to create user", http.StatusInternalServerError)
		return
	}
	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "User created successfully")
}
