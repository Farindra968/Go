package main

import (
	"github.com/Farindra968/mygo/auth"
	"github.com/Farindra968/mygo/user"
	"github.com/fatih/color"
)

func main() {
	auth.Authenticate("hari", "Admin1234@")
	token := auth.TokenSession()

	Data := user.User{
		ID:        1,
		Username:  "hari",
		Password:  "Admin1234@",
		Email:     "hari@example.com",
		FirstName: "Hari",
		LastName:  "Kumar",
	}

d := color.New(color.FgCyan, color.Bold)
	d.Println("Token:", token, Data)
} 