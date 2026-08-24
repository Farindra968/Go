package main

import "fmt"

// newCounter returns a function that keeps its own counter state.
func newCounter(start int) func() int {
	count := start

	return func() int {
		count++
		return count
	}
}

// newMultiplier returns a function that multiplies values by factor.
func newMultiplier(factor int) func(int) int {
	return func(value int) int {
		return value * factor
	}
}

func main() {
	firstCounter := newCounter(0)
	secondCounter := newCounter(10)

	fmt.Println("First counter:", firstCounter())
	fmt.Println("First counter:", firstCounter())
	fmt.Println("Second counter:", secondCounter())

	double := newMultiplier(2)
	triple := newMultiplier(3)
	fmt.Println("Double 5:", double(5))
	fmt.Println("Triple 5:", triple(5))
}
