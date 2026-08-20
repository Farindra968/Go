package main

import "fmt"

const (
	appName       = "Go Constants"
	version       = 1
	secondsInHour = 60 * 60
)

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

func main() {
	const pi float64 = 3.14
	const greeting = "Hello, Go!"
	const radius = 5
	const area = pi * radius * radius

	fmt.Println("Application:", appName)
	fmt.Println("Version:", version)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Value of pi:", pi)
	fmt.Println("Circle area:", area)
	fmt.Println("Seconds in an hour:", secondsInHour)
	fmt.Println("Days:", sunday, monday, tuesday, wednesday, thursday, friday, saturday)
}
