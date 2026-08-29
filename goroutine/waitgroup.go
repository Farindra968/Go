package main

import (
	"fmt"
	"sync"
)

func task1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Task %d started\n", id)
}

// func main1() {
// 	var wg sync.WaitGroup
// 	for i:=0; i<=20; i++ {
// 		wg.Add(1)
// 		task1(i, &wg)
// 	}
// 	wg.Wait()
// }