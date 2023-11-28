package database

import (


	"github.com/AbdulRafayZia/Gorilla-mux/internal/app/utils"
)

func GetProcessesById(claims *utils.MyClaims, id string) ([]utils.ResponseBody , error) {
	db := OpenDB()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM file_processing_data WHERE username = $1 AND id=$2", claims.Username, id)
	if err != nil {
		return nil , err
	}

	defer rows.Close()
	record := make([]utils.ResponseBody, 0)
	for rows.Next() {
		var response utils.ResponseBody
		var ProcessId int

		err := rows.Scan(&ProcessId, &response.Filename, &response.TotalWords, &response.TotalLines, &response.TotalPuncuations, &response.TotalVowels, &response.Routines, &response.Username, &response.ExecutionTime)
		if err != nil {
			return nil , err
		}
		record = append(record, response)

	}
	if err := rows.Err(); err != nil {
		return nil , err
	}
	return record,nil

}
