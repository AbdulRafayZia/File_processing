package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"gorilla/pkg"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)



// ResponseBody is the structure for the outgoing JSON response
type ResponseBody struct {
	TotalLines       int `json:"TotalLines"`
	TotalWords       int `json:"TotalWords"`
	TotalPuncuations int `json:"TotalPuncuations"`
	TotalVowels      int `json:"TotalVowels"`
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	

	err := r.ParseMultipartForm(10 << 20) // 10 MB max file size
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
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file from form data", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// Create a new file in the server's upload directory
	filePath := filepath.Join("./assests", handler.Filename)
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

	// fmt.Printf("data :%s\n", fileData)

	result := pkg.ProcessFile( fileData , routines)



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