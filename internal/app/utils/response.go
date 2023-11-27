package utils

type ResponseBody struct {
	TotalLines       int     `json:"Total_lines"`
	TotalWords       int     `json:"Total_words"`
	TotalPuncuations int     `json:"Total_puncuations"`
	TotalVowels      int     `json:"Total_vowels"`
	ExecutionTime    float64 `json:"Execution_Time"`
	Routines         int     `json:"No_of_Routines"`
	Filename         string  `json:"file_name"`
	// Id               int    `json:"id"`
	Username string `json:"username"`
};


type ExecutionData struct {
	AveragTime float64 `json:"average_execution_time"`
};