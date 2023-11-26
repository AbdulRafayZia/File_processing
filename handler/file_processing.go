package handler

import (
	
	"fmt"
	
	"net/http"

	// "golang.org/x/crypto/bcrypt"

	filehandle "github.com/AbdulRafayZia/Gorilla-mux/fileHandle"
)
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