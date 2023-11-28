package database

import(
	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
	"log"
)

func GetProcesses(claims *utils.MyClaims ) []utils.ResponseBody {
	db := OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data WHERE username = $1", claims.Username)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	record := make([]utils.ResponseBody, 0)
	for rows.Next() {
		var response utils.ResponseBody
		var ProcessId int

		err := rows.Scan(&ProcessId, &response.Filename, &response.TotalWords, &response.TotalLines, &response.TotalPuncuations, &response.TotalVowels, &response.Routines, &response.Username, &response.ExecutionTime)
		if err != nil {
			log.Fatal(err)
		}
	record = append(record, response)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return record

}