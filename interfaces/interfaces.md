# Interfaces in Go

An interface in Go is a type that defines a set of methods. Any type that implements those methods automatically satisfies the interface. This is one of the key ideas behind Go's approach to polymorphism.

A type does not need to declare that it implements an interface explicitly. If it has the required methods, it is considered compatible.

## Why Interfaces Are Useful

Interfaces help you write flexible code. Instead of depending on a specific concrete type, your code can work with any type that satisfies the same behavior.

This makes it easier to:

- write reusable functions
- swap implementations without changing code
- build loosely coupled systems
- work with multiple providers or services in the same way

## Interface Syntax

```go
type PaymentMethod interface {
    Pay(amount float64) string
    ProviderName() string
}
```

This interface requires any implementation to provide two methods:

- `Pay(amount float64) string`
- `ProviderName() string`

## Example: Payment Gateways

```go
package main

import "fmt"

type PaymentMethod interface {
    Pay(amount float64) string
    ProviderName() string
}

type Payment struct {
    gateway PaymentMethod
}

func (p Payment) ProcessPayment(amount float64) {
    result := p.gateway.Pay(amount)
    fmt.Printf("Using %s: %s\n", p.gateway.ProviderName(), result)
}

type Esewa struct{}

func (e Esewa) ProviderName() string {
    return "Esewa"
}

func (e Esewa) Pay(amount float64) string {
    return fmt.Sprintf("Payment approved for %.2f via Esewa.", amount)
}

type Khalti struct{}

func (k Khalti) ProviderName() string {
    return "Khalti"
}

func (k Khalti) Pay(amount float64) string {
    return fmt.Sprintf("Payment approved for %.2f via Khalti.", amount)
}

func main() {
    gateways := []PaymentMethod{
        Esewa{},
        Khalti{},
    }

    for _, gateway := range gateways {
        payment := Payment{gateway: gateway}
        payment.ProcessPayment(200.5)
    }
}
```

## How It Works

The `Payment` struct has a field named `gateway` of type `PaymentMethod`.

That means the `gateway` can be any type that implements `PaymentMethod`, such as:

- `Esewa`
- `Khalti`
- `IMEPay`

The `ProcessPayment` method calls:

```go
result := p.gateway.Pay(amount)
fmt.Printf("Using %s: %s\n", p.gateway.ProviderName(), result)
```

This is called polymorphism: the same method can work with different concrete types as long as they share the same interface.

## Interface Implementation Rules

A type satisfies an interface when it has all of that interface's methods.

For example, both `Esewa` and `Khalti` satisfy `PaymentMethod` because they provide:

```go
Pay(amount float64) string
ProviderName() string
```

If one of the methods is missing, the type does not implement the interface.

## Empty Interface

The empty interface is written as:

```go
interface{}
```

It can hold any type because it has no methods. This is useful for generic or flexible code, but it should be used carefully.

```go
var value any = "Go"
var number any = 42
```

In modern Go, `any` is the preferred name for the empty interface.

## Interface and Method Sets

An interface defines a method set. A value of a type implements the interface if it has those methods.

### Example

```go
type Speaker interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof"
}
```

Now `Dog` implements `Speaker`.

## Benefits of Interfaces

Interfaces let you design code around behavior rather than exact types. This is especially useful in:

- payment systems
- database access layers
- logging systems
- API integrations
- testing and mocking

## Important Points

- An interface declares behavior, not data.
- Go uses structural typing, not inheritance.
- A type can implement multiple interfaces.
- Interfaces make code more flexible and easier to extend.
- A value can be stored in an interface variable only if its dynamic type implements that interface.

## Summary

Interfaces are one of Go's most important features. They let different types share the same behavior while keeping the code flexible and clean. In this example, `Esewa`, `Khalti`, and `IMEPay` all implement the same payment behavior, so they can be used through the same `PaymentMethod` interface.
