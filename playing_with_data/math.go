package main

import (
	"fmt"
	"math"
)

func main() {
	// 1. Basic arithmetic functions
	fmt.Println("Square Root of 16:", math.Sqrt(16))         // Square root
	fmt.Println("Power of 2^3:", math.Pow(2, 3))             // Exponentiation
	fmt.Println("Absolute value of -3.14:", math.Abs(-3.14)) // Absolute value
	fmt.Println("Ceiling of 3.14:", math.Ceil(3.14))         // Round up to nearest integer
	fmt.Println("Floor of 3.14:", math.Floor(3.14))          // Round down to nearest integer

	// 2. Rounding functions
	fmt.Println("Round 3.14159:", math.Round(3.14159))      // Rounding to the nearest integer
	fmt.Println("Truncating 3.14159:", math.Trunc(3.14159)) // Removing the fractional part
}