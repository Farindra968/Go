# Generics in Go

Generics allow you to write functions, types, and methods that work with different data types without repeating the same logic for each one.

Instead of writing separate versions for `int`, `string`, or `float64`, you can define a single generic version using a type parameter.

## Why use generics?

Generics are useful when you want to:

- write reusable logic for different types
- avoid code duplication
- build data structures like stacks, lists, or maps that work with any type
- express constraints on allowed types

## Basic syntax

A generic function uses a type parameter inside square brackets.

```go
func printSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}
```

Here:

- `T` is the type parameter
- `any` is the constraint, meaning `T` can be any type
- `[]T` means the function accepts a slice of the type parameter

## Type parameters

You can define one or more type parameters in a function or type:

```go
func pair[A, B any](first A, second B) (A, B) {
    return first, second
}
```

This function works with any two types, such as `int` and `string`.

## Type constraints

A constraint restricts what types may be used.

### `any`

`any` means any type is allowed.

```go
func identity[T any](value T) T {
    return value
}
```

### `comparable`

`comparable` allows only types that can be compared with `==` and `!=`.

```go
type KeyValue[T comparable] struct {
    Key   T
    Value string
}
```

### Custom constraints

You can also create your own constraint using an interface.

```go
type Number interface {
    ~int | ~float64
}

func add[T Number](a, b T) T {
    return a + b
}
```

This restricts `T` to integer or floating-point types.

## Generic functions

### Example: print any slice

```go
package main

import "fmt"

func printSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}

func main() {
    numbers := []int{1, 2, 3}
    names := []string{"Go", "Generics", "Example"}

    printSlice(numbers)
    printSlice(names)
}
```

### Example: compare values

```go
func max[T comparable](a, b T) T {
    if a > b {
        return a
    }
    return b
}
```

This works only when the type supports ordering. For ordered types, use `cmp.Ordered` from the standard library, which is more appropriate for comparisons like `>` and `<`.

## Generic structs

You can also make structs generic.

```go
type Box[T any] struct {
    Value T
}

func main() {
    intBox := Box[int]{Value: 42}
    stringBox := Box[string]{Value: "hello"}

    fmt.Println(intBox.Value)
    fmt.Println(stringBox.Value)
}
```

The same type can be reused for different data types without rewriting the struct.

## Generic methods

Methods can be defined on generic types as well.

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(value T) {
    s.items = append(s.items, value)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }

    lastIndex := len(s.items) - 1
    value := s.items[lastIndex]
    s.items = s.items[:lastIndex]
    return value, true
}
```

This is a common pattern for creating reusable type-safe data structures.

## Type inference

Go can often infer the type parameter automatically.

```go
numbers := []int{1, 2, 3}
printSlice(numbers)
```

You do not always need to write `printSlice[int](numbers)`.

If the compiler can infer the type from the argument, it will do so automatically.

## Example: generic stack

```go
package main

import (
    "cmp"
    "fmt"
)

type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(value T) {
    s.elements = append(s.elements, value)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }

    lastIndex := len(s.elements) - 1
    value := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.elements) == 0 {
        var zero T
        return zero, false
    }

    return s.elements[len(s.elements)-1], true
}

func max[T cmp.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

func main() {
    stack := Stack[int]{}
    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    top, ok := stack.Peek()
    if ok {
        fmt.Println("Top element:", top)
    }

    popped, ok := stack.Pop()
    if ok {
        fmt.Println("Popped:", popped)
    }

    fmt.Println("Maximum:", max(12, 7))
}
```

## Common built-in constraints

Go provides several useful constraints:

- `any` — any type
- `comparable` — types that support `==` and `!=`
- `cmp.Ordered` — ordered types such as strings and numbers

These constraints help you write clear and safe generic code.

## Benefits of generics

Generics provide several advantages:

- reusability
- type safety
- cleaner APIs
- fewer duplicate implementations
- easier maintenance

## When to use generics

Use generics when you want to write logic once for many types, especially for:

- collections
- utility functions
- algorithms
- reusable data structures

## Summary

Generics are one of Go's most powerful features. They let you write flexible code without losing type safety.

The key ideas are:

- use type parameters like `T`
- apply constraints such as `any` or `comparable`
- write generic functions, types, and methods
- let Go infer types automatically when possible

Generics make code more reusable and cleaner while keeping Go's strong type checking.
