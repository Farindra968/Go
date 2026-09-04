# Go Learning Roadmap

## Introduction to Go

Go, also known as Golang, is a statically typed, compiled programming language designed for simplicity, speed, and reliability. It is widely used for backend services, CLI tools, APIs, cloud applications, and systems programming.

This repository is a hands-on Go learning workspace with small practical examples and notes organized by topic. Each section focuses on one concept so you can learn step by step and practice by running the code.

## Why learn Go?

- Simple and readable syntax
- Fast compilation and execution
- Strong standard library
- Great for backend development and APIs
- Excellent support for concurrency and networking
- Used by large companies and modern cloud systems

## Repository structure

This project is organized by topic. Each folder contains a Go example file and a matching markdown explanation.

## Learning order

1. [Basic Go](basic/go.md)
2. [Variables](variable/variable.md)
3. [Types](types/types.md)
4. [Constants](constants/constants.md)
5. [Arrays](arrays/arrays.md)
6. [Slices](slices/slices.md)
7. [Maps](maps/map.md)
8. [Range](range/range.md)
9. [Functions](functions/functions.md)
10. [Variadic Functions](variadic_functions/variadic_functions.md)
11. [Closures](closures/closures.md)
12. [Pointers](pointers/pointers.md)
13. [Structs](structs/structs.md)
14. [Interfaces](interfaces/interfaces.md)
15. [If Statements](if/if.md)
16. [Switch Statements](switch/switch.md)
17. [Loops](loop/loop.md)
18. [Generics](generics/generics.md)
19. [Enums](enums/enums.md)
20. [Files](files/files.md)
21. [Error Handling](error/error.md)
22. [Packages](packages/packages.md)

## Learning path

This order follows a beginner-friendly progression:

- start with syntax and variables
- move into types and data structures
- learn control flow and functions
- explore reusable abstractions like structs, interfaces, and generics
- understand file handling and error management
- finally study modular code with packages

## Project 1: Student API

The `project1` folder contains a small REST API built with Go's standard `net/http` package. It demonstrates:

- YAML configuration loading with environment support
- SQLite database storage
- Student creation and lookup
- Request validation with `validator`
- JSON responses and graceful server shutdown

### Run Project 1

```bash
cd project1
go run ./cmd/project1
```

The server starts at `http://localhost:8080`. The SQLite database is created at `storage/storage.db`.

### Project 1 endpoints

Check the server health:

```bash
curl http://localhost:8080/api/v1/health
```

Create a student:

```bash
curl -X POST http://localhost:8080/api/v1/students \
	-H "Content-Type: application/json" \
	-d '{"name":"Ada Lovelace","email":"ada@example.com","age":28,"password":"secret123"}'
```

Fetch a student by ID:

```bash
curl http://localhost:8080/api/v1/students/{id}
```

## How to use this repo

Open any topic folder and run the corresponding Go file.

Example:

```bash
go run generics/generics.go
```

Or run from a specific topic folder:

```bash
cd files
go run files.go
```

## Topics covered

This repo includes practice on:

- basic Go syntax
- variables and constants
- primitive and custom types
- arrays, slices, and maps
- loops, conditions, and switch cases
- functions and variadic functions
- closures and pointers
- structs and interfaces
- generics and enums
- file operations and working with data
- error handling and validation
- packages and modular code structure

## Current focus

The latest topic in the collection is error handling, which covers returning errors, wrapping errors, validating input, and handling panics safely.

## Notes

Each topic is designed to be simple, practical, and easy to follow. The goal is not just to read theory, but to write code, run it, and understand how Go works in real examples.

## Summary

This project is a compact Go learning repository for beginners and practicing developers. It combines small code examples with documentation so you can learn the language by doing, not just by reading.
