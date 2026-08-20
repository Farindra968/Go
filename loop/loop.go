package main

import "fmt"

func main() {
	// A condition-only for loop works like a while loop.
	count := 1
	for count <= 3 {
		fmt.Println("While-style:", count)
		count++
	}

	// The three-part form has an initializer, condition, and post statement.
	for number := 1; number <= 5; number++ {
		if number%2 == 0 {
			continue
		}
		fmt.Println("Odd number:", number)
	}

	// Range
	for u:= range 10 {
		fmt.Println("Character index:", u)
	}

	// range returns an index and a value while iterating over a collection.
	fruits := []string{"apple", "banana", "orange"}
	for index, fruit := range fruits {
		fmt.Println("Fruit:", index, fruit)
	}

	// break stops the nearest loop immediately.
	for number := 1; number <= 10; number++ {
		if number == 4 {
			break // exit the loop when number is 4
		}
		fmt.Println("Before break:", number)
	}

	// A loop without a condition is infinite; break gives it an exit.
	for {
		fmt.Println("Infinite loop exited safely")
		break
	}

	// A label lets break exit an outer loop from inside a nested loop.
outer:
	for row := 1; row <= 2; row++ {
		for column := 1; column <= 2; column++ {
			if row == 2 && column == 1 {
				break outer
			}
			fmt.Println("Position:", row, column)
		}
	}
}
