package pkg

import (
	// "io/ioutil"
	
	"sync"
)
func ProcessFile( Routines int, FileData string) Summary {
	var summary Summary
	var wg sync.WaitGroup
    channal := make(chan Summary)
    // content, err := ioutil.ReadFile("D:/gorilla/assests/file.txt")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // FileData := string(content)
    chunk := len(FileData) / Routines
    startIndex := 0
    endIndex := chunk
    for iterations := 0; iterations < Routines; iterations++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			Counts(FileData[start:end], channal)
		}(startIndex, endIndex)
        // go Counts(fileData[startIndex:endIndex], channal)
        startIndex = endIndex
        endIndex += chunk
      
    }
	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(channal)
	}()

	// Receive results from the channel and aggregate them
	for counts := range channal {
		summary.LineCount += counts.LineCount
		summary.WordsCount += counts.WordsCount
		summary.VowelsCount += counts.VowelsCount
		summary.PuncuationsCount += counts.PuncuationsCount
	}
	
    // for iterations := 0; iterations < routines; iterations++ {
    //     counts := <-channal
    //     // fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.LineCount)
    //     // fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.WordsCount)
    //     // fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.VowelsCount)
    //     // fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.PuncuationsCount)
	// 	summary.LineCount+=counts.LineCount;
	// 	summary.WordsCount+=counts.WordsCount;
	// 	summary.VowelsCount+=counts.VowelsCount;
	// 	summary.PuncuationsCount+=counts.PuncuationsCount;

    // }
	return summary
}
