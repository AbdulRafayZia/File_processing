package api

import (
	"encoding/json"
	"net/http"
	"gorilla/pkg"
	"time"
	"fmt"
)

// RequestBody is the structure for the incoming JSON payload
type RequestBody struct {
	Routines int `json:"value"`
}

// ResponseBody is the structure for the outgoing JSON response
type ResponseBody struct {
	TotalLines int `json:"TotalLines"`
	TotalWords int `json:"TotalWords"`
	TotalPuncuations int `json:"TotalPuncuations"`
	TotalVowels int `json:"TotalVowels"`


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

	// Call YourFunction with the provided value
	result := pkg.ProcessFile(requestBody.Routines)

	// Create the response payload
	responseBody := ResponseBody{
		TotalLines : result.LineCount,
		TotalWords:result.WordsCount ,
		TotalVowels: result.VowelsCount,
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
