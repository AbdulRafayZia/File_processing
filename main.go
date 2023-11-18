package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorilla/pkg"
	"io/ioutil"
	"log"

	
	"github.com/gorilla/mux"
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

// YourFunction is the function that processes the input and returns a result
func processFile( routines int) pkg.Summary {
    channal := make(chan pkg.Summary)
    content, err := ioutil.ReadFile("D:/gorilla/assests/file.txt")
    if err != nil {
        log.Fatal(err)
    }
    fileData := string(content)
    chunk := len(fileData) / routines
    startIndex := 0
    endIndex := chunk
    for iterations := 0; iterations < routines; iterations++ {
        go pkg.Counts(fileData[startIndex:endIndex], channal)
        startIndex = endIndex
        endIndex += chunk
    }
	var summary pkg.Summary
    for iterations := 0; iterations < routines; iterations++ {
        counts := <-channal
        // fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
        // fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
        // fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
        // fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
		summary.LineCount+=counts.LineCount;
		summary.WordsCount+=counts.WordsCount;
		summary.VowelsCount+=counts.VowelsCount;
		summary.PuncuationsCount+=counts.PuncuationsCount;

    }
	return summary
}

// HandlePostRequest handles the POST request and calls YourFunction
func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	// Decode the JSON payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call YourFunction with the provided value
	result := processFile(requestBody.Routines)

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
}

func main() {
	r := mux.NewRouter()

	// Define the route for the POST request
	r.HandleFunc("/api/yourendpoint", HandlePostRequest).Methods("POST")

	// Start the server
	port := 8080
	fmt.Printf("Server listening on :%d...\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
