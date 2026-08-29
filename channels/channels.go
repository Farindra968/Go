package main

import (
	"fmt"
	"time"
)

func processNum(chanNum chan int) {
	fmt.Println("Channel Number:", <-chanNum)
} 

func main() {
	channelNum := make(chan int)
	go processNum(channelNum)

	channelNum <- 42

	time.Sleep(time.Second*2) // Wait for the goroutine to finish

	// channelMsg := make(chan string)
	// channelMsg <- "Hello, World!"
	// msg := <-channelMsg

	// fmt.Println("Received Msg from channelMsg:", msg)
}