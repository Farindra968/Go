package main

import "fmt"

// OrderStatus is an enum-like type using int as the underlying type.
// The values are assigned automatically by iota.
type OrderStatus int

const (
	Received   OrderStatus = iota // 0
	Processing                    // 1
	Shipped                       // 2
	Delivered                     // 3
	Canceled                      // 4
)

// PaymentStatus is an enum-like type using string as the underlying type.
// String-based enums are commonly used for data that is read and stored as text.
type PaymentStatus string

const (
	PaymentPending   PaymentStatus = "Pending"
	PaymentCompleted PaymentStatus = "Completed"
	PaymentFailed    PaymentStatus = "Failed"
)

// changeOrderStatus prints the current order state.
func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

// changePaymentStatus prints the current payment state.
func changePaymentStatus(status PaymentStatus) {
	fmt.Println("Payment status changed to:", status)
}

// String returns a readable name for an OrderStatus value.
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

// String returns a readable name for a PaymentStatus value.
func (s PaymentStatus) String() string {
	return string(s)
}

// printOrderProgress shows an additional example of using enum values.
func printOrderProgress(status OrderStatus) {
	fmt.Printf("Current order state: %s\n", status.String())
}

func main() {
	// Existing examples: keep these as the main demonstration.
	changeOrderStatus(Received)
	changeOrderStatus(Processing)
	changePaymentStatus(PaymentPending)
	changePaymentStatus(PaymentCompleted)
	changePaymentStatus(PaymentFailed)

	// Additional examples for better understanding.
	fmt.Println("--- Extra enum examples ---")
	printOrderProgress(Shipped)
	fmt.Println("Current payment state:", PaymentPending.String())
	fmt.Println("Current payment state:", PaymentCompleted.String())
}
