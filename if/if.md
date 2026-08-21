
# If Statements in Go

Go uses `if` statements to run code only when a condition is true. The
condition must produce a boolean value (`true` or `false`), and parentheses
around the condition are not used.

## Basic Syntax

```go
if condition {
	// Runs when condition is true.
}
```

The opening brace must be on the same line as the condition. Braces are
required, even when the body contains only one statement.

## `else if` and `else`

Several alternatives can be connected in one conditional chain:

```go
if conditionA {
	// Runs when conditionA is true.
} else if conditionB {
	// Runs when conditionA is false and conditionB is true.
} else {
	// Runs when all previous conditions are false.
}
```

Go evaluates conditions from top to bottom and executes only the first
matching branch. The `else` branch is optional and runs when no earlier branch
matches.

## This Example

The program in [`if.go`](if.go) compares `age` with three ranges:

```go
age := 19

if age >= 18 {
	fmt.Println("You are an adult.")
} else if age >= 13 {
	fmt.Println("You are a teenager.")
} else {
	fmt.Println("You are a child.")
}
```

Because `19 >= 18` is true, the program prints:

```text
You are an adult.
```

The later conditions are not evaluated after the first matching branch. The
order matters: an age of `19` also satisfies `age >= 13`, but it is classified
as an adult because the adult condition comes first.

## Conditions

Conditions can use comparison and logical operators:

```go
if score >= 60 {
	fmt.Println("Passed")
}

if username != "" && password != "" {
	fmt.Println("Credentials provided")
}
```

Common comparison operators are `==`, `!=`, `<`, `<=`, `>`, and `>=`.
Logical operators are `&&` (and), `||` (or), and `!` (not). Go short-circuits
logical expressions: with `A && B`, `B` is skipped when `A` is false; with
`A || B`, `B` is skipped when `A` is true.

## Initialization Statement

An `if` statement may declare a value before its condition. That value is
available only inside the `if`, `else if`, and `else` branches:

```go
if length := len("hello"); length > 3 {
	fmt.Println("The string is longer than three characters.")
}
```

The initializer and condition are separated by a semicolon. This is useful
when a value is needed only for the decision.

## Important Rules

- The condition must be a `bool`; Go does not treat `0`, `1`, or empty strings
  as boolean values.
- `else` must appear on the same statement as the closing brace of the
  preceding branch: `} else {`.
- Conditions are not required to be surrounded by parentheses.
- Variables declared inside a branch are local to that branch.
- Use `switch` when many branches compare one value; use `if` for general
  boolean conditions or a small number of alternatives.

## Access-Control Example

The second conditional chain in [`if.go`](if.go) combines comparisons with
the `&&` operator:

```go
role := "user"
hasAccess := false

if role == "admin" && hasAccess {
	fmt.Println("Access granted, admin with permission.")
} else if role != "admin" {
	fmt.Println("Access denied, not an admin.")
} else {
	fmt.Println("Access denied, no permission.")
}
```

Access is granted only when both conditions are true. With the values shown,
the program prints `Access denied, not an admin.` The final `else` handles an
administrator who does not have permission.

## Run the Example

From the workspace root, run:

```bash
go run if/if.go
```

Expected output:

```text
You are an adult.
Access denied, not an admin.
```
