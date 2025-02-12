package main

import "fmt"

// Function to compute GCD using the Euclidean Algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b // Swap a with b, and b with a % b
	}
	return a
}

func main() {
	var num1, num2 int

	// Taking input
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Handling negative numbers
	if num1 < 0 {
		num1 = -num1
	}
	if num2 < 0 {
		num2 = -num2
	}

	// Compute GCD
	result := gcd(num1, num2)

	// Output the result
	fmt.Printf("GCD of %d and %d is: %d\n", num1, num2, result)
}
