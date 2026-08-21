# Arrays in Go

An array is a fixed-length sequence of values of the same type. Its length is
part of its type, so `[3]int` and `[4]int` are different types. Arrays are
value types: assigning an array or passing it to a function copies all of its
elements.

## Declaration and Zero Values

Declare an array with its length and element type:

```go
var numbers [5]int
```

Every element starts with the zero value for the element type. For `int`, the
zero value is `0`, so `numbers` initially contains `[0 0 0 0 0]`.

The length is fixed and can be read with the built-in `len` function:

```go
fmt.Println(len(numbers)) // 5
```

## Indexing and Assignment

Array indexes start at `0`. For an array of length `5`, valid indexes are
`0` through `4`:

```go
numbers[0] = 10
numbers[4] = 50
fmt.Println(numbers[0]) // 10
```

Using an index outside the array bounds causes a runtime panic. Go checks array
indexes, which prevents silent memory corruption.

## Array Literals

An array can be initialized with a literal:

```go
colors := [3]string{"red", "green", "blue"}
```

The number of values must fit the declared length. The `...` form lets the
compiler infer the length from the number of values:

```go
numbers := [...]int{2, 4, 6, 8}
fmt.Println(len(numbers)) // 4
```

You can also initialize selected indexes:

```go
days := [...]string{0: "Sunday", 6: "Saturday"}
```

Unspecified elements receive their type's zero value.

## Iterating Over an Array

The `range` form is the usual way to visit every element:

```go
for index, value := range numbers {
	fmt.Println(index, value)
}
```

Use `for index := range numbers` when you need to update elements in place:

```go
for index := range numbers {
	numbers[index] = (index + 1) * 10
}
```

If one value from `range` is unused, replace it with `_`:

```go
for _, value := range numbers {
	fmt.Println(value)
}
```

## Copying Arrays

Arrays are copied when assigned:

```go
copyOfNumbers := numbers
copyOfNumbers[0] = 100
```

Changing `copyOfNumbers` does not change `numbers`. This differs from slices,
which share an underlying array when copied in the usual way.

## Multidimensional Arrays

An array can contain other arrays:

```go
matrix := [2][2]int{{1, 2}, {3, 4}}
fmt.Println(matrix) // [[1 2] [3 4]]
```

The first length is the number of rows and the second is the number of columns.
An element can be accessed with two indexes, such as `matrix[1][0]`.

## Arrays as Function Arguments

Passing an array to a function passes a copy. Use a pointer when a function
must modify the original array:

```go
func reset(numbers *[5]int) {
	numbers[0] = 0
}

values := [5]int{1, 2, 3, 4, 5}
reset(&values)
```

In most everyday Go code, use a slice when the collection should have a
variable length or when avoiding a full array copy is important.

## Array Types and Comparisons

Arrays with comparable element types can be compared with `==` and `!=`.
Arrays are equal when they have the same length and every corresponding
element is equal:

```go
first := [2]int{1, 2}
second := [2]int{1, 2}
fmt.Println(first == second) // true
```

Arrays containing non-comparable elements, such as slices, cannot be compared.

## Example Output

Run [`arrays.go`](arrays.go) from the workspace root:

```bash
go run arrays/arrays.go
```

It demonstrates zero values, assignment, indexing, literals, iteration,
copying, and multidimensional arrays. Its output is:

```text
Initial array: [0 0 0 0 0]
Length of the array: 5
Array after assignment: [10 20 30 40 50]
First element: 10
Last element: 50
Array literal: [red green blue]
Inferred length: 4 [2 4 6 8]
numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
numbers[3] = 40
numbers[4] = 50
Sum: 150
Original after copy change: [10 20 30 40 50]
Copy after change: [100 20 30 40 50]
Two-dimensional array: [[1 2] [3 4]]
```
