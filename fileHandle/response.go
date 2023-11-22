package filehandle
type ResponseBody struct {
	TotalLines       int    `json:"Total_lines"`
	TotalWords       int    `json:"Total_words"`
	TotalPuncuations int    `json:"Total_puncuations"`
	TotalVowels      int    `json:"Total_vowels"`
	ExecutionTime    string `json:"Execution_Time"`
	Routines         int    `json:"No_of_Routines"`
}