package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// Example: 1. Basic error handling using return values.
	fmt.Println("=== Basic error handling ===")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// Example: 2. Creating custom errors with errors.New.
	fmt.Println("\n=== Custom error with errors.New ===")
	customErr := errors.New("invalid username")
	fmt.Println(customErr)

	// Example: 3. Wrapping errors with additional context.
	fmt.Println("\n=== Error wrapping ===")
	wrappedErr := fmt.Errorf("database query failed: %w", customErr)
	fmt.Println(wrappedErr)
	fmt.Println("Unwrap:", errors.Unwrap(wrappedErr))

	// Example: 4. Validating user input and returning errors.
	fmt.Println("\n=== Validating input ===")
	if err := validateAge(15); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	if err := validateAge(25); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	// Example: 5. Handling file open errors.
	fmt.Println("\n=== File handling error example ===")
	_, err = os.Open("missing-file.txt")
	if err != nil {
		fmt.Println("File open error:", err)
	}

	// Example: 6. Recovering from a panic.
	fmt.Println("\n=== Recovering from panic ===")
	recoverFromPanic()

	// Example: 7. Returning multiple values and parsing with errors.
	fmt.Println("\n=== Returning multiple values ===")
	n, err := parseNumber("123")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed number:", n)
	}

	n, err = parseNumber("abc")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed number:", n)
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d is too young: must be at least 18", age)
	}
	return nil
}

func recoverFromPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered panic:", r)
		}
	}()

	panic("something went wrong")
}

func parseNumber(value string) (int, error) {
	var number int
	_, err := fmt.Sscanf(value, "%d", &number)
	if err != nil {
		return 0, fmt.Errorf("invalid number %q: %w", value, err)
	}
	return number, nil
}
