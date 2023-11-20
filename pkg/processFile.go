package pkg

import (
	// "io/ioutil"
	// "log"
	"sync"
)
func ProcessFile( fileData string , routines int) Summary {
	var summary Summary
	var wg sync.WaitGroup
    channal := make(chan Summary)
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
	
	return summary
}
