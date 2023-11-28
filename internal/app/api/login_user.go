package api

import (
	"encoding/json"
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)

	role, err := service.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)
		return
	}

	validRole := service.CheckUserRole(role)
	if !validRole {
		http.Error(w, "Not a user", http.StatusUnauthorized)
		return

	}

	user, err := service.FindByName(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := service.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return
	}

	// Verify the password
	validPassword := service.VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return

	}
	if user == nil && !validPassword {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentials ", http.StatusUnauthorized)

		return

	}
	tokenString, err := service.CreateToken(request.Username, role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		http.Error(w, "error in generating toke ", http.StatusInternalServerError)
		return
	}
	
	token:=utils.Token{
		Token: tokenString,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}
