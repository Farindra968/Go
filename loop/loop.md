# Loops in Go

Go uses the `for` keyword for every kind of loop. Go does not have separate `while` or `do-while` keywords.

## Three-Part `for` Loop

The traditional form contains an initializer, a condition, and a post statement.

```go
for initializer; condition; post {
	// loop body
}
```

Example:

```go
for number := 1; number <= 5; number++ {
	fmt.Println(number)
}
```

The `number` variable exists only inside the loop.

## Condition-Only Loop

Leaving out the initializer and post statement creates a loop that behaves like a `while` loop.

```go
count := 1
for count <= 3 {
	fmt.Println(count)
	count++
}
```

Make sure the loop condition eventually becomes false, or the loop will never stop.

## Infinite Loop

Leaving out all three parts creates an infinite loop.

```go
for {
	fmt.Println("Running")
}
```

An infinite loop can stop with `break`, a `return`, or another control mechanism.

```go
for {
	if ready {
		break
	}
}
```

## `range` Loop

Use `range` to iterate over arrays, slices, maps, strings, or channels.

```go
fruits := []string{"apple", "banana", "orange"}

for index, fruit := range fruits {
	fmt.Println(index, fruit)
}
```

For a slice, `range` provides the index and value. Use `_` when one result is not needed.

```go
for _, fruit := range fruits {
	fmt.Println(fruit)
}
```

For a map, `range` provides a key and value:

```go
scores := map[string]int{"Ana": 90, "Ben": 85}

for name, score := range scores {
	fmt.Println(name, score)
}
```

Map iteration order is not guaranteed.

## Looping Over a String

When ranging over a string, the value is a `rune`, so UTF-8 characters are handled correctly.

```go
for position, character := range "Go!" {
	fmt.Println(position, character)
}
```

The position is a byte offset, while the character is a Unicode code point.

## `break`

`break` immediately stops the nearest loop.

```go
for number := 1; number <= 10; number++ {
	if number == 4 {
		break
	}
	fmt.Println(number)
}
```

This prints `1`, `2`, and `3`.

## `continue`

`continue` skips the rest of the current iteration and starts the next one.

```go
for number := 1; number <= 5; number++ {
	if number%2 == 0 {
		continue
	}
	fmt.Println(number)
}
```

This example prints only odd numbers.

## Nested Loops and Labels

Loops can be placed inside other loops. By default, `break` affects only the nearest loop.

Use a label when the outer loop should stop:

```go
outer:
for row := 1; row <= 3; row++ {
	for column := 1; column <= 3; column++ {
		if row == 2 && column == 1 {
			break outer
		}
		fmt.Println(row, column)
	}
}
```

Labels are useful for nested-loop control, but they should be used sparingly so the control flow remains easy to follow.

## Looping With `switch`

A `switch` can be used inside a loop to handle each iteration differently.

```go
for number := 0; number < 3; number++ {
	switch number {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("another number")
	}
}
```

## Running the Example

From the Go project folder, run:

```bash
go run loop/loop.go
```

The matching example is in [loop.go](loop.go).
