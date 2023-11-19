package pkg

import (
	// "io/ioutil"
	// "log"
	"sync"
)
func ProcessFile( fileData string , routines int) Summary {
    // routines:=1
	var summary Summary
	var wg sync.WaitGroup
    channal := make(chan Summary)
    // content, err := ioutil.ReadFile("D:/gorilla/assests/file.txt")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // fileData := string(content)
    chunk := len(fileData) / routines
    startIndex := 0
    endIndex := chunk
    for iterations := 0; iterations < routines; iterations++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			Counts(fileData[start:end], channal)
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
