package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "golang.org/x/crypto/bcrypt"

	filehandle "github.com/AbdulRafayZia/Gorilla-mux/fileHandle"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var resquest Credentials
	json.NewDecoder(r.Body).Decode(&resquest)
	fmt.Printf("The user request value %v", resquest)

	user, err := findUserByUsername(resquest.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := getHashedPassword(resquest.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return
	}

	// Verify the password
	if user != nil && verifyPassword(hashedPassword, resquest.Password) {

		tokenString, err := CreateToken(resquest.Username)
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

func findUserByUsername(username string) (*Credentials, error) {
	var user Credentials
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

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Missing authorization header")
		return
	}
	tokenString = tokenString[len("Bearer "):]

	err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	filehandle.GetFile(w, r)

}
