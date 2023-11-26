package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"


	// "golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("The user request value %v \n", request)

	user, err := findUserByUsername(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := getHashedPassword(request.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return
	}

	// Verify the password
	if user != nil && verifyPassword(hashedPassword, request.Password) {

		tokenString, err := CreateToken(request.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf(" Error in Generating Token")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}

}

func findUserByUsername(username string) (*utils.Credentials, error) {
	var user utils.Credentials 
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil // User not found
	} else if err != nil {
		return nil, err // Other error
	}
	return &user, nil
}

func getHashedPassword(username string) (string, error) {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		log.Printf("Error retrieving hashed password: %v", err)
		return "", err
	}
	return hashedPassword, nil
}
func verifyPassword(hash, password string) bool {
	// Compare the stored hash with the provided password
	// err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))\

	// if err!=nil {
	// 	log.Printf("Error in verify hashed password: %v", err)
	// 	return false

	// }
	// return true
	if hash == password {
		return true
	} else {
		log.Printf("Error in verify hashed password:")
		return false
	}

}