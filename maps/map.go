package main

import (
	"fmt"
	"maps"
	"sort"
)

func main() {
	// A map stores key-value pairs. The key type must be comparable.
	ages := map[string]int{
		"Alice": 28,
		"Bob":   32,
	}
	ages["Cara"] = 25
	ages["Bob"] = 33
	fmt.Println("Ages:", ages)

	// make creates an empty, writable map with an optional initial capacity.
	students := make(map[string]bool, 3)
	students["Alice"] = true
	students["Bob"] = false
	fmt.Println("Students:", students)

	// The comma-ok form distinguishes a missing key from a stored zero value.
	age, found := ages["Alice"]
	fmt.Println("Alice age:", age, "found:", found)
	missingAge, found := ages["Drew"]
	fmt.Println("Drew age:", missingAge, "found:", found)

	// len reports the number of entries. delete is safe even when the key is absent.
	delete(ages, "Cara")
	delete(ages, "Drew")
	fmt.Println("After delete:", ages, "length:", len(ages))

	// Map iteration order is intentionally unspecified, so sort keys for stable output.
	keys := make([]string, 0, len(ages))
	for name := range ages {
		keys = append(keys, name)
	}
	sort.Strings(keys)
	for _, name := range keys {
		fmt.Printf("%s: %d\n", name, ages[name])
	}

	// A nil map can be read from and deleted from, but assigning to it panics.
	var nilMap map[string]int
	fmt.Println("Nil lookup:", nilMap["missing"], "length:", len(nilMap))
	delete(nilMap, "missing")

	// Maps can contain other maps or slices as values.
	departments := map[string][]string{
		"engineering": {"Alice", "Drew"},
		"design":      {"Bob"},
	}
	departments["engineering"] = append(departments["engineering"], "Cara")
	fmt.Println("Departments:", departments)

	// The maps package provides generic helpers in Go 1.21 and newer.
	clone := maps.Clone(ages)
	fmt.Println("Clone equal:", maps.Equal(ages, clone))
	maps.Copy(clone, map[string]int{"Drew": 29})
	fmt.Println("Copied clone:", clone)
	maps.DeleteFunc(clone, func(name string, age int) bool {
		return age < 30
	})
	fmt.Println("After DeleteFunc:", clone)
	clear(clone)
	fmt.Println("After Clear:", clone)
}
