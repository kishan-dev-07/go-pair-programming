// Implement a Go program to convert a given decimal number to binary using loops and conditionals.

package main

import (
    "fmt"
    "strconv"
)

func main() {
	var decimal int
	fmt.Print("Enter a decimal number: ")
	fmt.Scanln(&decimal)

	binary := ""
	for decimal > 0 {
		// adding it as concatination
		binary = strconv.Itoa(decimal%2) + binary
		decimal /= 2
	}

	fmt.Println("Binary number:", binary)
}