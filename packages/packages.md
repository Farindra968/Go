# Go Packages

A package in Go is a way to organize related code into a reusable unit. Packages help keep programs structured, easier to maintain, and easier to share across files or projects.

## What is a package?

In Go, every file belongs to a package. A package can be:

- `main` — an executable program
- a custom package — reusable library code
- an external dependency — a package imported from another module or repository

## Package structure in this project

This folder contains three examples:

- `main` package in `main.go`
- `auth` package in `auth/auth.go` and `auth/token.go`
- `user` package in `user/user.go`

## Example: package auth

```go
package auth

import "fmt"

func Authenticate(userName, password string) {
    fmt.Printf("Authenticating user: %s\n", userName)
}
```

This file belongs to the `auth` package. It contains a reusable function that can be imported by another Go program.

## Example: package auth token helper

```go
package auth

func TokenSession() string {
    return "fv9sjm,fjerfj98y34r834wfi34i9u98I8I3UTI39JEU8"
}
```

This package provides a helper function that returns a token string.

## Example: package user

```go
package user

// User represents a user in the system
type User struct {
    ID        int
    Username  string
    Password  string
    Email     string
    FirstName string
    LastName  string
}
```

The `user` package defines a reusable data structure called `User`.

## Example: importing packages in main

```go
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
```

### What this does

- imports the `auth` package
- imports the `user` package
- imports an external package called `fatih/color`
- calls functions and uses the `User` struct from different packages

## Why packages are useful

Packages help us:

- organize code logically
- reuse functions across files
- keep code clean and readable
- split large programs into smaller units
- work with external libraries easily

## Package naming rules

- package names are usually lowercase
- `main` is reserved for executable programs
- package names should describe the purpose of the code
- exported names (capitalized) can be accessed from other packages

## Exported vs unexported names

If a function or variable starts with an uppercase letter, it can be used outside the package.

```go
func Authenticate(userName, password string) {
    // exported
}

func tokenSession() string {
    // unexported
}
```

`Authenticate` is exported because it begins with an uppercase letter, so it can be accessed from another package.

## Running the example

From the project root, run:

```bash
cd /c/Coding/Go/packages
go run main.go
```

This will execute the package example and print the authentication message and token information.

## Summary

Packages are one of the most important concepts in Go. They let you build modular and reusable programs. In this project, the `auth` and `user` packages demonstrate how Go code is organized into separate reusable units and imported into the main application.
