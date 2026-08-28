# Structs in Go

A **struct** is a user-defined type that groups related values into named
fields. Structs are useful for representing real-world data such as people,
orders, products, and employees.

```go
type Person struct {
	Name      string
	Age       int
	Address   string
	Amount    float64
	IsMarried bool
	Role      string
}
```

Each field has a name and a type. A struct value contains one value for every
field, and fields are accessed with the dot operator (`.`).

## Creating a Struct

The recommended form is a **keyed struct literal**. It explicitly identifies
each field, so the code remains readable if the struct definition changes.

```go
person := Person{
	Name:      "John Doe",
	Age:       30,
	Address:   "123 Main St",
	Amount:    100.50,
	IsMarried: true,
	Role:      "User",
}

fmt.Println(person.Name)
fmt.Println(person.Age)
```

Fields can be updated directly:

```go
person.Role = "Admin"
person.IsMarried = false
```

An unkeyed literal is also possible, but it depends on field order and must
provide values in exactly that order. Keyed literals are usually safer:

```go
person := Person{"John Doe", 30, "123 Main St", 100.50, true, "User"}
```

## Printing Structs

`fmt.Printf` with `%+v` prints field names as well as values:

```go
fmt.Printf("%+v\n", person)
```

The ordinary `%v` verb prints values without field names. `%#v` prints a Go
syntax representation that includes the type.

## Zero Values

Declaring a struct without an initializer gives every field its zero value:

```go
var emptyPerson Person
fmt.Printf("%+v\n", emptyPerson)
```

The zero values in this example are:

| Field | Zero value |
|---|---|
| `Name`, `Address`, `Role` | `""` |
| `Age` | `0` |
| `Amount` | `0` |
| `IsMarried` | `false` |

Go's useful zero values mean that a struct can often be used immediately after
declaration without a special initialization function.

## Methods on Structs

A method is a function with a receiver. The receiver connects the method to a
type.

### Pointer Receiver

`ChangeMaritalStatus` and `ChangeRole` use `*Person` receivers:

```go
func (p *Person) ChangeRole(role string) {
	p.Role = role
}
```

A pointer receiver can change the original struct. Go automatically takes the
address of an addressable value when calling the method:

```go
person.ChangeRole("Admin")
```

This is equivalent to calling the method with a pointer:

```go
(&person).ChangeRole("Admin")
```

Use a pointer receiver when the method must mutate the value or when copying a
large struct should be avoided.

### Value Receiver

`Summary` uses a value receiver because it only reads the struct:

```go
func (p Person) Summary() string {
	return fmt.Sprintf("%s (%d), role: %s", p.Name, p.Age, p.Role)
}
```

A value receiver receives a copy of the struct. Changes made to that copy do
not affect the original value.

## Constructor-Style Functions

Go does not have a built-in constructor keyword. A common Go pattern is a
function named `NewType` that initializes and returns a value or pointer:

```go
func NewPerson(
	name string,
	age int,
	address string,
	amount float64,
	isMarried bool,
	role string,
) *Person {
	return &Person{
		Name:      name,
		Age:       age,
		Address:   address,
		Amount:    amount,
		IsMarried: isMarried,
		Role:      role,
	}
}

person := NewPerson("Jane Smith", 28, "456 Elm St", 200.75, false, "Manager")
```

Returning `*Person` lets callers use pointer-receiver methods directly and
avoids copying the struct.

## Anonymous Structs

An anonymous struct is declared without a named type. It is useful for
short-lived data used in one function or for small temporary values:

```go
language := struct {
	Name       string
	IsCompiled bool
}{
	Name:       "Go",
	IsCompiled: true,
}

fmt.Printf("%+v\n", language)
```

Two anonymous struct types are identical only when their fields have the same
names, types, order, and tags. For data used in multiple places, define a
named type instead.

## Struct Embedding

Embedding places one named struct inside another without giving the embedded
field a separate name at the outer level:

```go
type Employee struct {
	Person
	EmployeeID string
	Department string
}
```

An `Employee` can initialize the embedded struct explicitly:

```go
employee := Employee{
	Person: Person{
		Name: "Jack",
		Age:  35,
		Role: "Developer",
	},
	EmployeeID: "E123",
	Department: "IT",
}
```

The embedded fields and methods are **promoted**, so these selectors are valid:

```go
fmt.Println(employee.Name)
employee.ChangeRole("Senior Developer")
fmt.Println(employee.Summary())
```

The explicit form is always available too:

```go
fmt.Println(employee.Person.Name)
employee.Person.ChangeRole("Developer")
```

Embedding is composition, not inheritance. `Employee` contains a `Person`, but
it is still a distinct type and can have its own fields and methods.

## Struct Pointers

A pointer to a struct stores the address of the struct:

```go
personPointer := &person
personPointer.Role = "Admin"
```

Go automatically dereferences a struct pointer for field access, so
`personPointer.Role` is equivalent to `(*personPointer).Role`.

## Exported and Unexported Fields

An identifier beginning with an uppercase letter is exported and can be used
from another package. An identifier beginning with a lowercase letter is
unexported and is accessible only inside its own package.

The fields in `Person` are exported because they begin with uppercase letters.
This makes them accessible to code importing the package. Keep fields
unexported when a type must control how its data is read or changed, and expose
methods for that controlled behavior.

## Comparing Structs

Structs are comparable with `==` when all their fields are comparable. Strings,
numbers, booleans, arrays, and structs are comparable. Slices, maps, and
functions are not comparable, so a struct containing one of those types cannot
be compared with `==`.

```go
first := Person{Name: "Ada", Age: 30}
second := Person{Name: "Ada", Age: 30}
fmt.Println(first == second) // true
```

## Structs with Slices and Maps

Struct fields can have any valid Go type, including slices and maps:

```go
type Team struct {
	Name    string
	Members []string
	Scores  map[string]int
}
```

The struct itself can still be copied, but the slice and map fields refer to
shared underlying data. Use a deep copy when independent nested data is
required.

## Complete Example

The runnable example in [`structs.go`](structs.go) demonstrates all of the
topics above:

```bash
go run ./structs/structs.go
```

Expected output:

```text
Person: John Doe (30), role: User
Full value: {Name:John Doe Age:30 Address:123 Main St Amount:100.5 IsMarried:true Role:User}
Updated person: John Doe (30), role: Admin
Updated value: {Name:John Doe Age:30 Address:123 Main St Amount:100.5 IsMarried:false Role:Admin}
Zero value: {Name: Age:0 Address: Amount:0 IsMarried:false Role:}
Constructed person: Jane Smith (28), role: Manager
Anonymous struct: {Name:Go IsCompiled:true}
Employee: {Person:{Name:Jack Age:35 Address:789 Oak St Amount:300 IsMarried:true Role:Developer} EmployeeID:E123 Department:IT}
Promoted method: Jack (35), role: Senior Developer
Employee name: Jack
Employee age: 35
Employee role: Senior Developer
```

## Key Takeaways

- Structs group related fields into a named type.
- Prefer keyed literals for readable and maintainable initialization.
- A struct's zero value initializes every field automatically.
- Pointer receivers can mutate the original struct; value receivers read a copy.
- `NewType` functions are the usual replacement for constructors.
- Anonymous structs are useful for local, short-lived data.
- Embedding promotes fields and methods, but it is composition rather than inheritance.
- A struct is comparable only when all of its fields are comparable.
