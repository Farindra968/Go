# Functions in Go

A **function** is a reusable block of code that performs a specific task. Functions help organize code, reduce repetition, and improve readability.

## Function Syntax

```go
func functionName(parameter1 type, parameter2 type) returnType {
    // Function body
    return value
}
```

### Parts of a Function

- `func` — keyword used to declare a function
- `functionName` — name of the function
- Parameters — input values passed to the function
- `returnType` — type of value returned by the function
- `return` — sends a value back to the caller

## Example

```go
package main

import "fmt"

// add takes two integers and returns their sum.
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Println("Sum of 3 and 5 is:", result)
}
```

### Explanation

The `add` function accepts two integer parameters, `a` and `b`. It adds them together and returns the result as an integer.

```go
func add(a int, b int) int
```

- `add` is the function name.
- `a` and `b` are parameters.
- Both parameters have type `int`.
- The function returns an `int` value.

The function is called from `main`:

```go
result := add(3, 5)
```

Here, `3` and `5` are arguments. The function returns `8`, which is stored in `result`.

## Function Without Parameters

```go
func greet() {
    fmt.Println("Hello, Go!")
}
```

## Function Without a Return Value

```go
func displayMessage(message string) {
    fmt.Println(message)
}
```

## Function With Multiple Return Values

Go functions can return more than one value.

```go
func divide(a int, b int) (int, int) {
    quotient := a / b
    remainder := a % b

    return quotient, remainder
}
```

Usage:

```go
quotient, remainder := divide(10, 3)
fmt.Println("Quotient:", quotient)
fmt.Println("Remainder:", remainder)
```

## Named Return Values

Return values can be given names in the function declaration.

```go
func addNumbers(a int, b int) (sum int) {
    sum = a + b
    return
}
```

## Short Parameter Syntax

Parameters with the same type can be written together.

```go
func multiply(a, b int) int {
    return a * b
}
```

## Variadic Functions

A variadic function accepts zero or more values of the same type.

```go
func total(numbers ...int) int {
    sum := 0

    for _, number := range numbers {
        sum += number
    }

    return sum
}
```

Usage:

```go
fmt.Println(total(1, 2, 3, 4))
```

## Anonymous Functions

An anonymous function is a function without a name.

```go
message := func() {
    fmt.Println("Hello from an anonymous function")
}

message()
```

## Function as a Value

Functions can be assigned to variables and passed to other functions.

```go
operation := add
fmt.Println(operation(4, 6))
```

## Benefits of Functions

- Reuse code
- Reduce repetition
- Improve readability
- Simplify testing
- Separate program responsibilities
- Make programs easier to maintain

## Important Notes

- Go uses the `func` keyword to declare functions.
- The `main` function is the entry point of an executable Go program.
- Function names beginning with an uppercase letter are exported.
- Function names beginning with a lowercase letter are private to the package.
- Arguments are passed to functions by value by default.