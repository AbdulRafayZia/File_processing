	package api

	import (
		"encoding/json"
		"fmt"
		"gorilla/pkg"
		"io"
		"net/http"
		"os"
		"time"
	)

	// RequestBody is the structure for the incoming JSON payload
	type RequestBody struct {
		Routines int `json:"routines"`
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
		 
		routines:=GetJsonData(w , r)
			// Parse the multipart form data
		// err := r.ParseMultipartForm(10 << 20) // 10 MB limit
		// if err != nil {
		// 	http.Error(w, "Unable to parse form", http.StatusBadRequest)
		// 	return
		// }

		// Get the file from the request	
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		// Create a temporary file to store the uploaded file
		tempFile, err := os.CreateTemp("uploads", "temp-*.txt")
		if err != nil {
			http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		// Copy the file data to the temporary file
		_, err = io.Copy(tempFile, file)
		if err != nil {
			http.Error(w, "Error copying file data", http.StatusInternalServerError)
			return
		}

		// Read the contents of the temporary file
		contents, err := os.ReadFile(tempFile.Name())
		if err != nil {
			http.Error(w, "Error reading file contents", http.StatusInternalServerError)
			return
		}
		fileData:=  string(contents)
		// Call YourFunction with the provided value
		// result := pkg.ProcessFile(fileData)
		result := pkg.ProcessFile(fileData , routines)

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
func GetJsonData(w http.ResponseWriter, r  *http.Request) int  {
	var requestBody RequestBody

		// Decode the JSON payload
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&requestBody)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return 0
		}
		return requestBody.Routines
	
}