package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"net/http"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/service"
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	database "github.com/AbdulRafayZia/Gorilla-mux/internal/infrastructure/Database"

	filehandle "github.com/AbdulRafayZia/Gorilla-mux/internal/app/fileHandle"
)

func ProcessFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString, err:=service.GetToken(w ,r )
	if tokenString == "" || err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		http.Error(w, "could not provide autherization bearer", http.StatusUnauthorized)
		return 
	}
	

	claims, err := service.VerifyToken(tokenString)
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

	stringRoutines := r.FormValue("routines")
	routines, err := strconv.Atoi(stringRoutines)

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

	response, err := filehandle.ReadFile(claims.Username, file, routines)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Not getting response")
		return
	}
	fmt.Println(claims.Username)

	endTime := time.Now()
	// Calculate the execution time
	executionTime := endTime.Sub(startTime)

	responseBody := utils.ResponseBody{
		TotalLines:       response.LineCount,
		TotalWords:       response.WordsCount,
		TotalVowels:      response.VowelsCount,
		TotalPuncuations: response.PuncuationsCount,
		ExecutionTime:    executionTime.Seconds(),
		Routines:         routines,
		Filename:         FileHeader.Filename,
		Username:         claims.Username,
	}
	//Insert response into Database
	err = database.InsertData(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Failed to INSERT file Data")
		fmt.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

	fmt.Printf("Execution time: %v\n", executionTime)

}
