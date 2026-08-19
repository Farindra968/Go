package main

import "fmt"

// Person demonstrates a struct type.
type Person struct {
	Name string
	Age  int
}

// Stringer demonstrates an interface type.
// It uses the String method to return a string representation of the type.
// The Stringer interface is implemented by any type that has a String method with the signature String() string.
type Stringer interface {
	String() string
}

// Product implements the Stringer interface.
type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s: $%.2f", p.Name, p.Price)
}


func main() {
	// Signed integer types store whole positive and negative numbers.
	var integer int = 42 // The maximum value for int is platform-dependent, but it is typically 2^31 - 1 on 32-bit systems and 2^63 - 1 on 64-bit systems.
	var integer8 int8 = -128 // The maximum value for int8 is 127, which is 2^7 - 1.
	var integer16 int16 = 32000 // The maximum value for int16 is 32767, which is 2^15 - 1.
	var integer32 int32 = 2147483647 // The maximum value for int32 is 2147483647, which is 2^31 - 1.
	var integer64 int64 = 9223372036854775807 // The maximum value for int64 is 9223372036854775807, which is 2^63 - 1.

	// Unsigned integer types store only zero and positive numbers.
	var unsigned uint = 100 // uint is an alias for uint32 or uint64 depending on the platform. Its maximum value is platform-dependent, but it is typically 2^32 - 1 on 32-bit systems and 2^64 - 1 on 64-bit systems.
	var byteValue byte = 'A' // byte is an alias for uint8, which can store values from 0 to 255. The maximum value for uint8 is 255, which is 2^8 - 1.
	var uint16Value uint16 = 65535 // The maximum value for uint16 is 65535, which is 2^16 - 1.

	// Floating-point types store decimal numbers.
	var decimal32 float32 = 3.14
	var decimal64 float64 = 99.99

	// Complex types store real and imaginary components.
	var complex64Value complex64 = 2 + 3i
	var complex128Value complex128 = 5 + 7i

	// bool stores either true or false.
	var active bool = true

	// string stores a sequence of UTF-8 characters.
	var message string = "Hello, Go!"

	// rune is an alias for int32 and stores a Unicode code point.
	var character rune = '世'

	// Arrays have a fixed length.
	numbers := [3]int{10, 20, 30}

	// Slices are flexible, dynamic views of arrays.
	scores := []int{80, 90, 95}
	scores = append(scores, 100)

	// Maps store key-value pairs.
	student := map[string]int{
		"Alice": 90,
		"Bob":   85,
	}

	// Structs group related values into one type.
	person := Person{Name: "Alice", Age: 25}

	// Pointers store the memory address of another variable.
	age := 30 // It is also called handling a variable by reference. A pointer is a variable that holds the memory address of another variable. In Go, pointers are denoted by the * operator, which is used to declare a pointer type and to dereference a pointer to access the value it points to. The & operator is used to get the memory address of a variable. Pointers are useful for passing large structs or arrays to functions without copying them, and for modifying variables in place.
	agePointer := &age
	*agePointer = 31

	// A function can be stored in a variable.
	add := func(a, b int) int {
		return a + b
	}

	// An interface can hold a value of any type.
	var value any = "interface value"

	// Type assertion extracts the concrete value from an interface.
	text, ok := value.(string)
	if ok {
		fmt.Println("Asserted string:", text)
	}

	// An interface value can contain a struct that implements its methods.
	var item Stringer = Product{Name: "Keyboard", Price: 49.99}

	fmt.Println("int:", integer)
	fmt.Println("int8:", integer8)
	fmt.Println("int16:", integer16)
	fmt.Println("int32:", integer32)
	fmt.Println("int64:", integer64)
	fmt.Println("uint:", unsigned)
	fmt.Println("byte:", byteValue)
	fmt.Println("uint16:", uint16Value)
	fmt.Println("float32:", decimal32)
	fmt.Println("float64:", decimal64)
	fmt.Println("complex64:", complex64Value)
	fmt.Println("complex128:", complex128Value)
	fmt.Println("bool:", active)
	fmt.Println("string:", message)
	fmt.Println("rune:", character)
	fmt.Println("array:", numbers)
	fmt.Println("slice:", scores)
	fmt.Println("map:", student)
	fmt.Println("struct:", person)
	fmt.Println("pointer value:", *agePointer)
	fmt.Println("function result:", add(3, 5))
	fmt.Println("interface:", value)
	fmt.Println("Stringer:", item.String())

	
}
