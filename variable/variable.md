# Variables in Go

A **variable** is a named storage location used to hold a value. The value of a variable can usually be changed while the program is running.

Variables allow programs to store and work with data such as numbers, text, and Boolean values.

## Declaring a Variable

The `var` keyword is used to declare a variable.

```go
var age int = 25
```

In this example:

- `var` declares the variable.
- `age` is the variable name.
- `int` is the data type.
- `25` is the initial value.

## Type Inference

Go can determine the type automatically when a value is assigned.

```go
name := "Alice"
age := 25
price := 99.99
active := true
```

The `:=` operator declares and initializes a new variable.

> Short declarations using `:=` can be used only inside functions.

## Explicit Declaration

A variable can be declared with its type and value.

```go
var number int = 100
var message string = "Hello, Go!"
```

## Declaration Without an Initial Value

If a variable is declared without a value, Go assigns its **zero value**.

```go
var score int
var name string
var active bool
```

The zero values are:

| Type | Zero Value |
|---|---|
| `int` | `0` |
| `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |
| Pointer, map, slice, function, interface | `nil` |

## Basic Data Types

```go
var count int = 10
var price float64 = 19.99
var grade rune = 'A'
var name string = "Alice"
var available bool = true
```

Common types include:

- `int` — whole numbers
- `float64` — decimal numbers
- `string` — text
- `bool` — `true` or `false`
- `rune` — a Unicode character
- `byte` — an alias for `uint8`

## Multiple Variable Declaration

Several variables can be declared in one statement.

```go
var x, y int = 10, 20
```

Variables can also be declared using a grouped block.

```go
var (
    name    = "Go"
    version = 1.0
    active  = true
)
```

## Multiple Assignment

Several variables can be assigned at the same time.

```go
x, y := 5, 10
x, y = y, x
```

The values of `x` and `y` are swapped without requiring a temporary variable.

## Changing a Variable

A variable declared with `var` can be reassigned.

```go
age := 25
age = 26
```

The new value must be compatible with the variable's type.

```go
var count int = 10
count = 20
```

## Constants

A constant is a fixed value that cannot be changed after declaration.

```go
const country = "India"
const pi = 3.14159
```

Constants are declared using the `const` keyword.

```go
const daysInWeek int = 7
```

This is invalid:

```go
const number = 10
number = 20 // Error: cannot assign to a constant
```

## Variable Scope

**Scope** defines where a variable can be accessed.

### Package Scope

A variable declared outside all functions has package scope.

```go
package main

import "fmt"

var message = "Hello"

func main() {
    fmt.Println(message)
}
```

### Function Scope

A variable declared inside a function can be used only inside that function.

```go
func main() {
    name := "Alice"
    fmt.Println(name)
}
```

### Block Scope

A variable declared inside `{}` is available only inside that block.

```go
func main() {
    {
        message := "Inside the block"
        fmt.Println(message)
    }
}
```

## Variable Shadowing

Variable shadowing occurs when a local variable has the same name as a variable from an outer scope.

```go
message := "Global message"

func main() {
    message := "Local message"
    fmt.Println(message)
}
```

The local variable is used inside `main`.

## Unused Variables

Go does not allow declared local variables to remain unused.

```go
func main() {
    name := "Alice"
    fmt.Println(name)
}
```

If `name` is declared but never used, the compiler reports an error.

Package-level variables may be declared without being used.

## Type Conversion

Go does not automatically convert between different numeric types. Explicit conversion is required.

```go
var number int = 10
var decimal float64 = float64(number)

fmt.Println(decimal)
```

Another example:

```go
var price float64 = 25.75
var wholeNumber int = int(price)

fmt.Println(wholeNumber)
```

The decimal part is removed during conversion.

## Complete Example

```go
package main

import "fmt"

var a int = 10
var b int = 20

func main() {
    name := "Alice"
    age := 25
    price := 99.99
    active := true

    fmt.Println("A:", a)
    fmt.Println("B:", b)
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Price:", price)
    fmt.Println("Active:", active)

    age = 26
    fmt.Println("Updated age:", age)
}
```

## Important Rules

- Use `var` for standard variable declarations.
- Use `:=` for short declarations inside functions.
- A variable's value can be changed.
- A constant's value cannot be changed.
- Variables must be used if declared inside a function.
- Go requires explicit type conversion.
- Variable names beginning with uppercase letters are exported.
- Variable names beginning with lowercase letters are private.
- Use meaningful names to make code easier to understand.