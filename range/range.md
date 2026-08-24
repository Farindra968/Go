# The `range` Clause in Go

The `range` clause is used with `for` loops to iterate over collections. Its
result depends on the value being ranged over:

- Arrays and slices produce an index and a copy of each element.
- Maps produce a key and its associated value.
- Strings produce a byte index and a decoded rune.
- Channels produce received values until the channel is closed.

## Ranging Over Slices

When ranging over a slice, the first value is the zero-based index and the
second value is a copy of the element at that index:

```go
numbers := []int{1, 2, 3}
for index, value := range numbers {
	fmt.Println(index, value)
}
```

Use the blank identifier `_` when the index is not needed:

```go
sum := 0
for _, value := range numbers {
	sum += value
}
```

Because `value` is a copy, changing it does not change the slice:

```go
for _, value := range numbers {
	value *= 2
}
fmt.Println(numbers) // [1 2 3]
```

To update the slice, use the index:

```go
for index := range numbers {
	numbers[index] *= 2
}
fmt.Println(numbers) // [2 4 6]
```

## Ranging Over Maps

Map iteration returns each key and value, but the order is unspecified and can
change between runs:

```go
marks := map[string]int{"Math": 90, "English": 85}
for subject, mark := range marks {
	fmt.Println(subject, mark)
}
```

Use only the key or only the value when the other one is unnecessary:

```go
for subject := range marks {
	fmt.Println(subject)
}

total := 0
for _, mark := range marks {
	total += mark
}
```

If output must be deterministic, collect and sort the keys before iterating:

```go
keys := make([]string, 0, len(marks))
for subject := range marks {
	keys = append(keys, subject)
}
sort.Strings(keys)
for _, subject := range keys {
	fmt.Println(subject, marks[subject])
}
```

## Ranging Over Strings

Ranging over a string decodes UTF-8. The first value is the byte offset of the
rune, not necessarily its character number, and the second value is a `rune`:

```go
text := "Go café"
for byteIndex, character := range text {
	fmt.Println(byteIndex, character, string(character))
}
```

The `é` character occupies two UTF-8 bytes, so the next byte index skips by
two. Invalid UTF-8 bytes are replaced with `utf8.RuneError`.

## Ranging Over Channels

A `for range` loop can receive values from a channel until the channel is
closed:

```go
values := make(chan int, 3)
values <- 1
values <- 2
values <- 3
close(values)

for value := range values {
	fmt.Println(value)
}
```

The sender, or another owner of the channel, should close it. Receiving from a
nil channel blocks forever, and ranging over an open channel waits for more
values.

## Range Variables and Mutation

The loop variables are assigned for each iteration. In modern Go, variables
declared by the loop are distinct per iteration when captured by a closure;
with older Go versions, take care when storing their addresses or capturing
them in goroutines. Regardless of version, the ranged element is a copy, so
mutate the original collection through its index when needed.

## Running the Example

Run [`range.go`](range.go) from the workspace root:

```bash
go run range/range.go
```

The example demonstrates traditional indexing, slice iteration, summing,
deterministic map iteration, map totals, and UTF-8 string iteration.
