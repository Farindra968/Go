package main

import "fmt"

// sum returns the total of zero or more integers.
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// printScores prints a student's name and any number of scores.
func printScores(name string, scores ...int) {
	fmt.Printf("%s's scores: %v\n", name, scores)
}

func main() {
	// A variadic function can be called with no arguments.
	fmt.Println("Empty sum:", sum())

	// It can also receive any number of individual arguments.
	fmt.Println("Sum:", sum(1, 2, 3, 4))

	// A slice can be expanded into individual arguments with ... .
	numbers := []int{5, 10, 15}
	fmt.Println("Slice sum:", sum(numbers...))

	// A regular parameter must come before the variadic parameter.
	printScores("Amina", 88, 92, 95)
}
