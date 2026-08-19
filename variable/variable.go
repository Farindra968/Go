package main

import "fmt"

// Package-level variables.
var a int = 10
var b int = 20

// Multiple variables can be declared together.
var (
	appName    = "Go Variables"
	version    = 1.0
	isComplete = true
)

func main() {
	// Explicit variable declaration.
	var age int = 25

	// Type inference.
	name := "Alice"

	// Variable declaration without an initial value.
	var score int

	// Multiple variable declaration.
	x, y := 5, 10

	// Different basic data types.
	var price float64 = 99.99
	var letter rune = 'G'
	var available bool = true

	// Constant values cannot be changed.
	const country = "India"

	fmt.Println("Value of a:", a)
	fmt.Println("Value of b:", b)
	fmt.Println("Application:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Complete:", isComplete)
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("X:", x)
	fmt.Println("Y:", y)
	fmt.Println("Price:", price)
	fmt.Println("Letter:", letter)
	fmt.Println("Available:", available)
	fmt.Println("Country:", country)

	// Variables can be changed after declaration.
	age = 26
	fmt.Println("Updated age:", age)

	// Swapping two variable values.
	x, y = y, x
	fmt.Println("After swapping - X:", x, "Y:", y)

	// A variable declared inside a block has local scope.
	{
		message := "This variable is inside a block"
		fmt.Println(message)
	}
}
