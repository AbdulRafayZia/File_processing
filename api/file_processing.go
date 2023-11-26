package api

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
	tokenString = tokenString[len("Bearer"):]

	Claims,err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Could not Get Claims")
		return
	}
	
	

	 response ,err:=  filehandle.GetFile(w, r)
	 if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not getting response")
		return
	 }
	 fmt.Println(Claims.Username)
	_, err = db.Exec("INSERT INTO file_processing_data( filename, words, lines, punctuations, vowels, execution_time, routines , username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",  response.Filename , response.TotalWords , response.TotalLines, response.TotalPuncuations , response.TotalVowels, response.ExecutionTime, response.Routines, Claims.Username)
	 if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to INSERT file Data")
		fmt.Println(err)
		return 
	 }
	 w.WriteHeader(http.StatusOK)


}