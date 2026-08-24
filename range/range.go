package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// A traditional for loop provides manual index control.
	for index := 0; index < len(numbers); index++ {
		fmt.Println("Index:", index, "Value:", numbers[index])
	}

	// range returns the index and a copy of each slice element.
	for index, value := range numbers {
		fmt.Println("Range index:", index, "value:", value)
	}

	// Use _ when the index is not needed.
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("Sum:", sum)

	// Map iteration returns each key and value, but its order is unspecified.
	mapData := map[string]string{
		"name":    "John",
		"country": "Japan",
	}
	mapKeys := make([]string, 0, len(mapData))
	for key := range mapData {
		mapKeys = append(mapKeys, key)
	}
	sort.Strings(mapKeys)
	for _, key := range mapKeys {
		fmt.Println("Key:", key, "Value:", mapData[key])
	}

	marks := map[string]int{
		"Math":    90,
		"Eng":     85,
		"Science": 95,
		"History": 80,
	}

	totalMarks := 0
	for subject, mark := range marks {
		fmt.Println("Subject:", subject, "Mark:", mark)
		totalMarks += mark
	}
	fmt.Println("Total Marks:", totalMarks)

	// When ranging over a string, the index is a byte offset and the value is a rune.
	text := "Go café"
	for byteIndex, character := range text {
		fmt.Println("Byte index:", byteIndex, "Rune:", character, "Character:", string(character))
	}
}
