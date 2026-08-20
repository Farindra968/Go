# Constants in Go

A **constant** is a value that is known at compile time and cannot be changed while the program runs. Constants are declared with the `const` keyword.

```go
const pi float64 = 3.14
const language = "Go"
```

## Declaring Constants

Constants can be declared with an explicit type or with type inference.

```go
const daysInWeek int = 7
const greeting = "Hello, Go!"
```

Use an explicit type when the constant should have a specific type. An inferred constant is initially **untyped**, so Go can use it with compatible numeric types when needed.

```go
const number = 10 // untyped integer constant
const price float64 = 19.99
```

## Constant Expressions

The value of a constant can be calculated from other constants and constant expressions.

```go
const minutesInHour = 60
const secondsInHour = minutesInHour * 60
const radius = 5
const pi = 3.14159
const area = pi * radius * radius
```

The expression must be computable at compile time. Values returned by functions, such as `time.Now()`, cannot be constants.

## Grouped Constants

Several constants can be declared in one `const` block.

```go
const (
	appName = "Go Constants"
	version = 1
	stable  = true
)
```

Grouping related constants makes them easier to read and maintain.

## `iota`

Inside a constant block, `iota` starts at `0` and increases by one for each constant specification.

```go
const (
	sunday = iota // 0
	monday        // 1
	tuesday       // 2
)
```

This is useful for sequential values such as days, levels, or bit flags.

```go
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

Each constant in this example receives the previous expression, `iota`, with the next value.

## Typed and Untyped Constants

A typed constant has a specific type.

```go
const typedNumber int = 10
```

An untyped constant does not have a fixed Go type until it is used in a context that requires one.

```go
const untypedNumber = 10

var integer int = untypedNumber
var decimal float64 = untypedNumber
```

The value must still be representable by the destination type.

## Constants Cannot Be Changed

After declaration, a constant cannot be assigned a new value.

```go
const maxUsers = 100
// maxUsers = 200 // Compile error: cannot assign to maxUsers
```

The commented line is invalid and is included only to show the rule. Uncommenting it prevents the program from compiling.

## Constants and Variables

Use a constant for a value that should never change. Use a variable when the value needs to change during execution.

```go
const taxRate = 0.18
var total = 100.0

total = total + total*taxRate
```

## Running the Example

From the Go project folder, run:

```bash
go run constants/constants.go
```

The matching example is in [constants.go](constants.go).
