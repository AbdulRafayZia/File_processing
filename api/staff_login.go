package api

import (
	
	"encoding/json"
	"fmt"
	
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/utils"


	// "golang.org/x/crypto/bcrypt"
)


func StaffLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Printf("The request body is %v\n", r.Body)

	var request utils.Credentials
	json.NewDecoder(r.Body).Decode(&request)
	fmt.Printf("The user request value %v \n", request)

	

	role,err:=GetRole(request.Username)
	if err!=nil{
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "Unauthozied username", http.StatusUnauthorized)

		return
	}
	 var  validRole bool
	if role=="staff"{
		validRole=true
	}else{
		validRole=false
	}
	if !validRole {
		http.Error(w, "Not a Staff Member", http.StatusUnauthorized)
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


	validPassword := VerifyPassword(hashedPassword, request.Password)
	if !validPassword {

		w.WriteHeader(http.StatusBadRequest)
		fmt.Errorf("Incorrect Password")
		return

	}
	if user != nil && validPassword && validRole {

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



