package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex
var signals = []string{"test"}

func main() {
	// go saySouthing("hello")
	// saySouthing("world")

	websites := []string{
		"https://gobyexample.com/",
		"https://www.jdn.co.il/",
		"https://www.geektime.co.il/",
		"https://www.google.co.il/",
	}

	for _, web := range websites {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

func saySouthing(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS!!")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
