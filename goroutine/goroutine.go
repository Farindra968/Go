package main

import (
	"fmt"
	"sync"
	"time"
)

// task simulates a background job that runs independently.
func task(id int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Task %d finished\n", id)
}

func main() {
	fmt.Println("Main goroutine started")

	var wg sync.WaitGroup

	// Start 10 goroutines at the same time.
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			task(id)
		}(i)
	}

	// Wait until all goroutines finish before exiting the program.
	wg.Wait()
	fmt.Println("All tasks completed")
}
