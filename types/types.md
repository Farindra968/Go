# Data Types in Go

A **data type** defines the kind of value a variable can store and the operations that can be performed on that value.

Go is a **statically typed** language. This means that a variable's type is checked at compile time.

```go
var age int = 25
var name string = "Alice"
```

Once a variable has a type, it normally cannot store a value of another type without explicit conversion.

---

## Basic Types

### Integer Types

Integer types store whole numbers.

| Type | Description |
|---|---|
| `int` | Integer whose size depends on the system architecture |
| `int8` | Signed 8-bit integer |
| `int16` | Signed 16-bit integer |
| `int32` | Signed 32-bit integer |
| `int64` | Signed 64-bit integer |
| `uint` | Unsigned integer |
| `uint8` | Unsigned 8-bit integer |
| `uint16` | Unsigned 16-bit integer |
| `uint32` | Unsigned 32-bit integer |
| `uint64` | Unsigned 64-bit integer |
| `uintptr` | Integer large enough to hold a pointer address |

Signed integers can contain negative and positive values. Unsigned integers can contain only zero and positive values.

```go
var count int = 100
var temperature int8 = -10
var fileSize uint64 = 5000
```

### `byte`

`byte` is an alias for `uint8`. It is commonly used for raw data and ASCII characters.

```go
var letter byte = 'A'
data := []byte("Hello")
```

### `rune`

`rune` is an alias for `int32`. It represents a Unicode code point and is useful for characters from different languages.

```go
var character rune = '世'
fmt.Println(character)
```

A string may contain multiple Unicode characters, while a rune represents one character.

---

## Floating-Point Types

Floating-point types store decimal values.

| Type | Description |
|---|---|
| `float32` | Single-precision decimal number |
| `float64` | Double-precision decimal number |

`float64` is generally preferred when greater precision is required.

```go
var temperature float32 = 36.5
var price float64 = 99.99
```

Floating-point calculations may contain small rounding differences because decimal values are represented in binary.

---

## Complex Types

Complex numbers contain a real part and an imaginary part.

| Type | Description |
|---|---|
| `complex64` | Real and imaginary parts use `float32` |
| `complex128` | Real and imaginary parts use `float64` |

```go
var number complex128 = 3 + 4i
```

The built-in functions `real` and `imag` retrieve the two parts.

```go
realPart := real(number)
imaginaryPart := imag(number)
```

---

## Boolean Type

The `bool` type stores one of two values:

- `true`
- `false`

```go
var loggedIn bool = true

if loggedIn {
    fmt.Println("User is logged in")
}
```

The zero value of `bool` is `false`.

---

## String Type

A `string` is an immutable sequence of bytes, usually containing UTF-8 text.

```go
message := "Hello, Go!"
```

Strings can be combined using the `+` operator.

```go
firstName := "Alice"
lastName := "Smith"
fullName := firstName + " " + lastName
```

Useful operations include:

```go
length := len(message)
firstByte := message[0]
```

`len` returns the number of bytes, not necessarily the number of Unicode characters.

For Unicode character processing, convert the string to a rune slice:

```go
characters := []rune("世界")
fmt.Println(len(characters))
```

---

## Zero Values

When a variable is declared without an initial value, Go assigns its zero value.

| Type | Zero Value |
|---|---|
| Numeric types | `0` |
| `float32`, `float64` | `0` |
| `bool` | `false` |
| `string` | `""` |
| Pointer | `nil` |
| Slice | `nil` |
| Map | `nil` |
| Function | `nil` |
| Interface | `nil` |
| Channel | `nil` |

```go
var number int
var name string
var active bool

fmt.Println(number) // 0
fmt.Println(name)   // ""
fmt.Println(active) // false
```

---

## Arrays

An array is a fixed-length collection of values with the same type.

```go
numbers := [3]int{10, 20, 30}
```

The array length is part of its type. Therefore, `[3]int` and `[4]int` are different types.

```go
var first [3]int
var second [4]int
```

Arrays are useful when the number of elements is known and fixed.

---

## Slices

A slice is a flexible, dynamically sized collection built on top of an array.

```go
scores := []int{80, 90, 95}
scores = append(scores, 100)
```

Slices have:

- A length, returned by `len`
- A capacity, returned by `cap`
- A reference to an underlying array

```go
fmt.Println(len(scores))
fmt.Println(cap(scores))
```

A slice can be created from an array:

```go
numbers := [5]int{1, 2, 3, 4, 5}
part := numbers[1:4]
```

---

## Maps

A map stores values using unique keys.

