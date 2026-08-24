# Closures in Go

A **closure** is a function value that can use variables from the scope where
it was created. The closure keeps access to those variables even after the
outer function has returned.

Closures are useful for preserving private state, creating customized
functions, and passing behavior to other parts of a program.

## Basic Syntax

```go
func makeFunction() func() int {
	value := 0

	return func() int {
		value++
		return value
	}
}
```

The anonymous function returned by `makeFunction` is a closure because it
captures and updates `value` from the outer function.

## Complete Example

```go
package main

import "fmt"

func newCounter(start int) func() int {
	count := start

	return func() int {
		count++
		return count
	}
}

func newMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func main() {
	firstCounter := newCounter(0)
	secondCounter := newCounter(10)

	fmt.Println("First counter:", firstCounter())
	fmt.Println("First counter:", firstCounter())
	fmt.Println("Second counter:", secondCounter())

	double := newMultiplier(2)
	triple := newMultiplier(3)
	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))
}
```

Expected output:

```text
First counter: 1
First counter: 2
Second counter: 11
Double 5: 10
Triple 5: 15
```

## How the Counter Closure Works

The `newCounter` function declares `count` and returns an anonymous function:

```go
func newCounter(start int) func() int {
	count := start

	return func() int {
		count++
		return count
	}
}
```

1. `newCounter(0)` creates `count` with an initial value of `0`.
2. The returned function captures `count`.
3. Each call increments the captured value and returns it.
4. The value is preserved between calls.

The return type, `func() int`, means that the function takes no arguments and
returns an `int`.

## Independent State

Each call to `newCounter` creates a separate captured variable:

```go
firstCounter := newCounter(0)
secondCounter := newCounter(10)
```

Calling `firstCounter` does not change `secondCounter`. They each have their
own `count` value.

## Capturing Configuration

Closures do not only preserve changing state. They can also capture a value
used as configuration:

```go
func newMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

double := newMultiplier(2)
triple := newMultiplier(3)
```

Both returned functions have the same shape, but each uses a different captured
factor.

## Closures and Function Values

A closure is a function value, so it can be stored in a variable, passed as an
argument, or returned from another function:

```go
double := newMultiplier(2)
result := double(5)
```

Here, `double` holds the closure and calling it returns `10`.

## Important Notes

- A closure can access variables from its enclosing scope.
- Captured variables remain available while the closure can still use them.
- Multiple closures created by separate calls can hold independent state.
- A closure can be used anywhere a compatible function value is accepted.
- Shared captured state can be modified by multiple closures, so concurrent use
  may require synchronization.

## Running the Example

From the repository root, run:

```bash
go run ./closures/closures.go
```
