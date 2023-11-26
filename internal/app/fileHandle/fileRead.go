package filehandle

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"strconv"
	"time"

	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"

)

func GetFile(w http.ResponseWriter, r *http.Request) (utils.ResponseBody, error) {
	startTime := time.Now()
	err := r.ParseMultipartForm(10000 << 20) // 10000 MB max file size
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}

	StringRoutines := r.FormValue("routines")
	routines, err := strconv.Atoi(StringRoutines)

	if err != nil {
		http.Error(w, "Invalid routines value", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}
	fmt.Printf("routienes :%d\n", routines)

	// Get file from form data
	file, FileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return utils.ResponseBody{}, err

	}
	defer file.Close()

	var fileContent bytes.Buffer
	_, err = io.Copy(&fileContent, file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return utils.ResponseBody{}, err
	}

	// Process file
	result := ProcessFile(fileContent.String(), routines)
	
	

	endTime := time.Now()

	// Calculate the execution time
	executionTime := endTime.Sub(startTime)
	TimeInSec := executionTime.String()
	responseBody := utils.ResponseBody{
		TotalLines:       result.LineCount,
		TotalWords:       result.WordsCount,
		TotalVowels:      result.VowelsCount,
		TotalPuncuations: result.PuncuationsCount,
		ExecutionTime:    TimeInSec,
		Routines:         routines,
		Filename: FileHeader.Filename,
		


	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

	fmt.Printf("Execution time: %v\n", executionTime)
	return responseBody, nil

}

