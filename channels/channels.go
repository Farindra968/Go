package main

import (
	"fmt"
	"time"
)

// Example 1: Unbuffered channel
// An unbuffered channel blocks until both sender and receiver are ready.
func addNumbers(a, b int, result chan int) {
	result <- a + b
}

// Example 2: Buffered channel
// A buffered channel can hold values without blocking until the buffer is full.
func sendMessage(msg chan string) {
	msg <- "Hello from goroutine"
}

// Example 3: Buffered channel with a goroutine
// This example demonstrates sending multiple emails using a buffered channel.
func emailSender(emails chan string, done chan bool) {
	defer func() { done <- true }()
	for email := range emails {
		fmt.Println("Email Sending:", email)
		time.Sleep(time.Second)

	}

}

// Example 4: Closing a channel and reading with range
// When a channel is closed, the loop ends automatically.
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d received job %d\n", id, job)
		time.Sleep(500 * time.Millisecond)
		results <- job * 2
	}
}

// Example 5: Select statement
// select lets a goroutine choose between multiple channel operations.
func demoSelect() {
	messages := make(chan string, 1)
	messages <- "first message"

	select {
	case msg := <-messages:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No message received yet")
	}

	select {
	case messages <- "second message":
		fmt.Println("Message sent successfully")
	default:
		fmt.Println("Channel is full, so send was skipped")
	}
}

// Example 6: Directional channels
// chan<- means only send, <-chan means only receive.
func sendData(ch chan<- string) {
	ch <- "Data sent to channel"
}

func receiveData(ch <-chan string) {
	fmt.Println("Received from channel:", <-ch)
}

func main() {
	// Example 1: Unbuffered channel
	result := make(chan int)
	go addNumbers(6, 9, result)
	fmt.Println("Sum from goroutine:", <-result)

	// Example 2: Buffered channel
	messages := make(chan string, 2)
	messages <- "Hello"
	messages <- "World"
	fmt.Println("First message:", <-messages)
	fmt.Println("Second message:", <-messages)

	// Example 3: Buffered channel with a goroutine
	// This example demonstrates sending multiple emails using a buffered channel.
	emailChann := make(chan string, 100)
	done := make(chan bool)
	go emailSender(emailChann, done)
	for i := 0; i <= 15; i++ {
		emailChann <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("Done! Email Sending")
	close(emailChann)
	<-done

	// Example 4: Close a channel and consume values using range
	jobs := make(chan int, 3)
	results := make(chan int, 3)

	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		for job := range jobs {
			results <- job * 10
		}
		close(results)
	}()

	for value := range results {
		fmt.Println("Processed value:", value)
	}

	// Example 5: select
	demoSelect()

	// Example 6: Directional channels
	data := make(chan string)
	go sendData(data)
	receiveData(data)

	// Example 7: Using a goroutine with a channel as a return value
	// This is a common pattern for communication between goroutines.
	reply := make(chan string)
	go func() {
		reply <- "Task completed successfully"
	}()
	fmt.Println(<-reply)
}
