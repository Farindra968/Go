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

type IMEPay struct{}

func (i IMEPay) ProviderName() string {
	return "IME Pay"
}

func (i IMEPay) Pay(amount float64) string {
	return fmt.Sprintf("Payment approved for %.2f via IME Pay.", amount)
}

func main() {
	gateways := []PaymentMethod{
		Esewa{},
		Khalti{},
		IMEPay{},
	}

	for _, gateway := range gateways {
		payment := Payment{gateway: gateway}
		payment.ProcessPayment(200.5)
	}
}
