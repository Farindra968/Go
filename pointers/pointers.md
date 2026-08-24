# Pointers in Go

A **pointer** is a value that stores the memory address of another variable.
Pointers allow a function to read or modify the original value instead of
working with a copy.

## Pointer Operators

Go uses two operators when working with pointers:

- `&` gets the address of a variable.
- `*` accesses the value stored at an address, which is called dereferencing.

For example:

```go
number := 5
pointer := &number

fmt.Println(*pointer) // 5
*pointer = 10
fmt.Println(number) // 10
```

`pointer` contains the address of `number`. Changing `*pointer` changes
`number` because both refer to the same storage.

## Complete Example

```go
package main

import "fmt"

func changeNumber(number *int) {
	*number = 10
}

func swap(first *int, second *int) {
	*first, *second = *second, *first
}

func main() {
	number := 5
	fmt.Println("Before changeNumber:", number)

	changeNumber(&number)
	fmt.Println("After changeNumber:", number)

	first, second := 3, 7
	fmt.Println("Before swap:", first, second)
	swap(&first, &second)
	fmt.Println("After swap:", first, second)
}
```

Expected output:

```text
Before changeNumber: 5
After changeNumber: 10
Before swap: 3 7
After swap: 7 3
```

## Passing a Pointer to a Function

Function arguments are passed by value in Go. When a function receives a normal
integer, it receives a copy:

```go
func changeCopy(number int) {
	number = 10
}
```

Calling `changeCopy(number)` does not change the original variable. To allow a
function to modify the original value, pass its address:

```go
func changeNumber(number *int) {
	*number = 10
}

number := 5
changeNumber(&number)
```

The parameter `number *int` means that `number` is a pointer to an `int`. The
expression `*number` accesses the actual integer stored at that address.

## Swapping Values with Pointers

The `swap` function receives pointers to two integers:

```go
func swap(first *int, second *int) {
	*first, *second = *second, *first
}
```

The right-hand side reads both original values before the assignments update
the two locations. Passing `&first` and `&second` lets the function modify the
variables in `main`.

## Nil Pointers

The zero value of a pointer is `nil`. A nil pointer does not point to a valid
value and must not be dereferenced:

```go
var pointer *int
fmt.Println(pointer == nil) // true
```

This would cause a runtime panic:

```go
fmt.Println(*pointer) // do not dereference a nil pointer
```

Check a pointer before dereferencing it when it may be nil:

```go
if pointer != nil {
	fmt.Println(*pointer)
}
```

## Pointers and Structs

Pointers are often used when a function needs to update a struct:

```go
type User struct {
	Name string
}

func rename(user *User, name string) {
	user.Name = name
}
```

Go automatically dereferences a pointer to a struct when accessing a field, so
`user.Name` is equivalent to `(*user).Name` in this example.

## Important Notes

- `&value` produces a pointer to `value`.
- `*pointer` reads or writes the value referenced by `pointer`.
- A pointer parameter lets a function modify the caller's value.
- A pointer's zero value is `nil`.
- Dereferencing a nil pointer causes a runtime panic.
- Use pointers when shared mutation or avoiding a large value copy is useful;
  use ordinary values when mutation is not needed.

## Running the Example

From the repository root, run:

```bash
go run ./pointers/pointers.go
```
