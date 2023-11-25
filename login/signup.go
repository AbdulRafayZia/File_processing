package login

import (
	"encoding/json"
	"fmt"
	"net/http"

)
type Users struct {
	Name     string `json:"name"`
	Password string `json: "password"`
	
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request parameters
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var u Users
	err:=json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)

    // Extract user information from the form
  

    // Insert the user into the database
    _, err = db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", u.Name, u.Password)
    if err != nil {
        http.Error(w, "Unable to create user", http.StatusInternalServerError)
        return
    }

    // Respond with a success message
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "User created successfully")
}

