package main

import "fmt"

// add function that takes two integers and returns their sum
func add(a int, b int) int {
	return a + b
}

// Named return values
// Name the return values in the function signature. This allows you to return values without explicitly specifying them in the return statement.
func split(sum int) (x, y int) {
	x = sum * 5 / 3
	y = sum + x
	return
}

func main() {
	fmt.Println("Sum of 3 and 5 is:", add(3, 5))
	fmt.Println(split(17))


}
