package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gorilla/pkg"
	"io"
	// "io/ioutil"
	"net/http"
	// "os"
	// "path/filepath"
	"strconv"
	"time"
)

// ResponseBody is the structure for the outgoing JSON response
type ResponseBody struct {
	TotalLines       int           `json:"Total no of Lines"`
	TotalWords       int           `json:"Total no of Words"`
	TotalPuncuations int           `json:"Totalno of Puncuations"`
	TotalVowels      int           `json:"Total no of Vowels"`
	ExecutionTime    time.Duration `json:"ExecutionTime"`
	Routines         int           `json:"No of Routines"`
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	err := r.ParseMultipartForm(1000 << 20) // 1000 MB max file size
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
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return
	}

	fileData := buf.String()

	result := pkg.ProcessFile(fileData, routines)

	endTime := time.Now()

	// Calculate the execution time
	executionTime := endTime.Sub(startTime)
	responseBody := ResponseBody{
		TotalLines:       result.LineCount,
		TotalWords:       result.WordsCount,
		TotalVowels:      result.VowelsCount,
		TotalPuncuations: result.PuncuationsCount,
		ExecutionTime:    executionTime,
		Routines:         routines,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

	w.WriteHeader(http.StatusOK)

	fmt.Printf("Execution time: %v\n", executionTime)
	// json.NewEncoder(w).Encode(executionTime)
}
