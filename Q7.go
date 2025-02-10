// Develop a Go program that calculates the factorial of a number using recursion.

package main

import (
    "fmt"
)

func fact(n int) int {
	if n == 1 {
		return 1
	}

	return n * fact(n-1)
}
func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)
	fmt.Println("Factorial of", n, "is", fact(n))
}
