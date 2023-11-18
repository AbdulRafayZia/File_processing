package api

import (
	"encoding/json"
	"fmt"
	"gorilla/pkg"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
	"os"
	"io"
)

// RequestBody is the structure for the incoming JSON payload
type RequestBody struct {
	Routines int `json:"value"`
}

// ResponseBody is the structure for the outgoing JSON response
type ResponseBody struct {
	TotalLines       int `json:"TotalLines"`
	TotalWords       int `json:"TotalWords"`
	TotalPuncuations int `json:"TotalPuncuations"`
	TotalVowels      int `json:"TotalVowels"`
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	var requestBody RequestBody

	// Decode the JSON payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}


	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error while getting the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	// Create a new file in the server's upload directory
	filePath := filepath.Join("uploads", handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create the file for writing", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

		// Copy the file content to the new file
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Unable to write the file", http.StatusInternalServerError)
			return
		}
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	fileData := string(fileBytes)

	// Call YourFunction with the provided value
	result := pkg.ProcessFile(requestBody.Routines , fileData)

	// Create the response payload
	responseBody := ResponseBody{
		TotalLines:       result.LineCount,
		TotalWords:       result.WordsCount,
		TotalVowels:      result.VowelsCount,
		TotalPuncuations: result.PuncuationsCount,
	}

	// Encode and send the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseBody)
	endTime := time.Now()

	// Calculate the execution time
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution time: %v\n", executionTime)
}
