package api

import (
	"encoding/json"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/validation"
)

func StaffLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)

	role, err := service.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)

		return
	}
	validRole := service.CheckStaffRole(role)
	if !validRole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
		return

	}

	user, err := service.FindByName(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := service.GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error in getting  from db", http.StatusBadRequest)
		return
	}

	validPassword := validation.VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "incorrect password ", http.StatusUnauthorized)
		return

	}
	if user == nil && !validPassword && !validRole {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentails ", http.StatusUnauthorized)

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
