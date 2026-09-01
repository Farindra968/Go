package auth

import "fmt"

func Authenticate(userName, password string) {
	fmt.Printf("Authenticating user: %s\n", userName)
}