package main

import "fmt"

func main() {
	age := 19
	if age >= 18 {
		fmt.Println("You are an adult.")
	} else if age >= 13 {
		fmt.Println("You are a teenager.")
	} else {
		fmt.Println("You are a child.")
	}

	role := "user"
	hasAccess := false
	if role == "admin" && hasAccess {
		fmt.Println("Access granted, admin with permission.")
	} else if role != "admin" {
		fmt.Println("Access denied, not an admin.")
	} else {
		fmt.Println("Access denied, no permission.")
	}
}
