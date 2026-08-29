# Goroutines in Go

A goroutine is a lightweight thread managed by the Go runtime. It allows a function to run concurrently with other functions without creating a heavy operating system thread for each task.

Goroutines are one of the main features that make Go great for concurrent programming.

## Why use goroutines?

Goroutines are useful when you want to:

- run tasks in parallel
- handle multiple requests or jobs at the same time
- keep the main program responsive
- improve performance for independent work

## Basic syntax

Use the `go` keyword before a function call to start a goroutine.

```go
go doWork()
```

This means `doWork()` will run concurrently while the rest of the program continues.

## Example: simple goroutine

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello()

    fmt.Println("Hello from main")
    time.Sleep(500 * time.Millisecond)
}
```

### Output

```text
Hello from main
Hello from goroutine
```

The order may vary because goroutines run concurrently.

## Why sleep is used

In Go, the main function exits when the program ends. If a goroutine is still running, it may be stopped with the program. That is why examples often use `time.Sleep()` to keep the program alive long enough to allow goroutines to finish.

However, in real programs, better synchronization tools like `WaitGroup` are preferred.

## Using WaitGroup

A `sync.WaitGroup` makes it easier to wait for multiple goroutines to finish.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func task(id int) {
    fmt.Printf("Task %d started\n", id)
    time.Sleep(500 * time.Millisecond)
    fmt.Printf("Task %d finished\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            task(id)
        }(i)
    }

    wg.Wait()
    fmt.Println("All tasks completed")
}
```

`wg.Add(1)` tells the wait group to expect one goroutine.

`wg.Done()` is called when the goroutine is finished.

`wg.Wait()` blocks until all goroutines complete.

## Important notes

- Goroutines are lightweight compared to OS threads.
- They are scheduled by the Go runtime.
- They run concurrently, not necessarily in order.
- Shared data must be handled carefully to avoid data races.
- Use synchronization primitives like channels, mutexes, and `WaitGroup` for safe coordination.

## Data race example

When multiple goroutines access the same variable at the same time, the result can be unpredictable.

```go
package main

import (
    "fmt"
)

func main() {
    counter := 0

    for i := 0; i < 1000; i++ {
        go func() {
            counter++
        }()
    }

    fmt.Println(counter)
}
```

This is not safe without synchronization because multiple goroutines may update `counter` at the same time.

## Summary

Goroutines let you run functions concurrently in Go. They are ideal for background tasks, parallel processing, and responsive programs.

The key points are:

- use `go` to start a goroutine
- use `WaitGroup` to wait for completion
- be careful with shared variables
- prefer synchronization when multiple goroutines access the same data

## Example file in this folder

The sample in this folder demonstrates launching multiple goroutines and waiting for them to finish using `sync.WaitGroup`.
