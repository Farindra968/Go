package main

import "fmt"

// Ordered defines which types can be used with comparison logic.
// It allows integers, floating-point numbers, and strings.
type Ordered interface {
	~int | ~float64 | ~string
}

// printSlice accepts any slice type and prints each element.
func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

// max returns the larger of two values using a type constraint.
func max[T Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// mapSlice transforms each item of one type into another type.
func mapSlice[T any, U any](items []T, transform func(T) U) []U {
	result := make([]U, 0, len(items))
	for _, item := range items {
		result = append(result, transform(item))
	}
	return result
}

// Stack is a generic data structure that can hold any type.
type Stack[T any] struct {
	elements []T
}

// Push adds a value to the top of the stack.
func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

// Pop removes and returns the last element from the stack.
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

// Peek returns the top element without removing it.
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	return s.elements[len(s.elements)-1], true
}

// Len returns the number of elements currently in the stack.
func (s *Stack[T]) Len() int {
	return len(s.elements)
}

func main() {
	// Generic slices with different data types.
	integers := []int{1, 2, 3, 4, 5}
	strings := []string{"Apple", "Banana", "Cherry"}

	fmt.Println("Int slice:")
	printSlice(integers)
	fmt.Println("String slice:")
	printSlice(strings)

	// Generic function with comparison constraint.
	fmt.Println("Max int:", max(10, 30))
	fmt.Println("Max string:", max("banana", "apple"))

	// Generic function that transforms one slice into another.
	doubled := mapSlice(integers, func(n int) int {
		return n * 2
	})
	fmt.Println("Doubled numbers:", doubled)

	// Generic stack for integers.
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Stack length:", stack.Len())

	// Peek at the last item without removing it.
	top, ok := stack.Peek()
	if ok {
		fmt.Println("Top element:", top)
	}

	// Remove the top item from the stack.
	popped, ok := stack.Pop()
	if ok {
		fmt.Println("Popped element:", popped)
	}

	fmt.Println("Remaining elements:", stack.elements)
}
