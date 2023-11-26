package api

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/AbdulRafayZia/Gorilla-mux/utils"
	// "golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)
	

	role,err:=GetRole(request.Username)
	if err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Errorf("Unauthorize perosn")
		return
	}
	 var  validRole bool
	if role=="user"{
		validRole=true
	}else{
		validRole=false
	}
	if !validRole {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Not a Staff member  ")
		return

	}

	user, err := FindByName(request.Username)
	if err != nil {
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}
	hashedPassword, err := GetPassword(request.Username)
	if err != nil {
		http.Error(w, "error getting hashed password", http.StatusBadRequest)
		return
	}

	// Verify the password
	validPassword := VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Incorrect Password")
		return

	}
	if user != nil && validPassword {

		tokenString, err := CreateToken(request.Username , role)
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
