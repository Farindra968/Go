# Go Mutex

A mutex (short for mutual exclusion) is a synchronization tool in Go that helps protect shared data when multiple goroutines run at the same time.

If two or more goroutines try to change the same variable at the same time, they may interfere with each other. This is called a race condition. A mutex helps prevent that by allowing only one goroutine to access the shared resource at a time.

## Why do we need a mutex?

When goroutines share data, this can happen:

- one goroutine reads a value while another writes it
- two goroutines write the same value at once
- the final result becomes incorrect or unpredictable

A mutex solves this by locking shared data during modification.

## Basic idea

```go
var mu sync.Mutex

mu.Lock()
// critical section
sharedValue++
mu.Unlock()
```

- `Lock()` stops other goroutines from entering the protected section.
- `Unlock()` allows the next goroutine to continue.

## Important rule

Always lock around the code that reads or writes shared data.

```go
mu.Lock()
value = value + 1
mu.Unlock()
```

If you forget to lock, your program may behave incorrectly.

## Example: safe increment

```go
package main

import (
    "fmt"
    "sync"
)

type Counter struct {
    value int
    mu    sync.Mutex
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func main() {
    var wg sync.WaitGroup
    counter := Counter{}

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            counter.Increment()
        }()
    }

    wg.Wait()
    fmt.Println("Final value:", counter.value)
}
```

This example is safe because all writes to `value` happen inside the mutex lock.

## When to use mutex

Use a mutex when:

- multiple goroutines share one variable
- one or more goroutines update the same data
- we need consistent final results

Common examples:

- counters
- caches
- maps accessed by multiple goroutines
- web app request counters

## When not to use mutex

Do not use a mutex for every small task. It is not needed when:

- data is local to one goroutine
- values are read-only and never shared
- channels are used for communication instead of shared memory

In Go, channels are often a better way to communicate between goroutines than sharing variables directly.

## Mutex vs channel

A mutex protects shared memory.
A channel passes data between goroutines.

Use:

- `mutex` for shared state protection
- `channel` for communication and coordination

## Example: reading safely

```go
func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}
```

This keeps reads and writes consistent.

## Practical example

```go
type Post struct {
    view int
    mu   sync.Mutex
}

func (p *Post) Inc() {
    p.mu.Lock()
    defer p.mu.Unlock()
    p.view++
}
```

If 100 goroutines call `Inc()` at the same time, the final value will still be correct because the mutex ensures only one increment happens at a time.

## Summary

- A mutex controls access to shared data.
- It prevents race conditions.
- Use `Lock()` before reading or writing shared values.
- Use `Unlock()` right after the critical section.
- Always keep the protected code as small as possible.

Mutexes are an essential part of Go concurrency, especially when multiple goroutines modify the same data.