```go
students := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

Values are accessed by their keys:

```go
score := students["Alice"]
```

The comma-ok syntax checks whether a key exists:

```go
score, exists := students["Charlie"]
if exists {
    fmt.Println(score)
}
```

A map can be updated and deleted:

```go
students["Charlie"] = 88
delete(students, "Bob")
```

---

## Structs

A struct groups related fields into a single type.

```go
type Person struct {
    Name string
    Age  int
}

person := Person{
    Name: "Alice",
    Age:  25,
}
```

Fields are accessed using the dot operator:

```go
fmt.Println(person.Name)
fmt.Println(person.Age)
```

Structs are commonly used to represent users, products, orders, and configuration objects.

---

## Pointers

A pointer stores the memory address of another variable.

```go
age := 25
pointer := &age

fmt.Println(*pointer)
```

- `&age` gets the address of `age`.
- `*pointer` accesses the value stored at that address.

A pointer can modify the original variable:

```go
*pointer = 26
```

The zero value of a pointer is `nil`.

---

## Functions as Types

Functions can be stored in variables, passed as arguments, and returned from other functions.

```go
add := func(a int, b int) int {
    return a + b
}

result := add(3, 5)
```

A function type can be declared as follows:

```go
type Operation func(int, int) int
```

---

## Interfaces

An interface defines a set of method signatures. A type satisfies an interface by implementing all its methods.

```go
type Stringer interface {
    String() string
}
```

Any type with a `String() string` method satisfies this interface.

```go
type Product struct {
    Name  string
    Price float64
}

func (p Product) String() string {
    return p.Name
}
```

Interfaces allow code to work with behavior instead of a specific concrete type.

---

## Empty Interface and `any`

The empty interface has no methods, so every Go type satisfies it.

```go
var value interface{} = 42
```

Go also provides `any` as an alias for `interface{}`.

```go
var data any = "Hello"
```

Because the concrete type is unknown, a type assertion or type switch can be used.

```go
text, ok := data.(string)
if ok {
    fmt.Println(text)
}
```

A type switch handles multiple possible types:

```go
switch value := data.(type) {
case string:
    fmt.Println("String:", value)
case int:
    fmt.Println("Integer:", value)
default:
    fmt.Println("Unknown type")
}
```

---

## Channels

A channel allows goroutines to communicate and exchange values safely.

```go
messages := make(chan string)

go func() {
    messages <- "Task completed"
}()

message := <-messages
fmt.Println(message)
```

A channel has a type, which defines the kind of values it can transport.

```go
numbers := make(chan int)
```

Channels may be buffered:

```go
buffered := make(chan int, 2)
buffered <- 10
buffered <- 20
```

---

## Type Conversion

Go requires explicit conversion between compatible types.

```go
number := 10
decimal := float64(number)
```

Another example:

```go
price := 19.99
wholeNumber := int(price)
```

The conversion from `float64` to `int` removes the decimal part.

Values cannot always be converted safely. For example, converting a string directly to an integer requires a package such as `strconv`.

```go
number, err := strconv.Atoi("123")
if err != nil {
    fmt.Println("Conversion error:", err)
}
```

---

## Type Aliases

A type alias gives another name to an existing type. The alias remains the same type.

```go
type ID = int
```

`ID` and `int` are identical types.

---

## Defined Types

A defined type creates a new, distinct type based on an existing type.

```go
type Age int

var userAge Age = 25
```

`Age` and `int` are different types, even though they have the same underlying representation.

Defined types improve readability and type safety.

---

## Type Inference

Go can determine a variable's type from its initial value.

```go
number := 10       // int
price := 19.99     // float64
name := "Alice"    // string
active := true     // bool
```

The inferred type is fixed after declaration.

---

## Choosing the Correct Type

- Use `int` for normal whole-number calculations.
- Use `int64` when a specific integer size is required.
- Use `float64` for decimal calculations.
- Use `bool` for true/false conditions.
- Use `string` for text.
- Use `byte` for raw bytes.
- Use `rune` for Unicode characters.
- Use slices for flexible collections.
- Use arrays for fixed-size collections.
- Use maps for key-value data.
- Use structs for related fields.
- Use interfaces for shared behavior.
- Use pointers when a function must modify the original value.

## Summary

Go provides simple primitive types such as `int`, `string`, `float64`, and `bool`, as well as powerful composite types such as arrays, slices, maps, structs, pointers, functions, interfaces, and channels.

Understanding Go's type system helps create code that is readable, safe, efficient, and easy to maintain.