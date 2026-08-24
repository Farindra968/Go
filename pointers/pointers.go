package main

import "fmt"

// changeNumber updates the integer stored at the address in number.
func changeNumber(number *int) {
	*number = 10
}

// swap exchanges the values stored at the addresses in first and second.
func swap(first *int, second *int) {
	*first, *second = *second, *first
}

func main() {
	number := 5
	fmt.Println("Before changeNumber:", number)

	changeNumber(&number)
	fmt.Println("After changeNumber:", number)

	first, second := 3, 7
	fmt.Println("Before swap:", first, second)
	swap(&first, &second)
	fmt.Println("After swap:", first, second)
}
