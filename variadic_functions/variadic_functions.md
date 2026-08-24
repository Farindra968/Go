# Variadic Functions in Go

A **variadic function** accepts zero or more arguments of the same type. The
variadic parameter is written with an ellipsis (`...`) before its type and is
available inside the function as a slice.

## Syntax

```go
func functionName(values ...type) {
	// values has type []type inside the function.
}
```

For example, `numbers ...int` means that the caller can provide any number of
`int` values. The function can range over `numbers` like any other slice.

## Complete Example

```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func printScores(name string, scores ...int) {
	fmt.Printf("%s's scores: %v\n", name, scores)
}

func main() {
	fmt.Println("Empty sum:", sum())
	fmt.Println("Sum:", sum(1, 2, 3, 4))

	numbers := []int{5, 10, 15}
	fmt.Println("Slice sum:", sum(numbers...))

	printScores("Amina", 88, 92, 95)
}
```

Expected output:

```text
Empty sum: 0
Sum: 10
Slice sum: 30
Amina's scores: [88 92 95]
```

## Calling a Variadic Function

### Individual Arguments

Pass zero or more values separated by commas:

```go
sum()
sum(1, 2, 3)
```

When no values are supplied, the variadic parameter contains no elements. In
the example, the initial total is `0`, so `sum()` returns `0`.

### Passing a Slice

Use `...` after a slice to expand its elements into individual arguments:

```go
numbers := []int{5, 10, 15}
result := sum(numbers...)
```

The slice element type must match the variadic parameter type. A `[]int` can be
expanded for `...int`, but it cannot be passed to `...float64` without creating
and converting another slice.

## Fixed and Variadic Parameters

A function may have regular parameters before its variadic parameter:

```go
func printScores(name string, scores ...int) {
	fmt.Println(name, scores)
}
```

The variadic parameter must be the final parameter. This is valid:

```go
func logMessage(level string, messages ...string) {}
```

This is invalid because another parameter follows the variadic parameter:

```go
func invalid(values ...int, count int) {}
```

## Important Notes

- A variadic parameter has slice type inside its function body.
- A variadic function may receive zero, one, or many values.
- All variadic arguments must have the declared element type.
- Only the final parameter can be variadic.
- Use `slice...` to pass an existing slice to a variadic parameter.
- Variadic functions are useful for totals, formatting helpers, logging, and
  other operations that accept a flexible number of values.

## Running the Example

From the repository root, run:

```bash
go run ./variadic_functions/variadic_functions.go
```
