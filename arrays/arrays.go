package main

import "fmt"

func main() {
	var numbers [5]int

	fmt.Println("Initial array:", numbers)
	fmt.Println("Length of the array:", len(numbers))
	// Out of bounds assignment will cause a compile-time error
	for index := 0; index < len(numbers); index++ {
		numbers[index] = (index + 1) * 10 // Out of bounds assignment will cause a compile-time error = Result: 10, 20, 30, 40, 50
	}
	fmt.Println("Array after assignment:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[len(numbers)-1])

	colors := [3]string{"red", "green", "blue"}

	// Inferred length array length can be inferred from the number of elements in the array literal
	inferred := [...]int{2, 4, 6, 8}
	fmt.Println("Array literal:", colors)
	fmt.Println("Inferred length:", len(inferred), inferred)

	var total int
	for index, number := range numbers {
		total += number
		fmt.Printf("numbers[%d] = %d\n", index, number)
	}
	fmt.Println("Sum:", total)

	copyOfNumbers := numbers
	copyOfNumbers[0] = 100
	fmt.Println("Original after copy change:", numbers)
	fmt.Println("Copy after change:", copyOfNumbers)

	// Multi-dimensional arrays
	// A two-dimensional array can be thought of as an array of arrays
	matrix := [2][2]int{{1, 2}, {3, 4}} // Two-dimensional array
	fmt.Println("Two-dimensional array:", matrix)
}
