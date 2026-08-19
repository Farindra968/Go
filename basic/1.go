package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, world!")

	// Sqrt returns the square root of a number.
	fmt.Println("Square root of 16:", math.Sqrt(16))

	// Pow returns the first number raised to the power of the second number.
	fmt.Println("2 raised to the power of 3:", math.Pow(2, 3))

	// Abs returns the absolute, or non-negative, value of a number.
	fmt.Println("Absolute value of -5:", math.Abs(-5))

	// Max returns the larger of two numbers.
	fmt.Println("Maximum of 10 and 20:", math.Max(10, 20))

	// Min returns the smaller of two numbers.
	fmt.Println("Minimum of 10 and 20:", math.Min(10, 20))

	// Round rounds a decimal number to the nearest integer.
	fmt.Println("Rounded value of 3.7:", math.Round(3.7))

	// Floor rounds a number down to the nearest integer.
	fmt.Println("Floor value of 3.7:", math.Floor(3.7))

	// Ceil rounds a number up to the nearest integer.
	fmt.Println("Ceiling value of 3.7:", math.Ceil(3.7))

	// Trunc removes the decimal part of a number.
	fmt.Println("Truncated value of 3.7:", math.Trunc(3.7))

	// Log returns the natural logarithm of a number.
	fmt.Println("Natural logarithm of 10:", math.Log(10))

	// Log10 returns the base-10 logarithm of a number.
	fmt.Println("Base-10 logarithm of 100:", math.Log10(100))

	// Log2 returns the base-2 logarithm of a number.
	fmt.Println("Base-2 logarithm of 8:", math.Log2(8))

	// Exp returns e raised to the specified power.
	fmt.Println("e raised to the power of 2:", math.Exp(2))

	// Expm1 returns e raised to x minus 1.
	fmt.Println("e raised to x minus 1:", math.Expm1(1))

	// Log1p returns the natural logarithm of 1 plus x.
	fmt.Println("Natural logarithm of 1 plus x:", math.Log1p(1))

	// Sin returns the sine of an angle in radians.
	fmt.Println("Sine of π/2:", math.Sin(math.Pi/2))

	// Cos returns the cosine of an angle in radians.
	fmt.Println("Cosine of π:", math.Cos(math.Pi))

	// Tan returns the tangent of an angle in radians.
	fmt.Println("Tangent of π/4:", math.Tan(math.Pi/4))

	// Asin returns the inverse sine of a value.
	fmt.Println("Arcsine of 1:", math.Asin(1))

	// Acos returns the inverse cosine of a value.
	fmt.Println("Arccosine of 0:", math.Acos(0))

	// Atan returns the inverse tangent of a value.
	fmt.Println("Arctangent of 1:", math.Atan(1))

	// Atan2 returns the angle formed by the coordinates y and x.
	fmt.Println("Arctangent of y/x:", math.Atan2(1, 1))

	// Sinh returns the hyperbolic sine of a value.
	fmt.Println("Hyperbolic sine of 1:", math.Sinh(1))

	// Cosh returns the hyperbolic cosine of a value.
	fmt.Println("Hyperbolic cosine of 1:", math.Cosh(1))

	// Tanh returns the hyperbolic tangent of a value.
	fmt.Println("Hyperbolic tangent of 1:", math.Tanh(1))

	// Asinh returns the inverse hyperbolic sine of a value.
	fmt.Println("Inverse hyperbolic sine of 1:", math.Asinh(1))

	// Acosh returns the inverse hyperbolic cosine of a value.
	fmt.Println("Inverse hyperbolic cosine of 2:", math.Acosh(2))

	// Atanh returns the inverse hyperbolic tangent of a value.
	fmt.Println("Inverse hyperbolic tangent of 0.5:", math.Atanh(0.5))

	// Cbrt returns the cube root of a number.
	fmt.Println("Cube root of 27:", math.Cbrt(27))

	// Mod returns the remainder after floating-point division.
	fmt.Println("Remainder of 10 divided by 3:", math.Mod(10, 3))

	// Hypot returns the hypotenuse of a right-angled triangle.
	fmt.Println("Hypotenuse of sides 3 and 4:", math.Hypot(3, 4))

	// Dim returns the positive difference between two numbers.
	fmt.Println("Positive difference between 10 and 4:", math.Dim(10, 4))

	// Copysign returns the first value with the sign of the second value.
	fmt.Println("Sign of -10 copied to 5:", math.Copysign(5, -10))

	// Signbit reports whether a number is negative.
	fmt.Println("Is -5 negative:", math.Signbit(-5))

	// IsNaN reports whether a value is not a number.
	fmt.Println("Is NaN:", math.IsNaN(math.NaN()))

	// IsInf reports whether a value represents positive or negative infinity.
	fmt.Println("Is positive infinity:", math.IsInf(math.Inf(1), 1))

	// Int returns a random non-negative integer.
	fmt.Println("Random integer:", rand.Int())

	// Intn returns a random integer from 0 up to, but not including, n.
	fmt.Println("Random integer from 0 to 29:", rand.Intn(30))

	// Int31 returns a random non-negative 31-bit integer.
	fmt.Println("Random 31-bit integer:", rand.Int31())

	// Int31n returns a random 31-bit integer from 0 up to n.
	fmt.Println("Random 31-bit integer from 0 to 99:", rand.Int31n(100))

	// Int63 returns a random non-negative 63-bit integer.
	fmt.Println("Random 63-bit integer:", rand.Int63())

	// Int63n returns a random 63-bit integer from 0 up to n.
	fmt.Println("Random 63-bit integer from 0 to 999:", rand.Int63n(1000))

	// Float32 returns a random float32 from 0.0 up to, but not including, 1.0.
	fmt.Println("Random float32:", rand.Float32())

	// Float64 returns a random float64 from 0.0 up to, but not including, 1.0.
	fmt.Println("Random float64:", rand.Float64())

	// NormFloat64 returns a normally distributed random number.
	fmt.Println("Normally distributed random number:", rand.NormFloat64())

	// ExpFloat64 returns an exponentially distributed random number.
	fmt.Println("Exponentially distributed random number:", rand.ExpFloat64())

	// Perm returns a random permutation of integers from 0 to n-1.
	fmt.Println("Random permutation:", rand.Perm(5))

	// Shuffle randomly rearranges the elements of a slice.
	numbers := []int{1, 2, 3, 4, 5}
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	fmt.Println("Shuffled numbers:", numbers)

	// Read fills a byte slice with random bytes.
	randomBytes := make([]byte, 8)
	count, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return
	}

	fmt.Printf("Random bytes written: %d, data: %v\n", count, randomBytes)
}
