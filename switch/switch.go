package main

import (
	"fmt"
	"time"
)

func main() {
	grade := "B"
	switch grade {
	case "A":
		fmt.Println("Excellent!")
	case "B":
		fmt.Println("Good!")
	case "C", "D":
		fmt.Println("Average.")
	default:
		fmt.Println("Invalid grade.")
	}

	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It is the weekend.")
	default:
		fmt.Println("It is a weekday.")
	}

	switch score := 87; {
	case score >= 90:
		fmt.Println("Score: excellent")
	case score >= 60:
		fmt.Println("Score: passing")
	default:
		fmt.Println("Score: failing")
	}

	// Multiple Conditions in a Single Case
	temperature := 25
	switch {
	case temperature < 0:
		fmt.Println("It's freezing!")
	case temperature >= 0 && temperature <= 20:
		fmt.Println("It's cold.")
	case temperature > 20 && temperature <= 30:
		fmt.Println("It's warm.")
	default:
		fmt.Println("It's hot!")
	}

	// Multiple Conditions in a Single Case
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend!")
		default:
			fmt.Println("It's a weekday.")
	}

	fallthroughExample := 1
	switch fallthroughExample {
	case 1:
		fmt.Println("Fallthrough: first case")
		fallthrough
	case 2:
		fmt.Println("Fallthrough: second case")
	}

	printType("hello")
	printType(42)
	printType(true)
}

func printType(value interface{}) {
	switch value := value.(type) {
	case string:
		fmt.Printf("Type: string (%q)\n", value)
	case int:
		fmt.Printf("Type: int (%d)\n", value)
	case bool:
		fmt.Printf("Type: bool (%t)\n", value)
	default:
		fmt.Printf("Type: %T\n", value)
	}
}
