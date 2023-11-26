package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"net/http"

	// "golang.org/x/crypto/bcrypt"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"

	filehandle "github.com/AbdulRafayZia/Gorilla-mux/internal/app/fileHandle"
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
	
	
	startTime := time.Now()
	err = r.ParseMultipartForm(10000 << 20) // 10000 MB max file size
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return 

	}

	StringRoutines := r.FormValue("routines")
	routines, err := strconv.Atoi(StringRoutines)

	if err != nil {
		http.Error(w, "Invalid routines value", http.StatusBadRequest)
		return 

	}
	fmt.Printf("routienes :%d\n", routines)

	// Get file from form data
	file, FileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return 

	}
	defer file.Close()

	 response ,err:=  filehandle.ReadFile(Claims.Username, file , routines)
	 if err!= nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not getting response")
		return
	 }
	 fmt.Println(Claims.Username)

	 w.WriteHeader(http.StatusOK)
	 endTime := time.Now()

	// Calculate the execution time
	executionTime := endTime.Sub(startTime)
	TimeInSec := executionTime.String()
	responseBody := utils.ResponseBody{
		TotalLines:       response.LineCount,
		TotalWords:       response.WordsCount,
		TotalVowels:      response.VowelsCount,
		TotalPuncuations: response.PuncuationsCount,
		ExecutionTime:    TimeInSec,
		Routines:         routines,
		Filename: FileHeader.Filename,
		Username: Claims.Username,
	}
	_, err = db.Exec("INSERT INTO file_processing_data( filename, words, lines, punctuations, vowels, execution_time, routines , username) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",  responseBody.Filename , responseBody.TotalWords , responseBody.TotalLines, responseBody.TotalPuncuations , responseBody.TotalVowels, responseBody.ExecutionTime, responseBody.Routines, responseBody.Username)
	 if err!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to INSERT file Data")
		fmt.Println(err)
		return 
	 }
	 
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

	fmt.Printf("Execution time: %v\n", executionTime)


}