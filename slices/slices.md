# Slices in Go

A slice is a flexible, variable-length view of an array. A slice contains a
pointer to an underlying array, a length, and a capacity. The length is the
number of elements currently in the slice. The capacity is the number of
elements available from the first element of the slice through its backing
array.

Unlike arrays, slices do not include their length in their type. `[]int` is the
type for every slice of integers, regardless of its length.

## Declaring Slices

The zero value of a slice is `nil`:

```go
var numbers []int
fmt.Println(numbers)    // []
fmt.Println(len(numbers)) // 0
fmt.Println(cap(numbers)) // 0
```

A nil slice is safe to read, pass to `len` and `cap`, and append to. A slice
literal creates a slice with initial values:

```go
colors := []string{"red", "green", "blue"}
```

Use `make` when you want to specify a length or capacity:

```go
ages := make([]int, 3, 5)
```

This creates a slice with length `3` and capacity `5`. Its three elements are
initially zero. If capacity is omitted, it is at least the requested length.

## Length and Capacity

The built-in functions `len` and `cap` inspect a slice:

```go
values := make([]int, 2, 4)
fmt.Println(len(values)) // 2
fmt.Println(cap(values)) // 4
```

Capacity is an implementation detail that helps describe when `append` may
need to allocate a new backing array. Code should rely on the slice returned by
`append`, not on a particular capacity-growth strategy.

## Indexing and Iteration

Slice indexes start at zero. Assigning through an index changes the element:

```go
numbers := []int{10, 20, 30}
numbers[1] = 25
fmt.Println(numbers[0], numbers[len(numbers)-1]) // 10 30
```

Use `range` to iterate over indexes and values:

```go
for index, value := range numbers {
	fmt.Println(index, value)
}

for _, value := range numbers {
	fmt.Println(value)
}
```

Indexing outside the range causes a runtime panic, so check the length before
accessing an index when the input may be empty.

## Slicing a Slice

A slice expression uses the half-open range `[start:end]`: it includes `start`
and excludes `end`.

```go
numbers := []int{10, 20, 30, 40, 50}
middle := numbers[1:4]
fmt.Println(middle) // [20 30 40]
```

The resulting slice usually shares the same backing array. Changing an element
through one slice can therefore change the other:

```go
middle[0] = 99
fmt.Println(numbers) // [10 99 30 40 50]
```

The full slice expression `numbers[start:end:max]` can limit the capacity of a
subslice and prevent an append from overwriting elements beyond `end`.

## Adding Elements with `append`

`append` returns a slice containing the original elements followed by the new
ones. Always assign its result:

```go
numbers := []int{1, 2}
numbers = append(numbers, 3, 4)
moreNumbers := []int{5, 6}
numbers = append(numbers, moreNumbers...)
```

When the current capacity is insufficient, Go allocates a new backing array.
This is why keeping the returned slice is required.

## Copying Slices

Assigning a slice copies only its slice header, so the two slices normally share
elements. Use `copy` when an independent copy is needed:

```go
source := []int{1, 2, 3}
destination := make([]int, len(source))
copy(destination, source)
destination[0] = 100
fmt.Println(source)      // [1 2 3]
fmt.Println(destination) // [100 2 3]
```

`copy` returns the number of elements copied. It copies the smaller of the two
slice lengths.

## Inserting and Deleting

There is no built-in insert or delete operation. `append` and `copy` provide
the usual patterns:

```go
numbers := []int{10, 20, 40}
numbers = append(numbers, 0)
copy(numbers[3:], numbers[2:])
numbers[2] = 30
```

Delete the element at index `2` with:

```go
numbers = append(numbers[:2], numbers[3:]...)
```

These operations may reuse the backing array. If the removed value contains
pointers or other references and the slice remains long-lived, clear the
removed slot first when retaining the backing array matters.

## Nil and Empty Slices

Both nil and empty slices have length zero, but only the nil slice compares
equal to `nil`:

```go
var nilSlice []int
emptySlice := []int{}

fmt.Println(nilSlice == nil)   // true
fmt.Println(emptySlice == nil) // false
```

Slices cannot be compared with `==` except against `nil`. Use
`slices.Equal` from the standard library when comparing elements in Go 1.21 or
newer:

```go
import "slices"

slices.Equal([]int{1, 2}, []int{1, 2}) // true
```

## Passing Slices to Functions

A slice passed to a function is passed by value, but its header still points to
the same backing array. A function can change existing elements. If it uses
`append`, return the updated slice to the caller:

```go
func addValue(numbers []int, value int) []int {
	numbers = append(numbers, value)
	return numbers
}
```

## Running the Example

Run [`slices.go`](slices.go) from the workspace root:

```bash
go run slices/slices.go
```

The example demonstrates nil and empty slices, literals, `make`, indexing,
iteration, slicing, `append`, `copy`, insertion, and deletion.
