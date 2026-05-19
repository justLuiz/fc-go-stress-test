package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	urlFlag := flag.String("url", "", "URL to test")
	requestsFlag := flag.Int("requests", 0, "Total number of requests")
	concurrencyFlag := flag.Int("concurrency", 0, "Number of parallel workers")

	flag.Parse()

	if *urlFlag == "" || *requestsFlag <= 0 || *concurrencyFlag <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	startTime := time.Now()

	workChannel := make(chan struct{}, *concurrencyFlag)
	resultsChannel := make(chan int, *requestsFlag)

	var waitGroup sync.WaitGroup

	for workerIndex := 0; workerIndex < *concurrencyFlag; workerIndex++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for range workChannel {
				response, requestError := http.Get(*urlFlag)
				if requestError != nil {
					resultsChannel <- 0
					continue
				}
				statusCode := response.StatusCode
				response.Body.Close()
				resultsChannel <- statusCode
			}
		}()
	}

	for requestIndex := 0; requestIndex < *requestsFlag; requestIndex++ {
		workChannel <- struct{}{}
	}
	close(workChannel)

	waitGroup.Wait()
	close(resultsChannel)

	statusCodeCount := make(map[int]int)
	for statusCode := range resultsChannel {
		statusCodeCount[statusCode]++
	}

	totalTime := time.Since(startTime)

	totalCompleted := 0
	for _, count := range statusCodeCount {
		totalCompleted += count
	}

	fmt.Printf("Total time: %.0fs\n", totalTime.Seconds())
	fmt.Printf("Total requests: %d\n", totalCompleted)
	fmt.Printf("Requests with status 200: %d\n", statusCodeCount[200])

	for statusCode, count := range statusCodeCount {
		if statusCode != 200 {
			fmt.Printf("Status %d: %d requests\n", statusCode, count)
		}
	}
}
