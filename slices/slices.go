package main

import (
	"fmt"
	"slices"
)

func main() {
	// A nil slice has no elements and no backing array. It is ready for append.
	var numbers []int
	fmt.Println("Nil slice:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	// Slice literals and make create initialized slices.
	colors := []string{"red", "green", "blue"}
	ages := make([]int, 3, 5) // make creates a slice with length 3 and capacity 5
	ages[0], ages[1], ages[2] = 20, 30, 40
	fmt.Println("Literal:", colors)
	fmt.Println("Made slice:", ages, "len:", len(ages), "cap:", cap(ages))

	// Slices use zero-based indexes and can be changed in place.
	colors[1] = "yellow"
	fmt.Println("Updated colors:", colors)
	fmt.Println("First color:", colors[0], "last color:", colors[len(colors)-1])

	// range visits indexes and values. Use _ when the index is not needed.
	for index, color := range colors {
		fmt.Printf("colors[%d] = %s\n", index, color)
	}

	// A slice expression creates a view over the same backing array.
	firstTwo := colors[:2]
	lastTwo := colors[1:]
	fmt.Println("First two:", firstTwo)
	fmt.Println("Last two:", lastTwo)

	// append returns the updated slice and may allocate a new backing array.
	numbers = append(numbers, 10, 20, 30)
	numbers = append(numbers, 40)
	fmt.Println("After append:", numbers, "len:", len(numbers), "cap:", cap(numbers))

	// copy copies values into an existing destination slice.
	numbersCopy := make([]int, len(numbers))
	copy(numbersCopy, numbers)
	numbersCopy[0] = 100
	fmt.Println("Original after copy:", numbers)
	fmt.Println("Copy after change:", numbersCopy)

	// Insert and delete with append. The original slice is updated here.
	numbers = append(numbers, 0)
	copy(numbers[3+1:], numbers[3:])
	numbers[3] = 35
	fmt.Println("After insert:", numbers)
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println("After delete:", numbers)

	// An empty slice is non-nil, while both nil and empty slices have length zero.
	empty := []int{}
	fmt.Println("Nil:", numbers == nil, "empty:", empty == nil)
	
	var num = make([]int, 3)

	num = append(num, 2) // Output: Num: [0 0 0	 2]
	num = append(num, 3) // Output: Num: [0 0 0	 2 3]
	num = append(num, 4) // Output: Num: [0 0 0	 2 3 4]


	// Slice Operations
	var slice1 = []int{1, 2, 3, 4, 5}
	var slice2 = []int{6, 7, 8, 9, 10}
	fmt.Println("Slice2:", slice2)

	// Slice Package Functions
	// Slices.clone()
	// Clone returns a copy of the slice s. The copy is a new slice backed by a new array.
	fmt.Println("", slices.Clone(slice1)) // Output: [1 2 3 4 5]

	// Slices.contains()
	// Contains reports whether the slice s contains the value v.
	fmt.Println("", slices.Contains(slice1, 3)) // Output: true

	// Slices.delete()
	// Delete removes the element at index i from s. The order of the remaining elements is preserved.
	slice1 = slices.Delete(slice1, 2, 3) // Output: [1 2 4 5]
	fmt.Println("Slice1 after delete:", slice1)

	// Slices.insert()
	// Insert inserts the values v into s at index i. The order of the existing elements is preserved.
	slice1 = slices.Insert(slice1, 2, 3) // Output: [1 2 3 4 5]
	fmt.Println("Slice1 after insert:", slice1)

	// Slices.reverse()
	// Reverse reverses the order of the elements in s.
	slices.Reverse(slice1) // Output: [5 4 3 2 1]
	fmt.Println("Slice1 after reverse:", slice1)

	// Slices.sort()
	// Sort sorts the elements of s in ascending order.
	slices.Sort(slice1) // Output: [1 2 3 4 5]
	fmt.Println("Slice1 after sort:", slice1)

	// Slices.unique()
	// Unique returns a new slice containing the unique elements of s, in the order they first appear.
	slice3 := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	slices.Sort(slice3) // Output: [1 2 2 3 3 3 4 4 4 4]
	uniqueSlice := slices.Compact(slice3) // Output: [1 2 3 4]
	fmt.Println("Slice3 after unique:", uniqueSlice)


	// slices.grow()
	// Grow returns a slice with the same elements as s, but with a capacity of at least n.
	// If n is less than or equal to the capacity of s, Grow returns s.
	// Otherwise, Grow allocates a new slice with the same elements as s and a capacity of at least n.
	num = slices.Grow(num, 10) // Output: Num: [0 0 0 2 3 4 0 0 0 0]

	// slices.Equal()
	// Equal reports whether the two slices s and t are equal.
	slice4 := []int{1, 2, 3, 4, 5}
	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice4 and Slice5 are equal:", slices.Equal(slice4, slice5)) // Output: true



	fmt.Println("Num:", num )
}

