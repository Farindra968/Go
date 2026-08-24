\# Maps in Go

A map is an unordered collection of key-value pairs. A map type is written as
`map[KeyType]ValueType`. Keys must be comparable, so strings, numbers, and
booleans can be keys, while slices and maps cannot.

## Creating and Updating Maps

Use a map literal when the initial values are known:

```go
ages := map[string]int{
	"Alice": 28,
	"Bob":   32,
}
ages["Cara"] = 25
ages["Bob"] = 33 // Assignment updates an existing key.
```

Use `make` to create an empty, writable map. The optional capacity is only a
hint; it does not limit the number of entries:

```go
students := make(map[string]bool, 3)
students["Alice"] = true
```

The zero value of a map is `nil`. A nil map can be read from and deleted from,
but assigning a value to it causes a runtime panic. Initialize it with `make`
before writing:

```go
var nilMap map[string]int
fmt.Println(nilMap["missing"]) // 0
delete(nilMap, "missing")     // Safe; no change.

nilMap = make(map[string]int)
nilMap["answer"] = 42
```

## Reading Values

Reading a missing key returns the zero value for the map's value type. Use the
comma-ok form when it matters whether the key exists:

```go
age, found := ages["Alice"]
fmt.Println(age, found) // 28 true

missingAge, found := ages["Drew"]
fmt.Println(missingAge, found) // 0 false
```

`len` returns the number of entries. `delete` removes a key and is safe when
the key is missing:

```go
delete(ages, "Cara")
delete(ages, "Drew")
fmt.Println(len(ages)) // 2
```

## Iterating Over Maps

Use `range` to visit keys and values. Map iteration order is not specified and
may differ between runs:

```go
for name, age := range ages {
	fmt.Println(name, age)
}
```

Sort keys first when output must be stable:

```go
keys := make([]string, 0, len(ages))
for name := range ages {
	keys = append(keys, name)
}
sort.Strings(keys)
for _, name := range keys {
	fmt.Println(name, ages[name])
}
```

## Maps as Values

Map values can be slices, structs, pointers, or other maps:

```go
departments := map[string][]string{
	"engineering": {"Alice", "Drew"},
	"design":      {"Bob"},
}
departments["engineering"] = append(departments["engineering"], "Cara")
```

When a map value is a struct, retrieve it, modify the copy, and assign it back
unless the value is a pointer. Slices and pointers stored as values can refer
to data that is modified through the map entry.

## Comparing and Copying

Maps cannot be compared with `==`, except against `nil`. Use the generic
`maps` package from Go 1.21 or newer for common operations:

```go
import "maps"

clone := maps.Clone(ages) // A shallow copy of the map.
fmt.Println(maps.Equal(ages, clone)) // true

maps.Copy(clone, map[string]int{"Drew": 29})
maps.DeleteFunc(clone, func(name string, age int) bool {
	return age < 30
})
clear(clone)
```

`maps.Copy` overwrites destination values when keys overlap. `maps.Clone` is
shallow: if values are pointers, slices, or maps, the referenced data is still
shared. The built-in `clear` function removes all entries while retaining the
map itself.

## Running the Example

Run [`map.go`](map.go) from the workspace root:

```bash
go run maps/map.go
```

The example demonstrates creation, updates, safe lookup, deletion, iteration,
nil-map behavior, nested values, and the Go 1.21 `maps` helper functions.
