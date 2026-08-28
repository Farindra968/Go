package main

import "fmt"

// Person represents a person.
type Person struct {
	Name      string
	Age       int
	Address   string
	Amount    float32
	IsMarried bool
	Role      string
}

// ChangeMaritalStatus updates the person's marital status.
func (p *Person) ChangeMaritalStatus(status bool) {
	p.IsMarried = status
}

// ChangeRole updates the person's role.
func (p *Person) ChangeRole(role string) {
	p.Role = role
}

// Summary returns a short summary of the person.
func (p Person) Summary() string {
	return fmt.Sprintf("%s (%d), role: %s", p.Name, p.Age, p.Role)
}

// NewPerson is a constructor-like function for Person.
func NewPerson(
	name string,
	age int,
	address string,
	amount float32,
	isMarried bool,
	role string) *Person {
	return &Person{
		Name:      name,
		Age:       age,
		Address:   address,
		Amount:    amount,
		IsMarried: isMarried,
		Role:      role,
	}
}

// Employee demonstrates struct embedding.
// Person's fields and methods are promoted to Employee.
type Employee struct {
	Person
	EmployeeID string
	Department string
}

func main() {

	// --------------------------------------------------
	// 1. Basic struct initialization
	// --------------------------------------------------

	// A keyed literal is clear and does not depend on field order.
	person := Person{
		Name:      "John Doe",
		Age:       30,
		Address:   "123 Main St",
		Amount:    100.50,
		IsMarried: true,
		Role:      "User",
	}

	fmt.Println("Person:", person.Summary())
	fmt.Printf("Full value: %+v\n", person)

	// Calling methods on a struct.
	person.ChangeMaritalStatus(false)
	person.ChangeRole("Admin")

	fmt.Println("Updated person:", person.Summary())
	fmt.Printf("Updated value: %+v\n", person)

	// --------------------------------------------------
	// 2. Zero-value struct
	// --------------------------------------------------

	// The zero value of every field is used when
	// no literal is provided.
	var emptyPerson Person

	fmt.Printf("Zero value: %+v\n", emptyPerson)

	// --------------------------------------------------
	// 3. Constructor-like function
	// --------------------------------------------------

	createdPerson := NewPerson(
		"Jane Smith",
		28,
		"456 Elm St",
		200.75,
		false,
		"Manager",
	)

	fmt.Println("Constructed person:", createdPerson.Summary())

	// --------------------------------------------------
	// 4. Anonymous struct
	// --------------------------------------------------

	// Useful for short-lived local data.
	language := struct {
		name       string
		isCompiled bool
	}{
		name:       "Go",
		isCompiled: true,
	}

	fmt.Printf("Anonymous struct: %+v\n", language)

	// --------------------------------------------------
	// 5. Struct embedding
	// --------------------------------------------------

	employee := Employee{
		Person: Person{
			Name:      "Jack",
			Age:       35,
			Address:   "789 Oak St",
			Amount:    300.00,
			IsMarried: true,
			Role:      "User",
		},
		EmployeeID: "E123",
		Department: "IT",
	}

	fmt.Printf("Employee: %+v\n", employee)

	// Person's method is promoted to Employee.
	employee.ChangeRole("Senior Developer")

	fmt.Println("Promoted method:", employee.Summary())

	// Person's fields are also promoted.
	fmt.Println("Employee name:", employee.Name)
	fmt.Println("Employee age:", employee.Age)
	fmt.Println("Employee role:", employee.Role)
}
