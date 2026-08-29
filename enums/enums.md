# Enums in Go

Go does not have a built-in `enum` keyword like some other languages such as Java or C#. Instead, Go developers usually model enums using named constants.

There are two common patterns:

- integer-based enums using `iota`
- string-based enums using string constants

These are often called "enum-like" values in Go.

## Why Use Enums?

Enums help make code more readable and safer by giving meaningful names to fixed values instead of raw numbers or strings.

Examples:

- order status
- payment status
- user role
- error types
- category names

## 1. Integer-Based Enum with iota

`iota` is a Go feature that automatically increments values in a constant block.

```go
package main

import "fmt"

type OrderStatus int

const (
    Received  OrderStatus = iota
    Processing
    Shipped
    Delivered
    Canceled
)

func main() {
    fmt.Println(Received)   // 0
    fmt.Println(Processing) // 1
    fmt.Println(Shipped)    // 2
}
```

In this example:

- `Received` gets `0`
- `Processing` gets `1`
- `Shipped` gets `2`
- `Delivered` gets `3`
- `Canceled` gets `4`

This is useful when the values are stored internally as numbers but you want readable names in code.

## 2. String-Based Enum

If you want the value to be stored as text, use a string-based type.

```go
type PaymentStatus string

const (
    PaymentPending  PaymentStatus = "Pending"
    PaymentCompleted PaymentStatus = "Completed"
    PaymentFailed   PaymentStatus = "Failed"
)
```

This is often better for APIs, database values, or JSON responses because the value is human-readable.

## Example: Order and Payment Status Together

```go
package main

import "fmt"

type OrderStatus int

type PaymentStatus string

const (
    Received OrderStatus = iota
    Processing
    Shipped
    Delivered
    Canceled
)

const (
    PaymentPending  PaymentStatus = "Pending"
    PaymentCompleted PaymentStatus = "Completed"
    PaymentFailed   PaymentStatus = "Failed"
)

func changeOrderStatus(status OrderStatus) {
    fmt.Println("Order status changed to:", status)
}

func changePaymentStatus(status PaymentStatus) {
    fmt.Println("Payment status changed to:", status)
}

func main() {
    changeOrderStatus(Received)
    changeOrderStatus(Processing)

    changePaymentStatus(PaymentPending)
    changePaymentStatus(PaymentCompleted)
    changePaymentStatus(PaymentFailed)
}
```

## Custom String Method

Go allows you to define a `String()` method on your enum type to make output more readable.

```go
func (s OrderStatus) String() string {
    switch s {
    case Received:
        return "Received"
    case Processing:
        return "Processing"
    case Shipped:
        return "Shipped"
    case Delivered:
        return "Delivered"
    case Canceled:
        return "Canceled"
    default:
        return "Unknown"
    }
}
```

Now printing the value gives a clearer result instead of only a number.

```go
fmt.Println(Shipped.String())
// Output: Shipped
```

## Using Enums in Real Applications

Enums are very common in real systems.

Examples:

- order lifecycle: `Received`, `Processing`, `Shipped`, `Delivered`
- payment lifecycle: `Pending`, `Completed`, `Failed`
- user roles: `Admin`, `User`, `Guest`
- product categories: `Electronics`, `Clothing`, `Food`

## Benefits of Enums

- improves readability
- reduces magic numbers and raw strings
- makes code easier to maintain
- reduces mistakes caused by invalid values
- helps create clearer logic in applications

## When to Use Enums

Use enums when a variable should only take a fixed set of valid values.

For example:

```go
if status == Delivered {
    fmt.Println("Order has been delivered.")
}
```

This is much clearer than comparing with raw numbers like `3`.

## Important Notes

- Go does not have a built-in enum type.
- `iota` is useful for sequential integer states.
- String-based enums are easier to read in logs and APIs.
- You can define helper methods such as `String()` for better output.

## Summary

Enums in Go are usually implemented as named constants. They make the code expressive and keep values within a controlled set.

In this example:

- `OrderStatus` is an integer-based enum using `iota`
- `PaymentStatus` is a string-based enum
- both are used to track different states clearly and consistently

This pattern is widely used in Go applications for state management, workflows, and status tracking.
