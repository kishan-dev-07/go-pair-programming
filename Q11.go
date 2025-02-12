// Implement a program to check if a number is prime or not using conditionals and loops.

package main

import "fmt"

func main() {
	var num int

    fmt.Print("Enter a number: ")
    fmt.Scanln(&num)

	divCount := 0
    for i:=2; i<num/2; i++ {
		if num % i == 0 {
			divCount++
		}
	}

	if divCount == 0 {
        fmt.Printf("%d is a prime number.\n", num)
    } else {
        fmt.Printf("%d is not a prime number.\n", num)
    }
}