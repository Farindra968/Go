# Switch Statements in Go

Go's `switch` statement selects one branch from a group of alternatives. It
is useful when a value can match several cases and is often clearer than a
long `if` and `else if` chain.

## Basic Syntax

```go
switch expression {
case value1:
	// Runs when expression == value1.
case value2:
	// Runs when expression == value2.
default:
	// Runs when no case matches.
}
```

Go evaluates cases from top to bottom and runs the first matching case. A
matching case does not continue into the next case automatically, so an
explicit `break` is normally unnecessary. The `default` case is optional.

## Expression Switch

The first switch in [`switch.go`](switch.go) compares the `grade` string with
the case values:

```go
grade := "B"
switch grade {
case "A":
	fmt.Println("Excellent!")
case "B":
	fmt.Println("Good!")
case "C", "D":
	fmt.Println("Average.")
default:
	fmt.Println("Invalid grade.")
}
```

Multiple values can be listed in one case, separated by commas. Each value is
compared with the switch expression using equality. The example prints
`Good!` because `grade` is `"B"`.

## Multiple Case Values

The `day` example groups Saturday and Sunday into one branch:

```go
switch day {
case "Saturday", "Sunday":
	fmt.Println("It is the weekend.")
default:
	fmt.Println("It is a weekday.")
}
```

This example prints `It is the weekend.`

## Condition-Only Switch

Omitting the switch expression creates a condition-only switch. Each case must
be a boolean expression, and the first true case runs:

```go
switch score := 87; {
case score >= 90:
	fmt.Println("Score: excellent")
case score >= 60:
	fmt.Println("Score: passing")
default:
	fmt.Println("Score: failing")
}
```

This is the switch equivalent of an `if`, `else if`, and `else` chain. An
initialization statement may appear before the semicolon; its variables are
scoped to the switch.

## Multiple Conditions

A condition-only switch can contain several boolean conditions. The first
condition that evaluates to `true` is selected:

```go
temperature := 25
switch {
case temperature < 0:
	fmt.Println("It's freezing!")
case temperature >= 0 && temperature <= 20:
	fmt.Println("It's cold.")
case temperature > 20 && temperature <= 30:
	fmt.Println("It's warm.")
default:
	fmt.Println("It's hot!")
}
```

The `&&` operator requires both comparisons to be true. With a temperature of
`25`, the third case matches and the program prints `It's warm.` The ranges
are ordered and non-overlapping, so exactly one branch handles each normal
temperature value.

## `fallthrough`

Go does not fall through to the next case by default. The `fallthrough`
statement explicitly transfers control to the next case without checking its
condition again:

```go
switch value {
case 1:
	fmt.Println("first case")
	fallthrough
case 2:
	fmt.Println("second case")
}
```

When `value` is `1`, both messages are printed. `fallthrough` must be the final
non-empty statement in its case and cannot be used in the final case. Use it
sparingly because it can make control flow less obvious.

## Type Switch

A type switch selects a case based on the dynamic type stored in an interface:

```go
func printType(value any) {
	switch value := value.(type) {
	case string:
		fmt.Printf("Type: string (%q)\n", value)
	case int:
		fmt.Printf("Type: int (%d)\n", value)
	case bool:
		fmt.Printf("Type: bool (%t)\n", value)
	default:
		fmt.Printf("Type: %T\n", value)
	}
}
```

Inside each case, `value` has the corresponding case type. The `default`
case handles types that are not listed. A type switch requires an interface
value; `any` is an alias for `interface{}`.

## Important Rules

- A case value must be comparable with the switch expression.
- Cases are checked in source order, and only the first matching case runs.
- Cases do not need to be constants; they may contain expressions when the
  switch form allows them.
- A switch can have no expression, which is equivalent to switching on `true`.
- `break` exits a switch early but is usually unnecessary because cases do not
  fall through.
- `fallthrough` is not allowed in a type switch.
- An empty switch has no cases to execute and is valid Go.

## Switch Compared with `if`

Use `switch` when several alternatives belong to one decision, especially
when comparing one value with many possible cases. Use `if` when conditions
are unrelated, involve complex boolean logic, or need a simple two-way test.

## Run the Example

From the workspace root:

```bash
go run switch/switch.go
```

Expected output:

```text
Good!
It is the weekend.
Score: passing
It's warm.
Fallthrough: first case
Fallthrough: second case
Type: string ("hello")
Type: int (42)
Type: bool (true)
```
