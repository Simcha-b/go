package main

import (
	"fmt"
	"runtime"
	"time"
)

// Function that simulates heavy work - for example, downloading a file from the internet
func downloadFile(fileName string) {
	fmt.Printf("Starting to download: %s\n", fileName)
	time.Sleep(2 * time.Second) // Simulates a 2-second download
	fmt.Printf("Finished downloading: %s\n", fileName)
}

// ===== Approach 1: Without concurrency - very slow =====
func withoutConcurrency() {
	fmt.Println("=== Without concurrency ===")
	start := time.Now()

	files := []string{"file1.zip", "file2.zip", "file3.zip", "file4.zip", "file5.zip"}

	for _, file := range files {
		downloadFile(file) // Each file is loaded one after the other
	}

	fmt.Printf("Total time: %v\n\n", time.Since(start))
}

// ===== Approach 2: A goroutine for each task - dangerous! =====
func withUnlimitedGoroutines() {
	fmt.Println("=== With a goroutine for each file (dangerous!) ===")
	start := time.Now()

	// Let's assume we have 1000 files to download...
	files := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		files[i] = fmt.Sprintf("file%d.zip", i+1)
	}

	// Create a goroutine for each file - this can crash the system!
	done := make(chan bool, len(files))

	for _, file := range files {
		go func(fileName string) {
			// fmt.Printf("Goroutine #%d downloading: %s\n", runtime.NumGoroutine(), fileName)
			time.Sleep(100 * time.Millisecond) // Faster download for example
			done <- true
		}(file)
	}

	// Wait for all to finish
	for i := 0; i < len(files); i++ {
		<-done
	}

	fmt.Printf("Created %d goroutines! Total time: %v\n\n", len(files), time.Since(start))
}

// ===== Approach 3: Worker Pool - the correct solution! =====
func withWorkerPool() {
	fmt.Println("=== With Worker Pool (the correct solution!) ===")
	start := time.Now()

	// We have many files to download
	files := make([]string, 20)
	for i := 0; i < 20; i++ {
		files[i] = fmt.Sprintf("file%d.zip", i+1)
	}

	// But only 3 "downloaders" at the same time (3 workers)
	const numWorkers = 3
	jobs := make(chan string, len(files))

	// Create only 3 workers
	for w := 1; w <= numWorkers; w++ {
		go func(workerID int) {
			for fileName := range jobs {
				fmt.Printf("Worker #%d downloading: %s\n", workerID, fileName)
				time.Sleep(500 * time.Millisecond) // Download
				fmt.Printf("Worker #%d finished: %s\n", workerID, fileName)
			}
		}(w)
	}

	// Send all files to the queue
	for _, file := range files {
		jobs <- file
	}
	close(jobs) // Indicate no more files

	// Wait a bit for everything to finish
	time.Sleep(8 * time.Second)

	fmt.Printf("Used only %d workers! Total time: %v\n\n", numWorkers, time.Since(start))
}

func main() {
	fmt.Printf("Number of CPU cores: %d\n", runtime.NumCPU())
	fmt.Printf("Number of goroutines at the start: %d\n\n", runtime.NumGoroutine())

	// 1. Without concurrency - slow
	withoutConcurrency()

	// 2. With a goroutine for each task - dangerous
	fmt.Printf("Number of goroutines before: %d\n", runtime.NumGoroutine())
	withUnlimitedGoroutines()
	fmt.Printf("Number of goroutines after: %d\n\n", runtime.NumGoroutine())

	// 3. With Worker Pool - perfect!
	withWorkerPool()

	fmt.Println("=== Summary ===")
	fmt.Println("✗ Without concurrency: Very slow")
	fmt.Println("✗ Goroutine for each task: Fast but can crash the system")
	fmt.Println("✓ Worker Pool: Fast, stable, and resource-controlled!")
}
