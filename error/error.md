# Error Handling in Go

Error handling in Go is explicit, simple, and predictable. Instead of exceptions, Go usually returns an `error` value from a function. This makes problems visible and forces the caller to handle them.

## Why Go uses errors

Go does not use `try/catch` like some other languages. Instead, functions commonly return `(value, error)`.

This approach makes the code honest:

- the function returns a result
- the caller checks whether it succeeded
- errors are handled at the call site

## Basic pattern

```go
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)
```

A common pattern is:

```go
value, err := someFunction()
if err != nil {
    // handle error
}
```

## Returning errors from functions

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

In this example:

- the function returns an integer and an `error`
- if the operation is valid, it returns the result and `nil`
- if something fails, it returns an error message

## Creating errors

### Using errors.New

```go
import "errors"

err := errors.New("invalid input")
```

This is the simplest method for creating an error with a fixed message.

### Using fmt.Errorf

```go
err := fmt.Errorf("invalid age: %d", age)
```

`fmt.Errorf` is useful when you want to include dynamic values in the error message.

## Wrapping errors

Go 1.13 introduced error wrapping with `%w`:

```go
err := fmt.Errorf("read config failed: %w", originalErr)
```

This allows the original error to be unwrapped later:

```go
fmt.Println(errors.Unwrap(err))
```

Wrapping is useful when you want to add context without losing the original cause.

## Validating input

```go
func validateAge(age int) error {
    if age < 18 {
        return fmt.Errorf("age %d is too young: must be at least 18", age)
    }
    return nil
}
```

Then you can check it like this:

```go
if err := validateAge(15); err != nil {
    fmt.Println("Validation error:", err)
}
```

## File errors

Many file operations return an error when something fails:

```go
file, err := os.Open("missing-file.txt")
if err != nil {
    fmt.Println("File open error:", err)
    return
}
```

This is a very common pattern in real Go programs because file access can fail because of permission issues, missing files, or invalid paths.

## Panic and recover

Go also has `panic` and `recover` for unexpected fatal situations. These are not the main error-handling mechanism in normal application code.

```go
func recoverFromPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered panic:", r)
        }
    }()

    panic("something went wrong")
}
```

Use `panic` only for truly exceptional cases. For regular application validation, prefer returning `error` values.

## Error handling best practices

- Check errors immediately after function calls.
- Return errors from functions instead of printing them and continuing.
- Add useful context with `fmt.Errorf`.
- Prefer `nil` for success and an error value for failure.
- Use wrapping when the original cause matters.
- Avoid swallowing errors silently.

## Example program

The sample program is available in [error/error.go](error/error.go). It demonstrates:

- `errors.New`
- returning `(value, error)`
- `fmt.Errorf`
- error wrapping with `%w`
- input validation
- file-related error handling
- panic recovery

## Summary

Error handling in Go is based on explicit return values. By checking `err != nil`, you keep programs predictable and easier to debug. This pattern makes Go code safe, readable, and robust.

A typical Go function follows this rule:

```go
result, err := doWork()
if err != nil {
    return err
}
```

This simple model is one of the reasons Go code is known for being clear and maintainable.
