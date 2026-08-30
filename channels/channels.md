# Go Channels

A channel in Go is a communication mechanism used to send and receive values between goroutines. It helps goroutines share data safely and coordinate work.

## Why use channels?

Channels are useful when:

- one goroutine needs to send data to another
- tasks need to be synchronized
- a goroutine should report a result back to the main goroutine
- multiple goroutines communicate with each other

Without channels, goroutines may run independently and data races can happen if shared variables are modified without synchronization.

## Creating a channel

```go
ch := make(chan int)
```

This creates an unbuffered channel of type int.

## Sending and receiving values

```go
ch := make(chan int)

go func() {
    ch <- 42
}()

value := <-ch
fmt.Println(value) // 42
```

- `ch <- 42` sends data into the channel.
- `<-ch` receives data from the channel.

If the channel is unbuffered, the send and receive must happen at the same time.

## Unbuffered channel

An unbuffered channel blocks until a sender and receiver are both ready.

```go
package main

import "fmt"

func main() {
    ch := make(chan int)

    go func() {
        fmt.Println("Sending value...")
        ch <- 10
    }()

    fmt.Println("Receiving value...", <-ch)
}
```

This pattern is useful when you want strong synchronization between goroutines.

## Buffered channel

A buffered channel has a capacity. It can store values without blocking until the buffer is full.

```go
package main

import "fmt"

func main() {
    ch := make(chan string, 2)

    ch <- "Hello"
    ch <- "World"

    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

Here, the channel can hold up to 2 values before the sender blocks.

## Closing a channel

When a sender is done, it can close the channel.

```go
close(ch)
```

Closing a channel tells receivers that no more values will be sent.

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)

    for value := range ch {
        fmt.Println(value)
    }
}
```

The `range` loop automatically stops when the channel is closed.

## Range over channel

This pattern is common for reading values until the channel is closed.

```go
for value := range ch {
    fmt.Println(value)
}
```

A `range` loop keeps receiving until the channel is closed.

## Select statement

The `select` statement lets a goroutine wait on multiple channel operations.

```go
package main

import "fmt"

func main() {
    ch1 := make(chan string, 1)
    ch2 := make(chan string, 1)

    ch1 <- "one"
    ch2 <- "two"

    select {
    case msg1 := <-ch1:
        fmt.Println("Received from ch1:", msg1)
    case msg2 := <-ch2:
        fmt.Println("Received from ch2:", msg2)
    default:
        fmt.Println("No channel ready")
    }
}
```

`select` is often used for timeouts, multiple goroutines, and non-blocking communication.

## Directional channels

Channels can be restricted to only send or only receive.

```go
func sendData(ch chan<- string) {
    ch <- "Hello"
}

func readData(ch <-chan string) {
    fmt.Println(<-ch)
}
```

- `chan<-` means send-only channel
- `<-chan` means receive-only channel

This improves code clarity and safety.

## Example: passing results from a goroutine

```go
package main

import "fmt"

func add(a, b int, result chan int) {
    result <- a + b
}

func main() {
    result := make(chan int)

    go add(10, 20, result)

    sum := <-result
    fmt.Println("Sum is:", sum)
}
```

This is one of the most common uses of channels in Go.

## Important points

- Channels are used for communication between goroutines.
- Unbuffered channels block until both sender and receiver are ready.
- Buffered channels allow short-term storage of values.
- Close channels when no more values will be sent.
- Use `range` to read until closure.
- Use `select` to handle multiple channel cases.

## Summary

Channels are one of the most important concepts in Go concurrency. They allow goroutines to communicate safely and coordinate tasks. Learning how to use them correctly is essential for writing concurrent Go programs.

## Practice example

```go
package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
    for job := range jobs {
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 3)
    results := make(chan int, 3)

    jobs <- 1
    jobs <- 2
    jobs <- 3
    close(jobs)

    go worker(jobs, results)

    for result := range results {
        fmt.Println(result)
    }
}
```

This example shows how workers can receive jobs and send results through channels.
