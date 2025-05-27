package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	URL      string
	Status   string
	Duration time.Duration
}

func checkURL(url string, wg *sync.WaitGroup, results chan<- Result) {
	defer wg.Done()

	done := make(chan Result, 1)
	go func() {

		start := time.Now()
		res, err := http.Get(url)
		duration := time.Since(start)

		status := "Error"

		if err == nil {
			status = res.Status
			res.Body.Close()
		}

		done <- Result{
			URL:      url,
			Status:   status,
			Duration: duration,
		}
	}()
	
	select {
	case res := <-done:
		results <- res
	case <-time.After(2 * time.Second):
		results <- Result{
			URL:      url,
			Status:   "Timeout",
			Duration: 2 * time.Second,
		}
	}

}

// func check(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {

	urls := []string{
		"https://google.com",
		"https://golang.org",
		"https://thisurldoesnotexist.tld",
	}

	wg := &sync.WaitGroup{}
	results := make(chan Result, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, wg, results)
	}
	wg.Wait()
	close(results)

	fmt.Println("\n🔍 תוצאות הבדיקה:")
	for result := range results {
		fmt.Printf("%-30s --> %-20s (%v)\n", result.URL, result.Status, result.Duration)
	}

}
