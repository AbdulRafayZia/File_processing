package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
)

func StaffLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("The user request value %v \n", request)

	role, err := service.GetRole(request.Username)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)

		return
	}
	var validRole bool
	if role == "staff" {
		validRole = true
	} else {
		validRole = false
	}
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

	validPassword := service.VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "incorrect password ", http.StatusUnauthorized)
		return

	}
	if user != nil && validPassword && validRole {

		tokenString, err := service.CreateToken(request.Username, role)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			http.Error(w, "error in generating toke ", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return

	} else {
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "invalid credentails ", http.StatusUnauthorized)
	}

}
