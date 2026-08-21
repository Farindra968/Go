package main

import "fmt"

func main() {
	// A nil slice has no elements and no backing array. It is ready for append.
	var numbers []int
	fmt.Println("Nil slice:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	// Slice literals and make create initialized slices.
	colors := []string{"red", "green", "blue"}
	ages := make([]int, 3, 5) // make creates a slice with length 3 and capacity 5
	ages[0], ages[1], ages[2] = 20, 30, 40
	fmt.Println("Literal:", colors)
	fmt.Println("Made slice:", ages, "len:", len(ages), "cap:", cap(ages))

	// Slices use zero-based indexes and can be changed in place.
	colors[1] = "yellow"
	fmt.Println("Updated colors:", colors)
	fmt.Println("First color:", colors[0], "last color:", colors[len(colors)-1])

	// range visits indexes and values. Use _ when the index is not needed.
	for index, color := range colors {
		fmt.Printf("colors[%d] = %s\n", index, color)
	}

	// A slice expression creates a view over the same backing array.
	firstTwo := colors[:2]
	lastTwo := colors[1:]
	fmt.Println("First two:", firstTwo)
	fmt.Println("Last two:", lastTwo)

	// append returns the updated slice and may allocate a new backing array.
	numbers = append(numbers, 10, 20, 30)
	numbers = append(numbers, 40)
	fmt.Println("After append:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	// copy copies values into an existing destination slice.
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	numbersCopy[0] = 100
	fmt.Println("Original after copy:", numbers)
	fmt.Println("Copy after change:", numbersCopy)

	// Insert and delete with append. The original slice is updated here.
	numbers = append(numbers, 0)
	copy(numbers[3+1:], numbers[3:])
	numbers[3] = 35
	fmt.Println("After insert:", numbers)
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("After delete:", numbers)

	// An empty slice is non-nil, while both nil and empty slices have length zero.
	empty := []int{}
	fmt.Println("Nil:", numbers == nil, "empty:", empty == nil)
}
