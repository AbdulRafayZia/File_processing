package login

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	filehandle "github.com/AbdulRafayZia/Gorilla-mux/fileHandle"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	fmt.Printf("The user request value %v", u)


	user, err := findUserByUsername(u.Username)
    if err != nil {
        http.Error(w, "Error finding user", http.StatusInternalServerError)
        return
    }

    // Verify the password
    if user != nil && verifyPassword(u.Password, password) {
        // Password is correct, perform login
        // You can set a session or generate a token for the logged-in user
        // Redirect the user to a dashboard or home page
        http.Redirect(w, r, "/dashboard", http.StatusFound)
        return
    }

	if u.Username == "Chek" && u.Password == "123456" {
		tokenString, err := CreateToken(u.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Errorf("No username found")
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}

func findUserByUsername(username string) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
    if err == sql.ErrNoRows {
        return nil, nil // User not found
    } else if err != nil {
        return nil, err // Other error
    }
    return &user, nil
}
func verifyPassword(hash, password string) bool {
    // Implement a secure password comparison function (e.g., bcrypt.CompareHashAndPassword)
    // Compare the stored hash with the provided password
    // Return true if they match, false otherwise
    return true
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

	
	filehandle.GetFile(w ,r)



}
