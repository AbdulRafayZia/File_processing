// mypackage/summary.go
package pkg

type Summary struct {
	LineCount        int
	WordsCount       int
	VowelsCount      int
	PuncuationsCount int
}

func Counts(data string, channal chan Summary) {
	DocCounts := Summary{}

	for _, words := range data {
		switch {
		case words == '\n':
			DocCounts.LineCount++
		case words == 32:
			DocCounts.WordsCount++
		case words == 65 || words == 69 || words == 73 || words == 79 || words == 85 || words == 97 || words == 101 || words == 105 || words == 111 || words == 117:
			DocCounts.VowelsCount++
		case (words >= 33 && words <= 47) || (words >= 58 && words <= 64) || (words >= 91 && words <= 96) || (words >= 123 && words <= 126):
			DocCounts.PuncuationsCount++
		}
	}

	channal <- DocCounts
}