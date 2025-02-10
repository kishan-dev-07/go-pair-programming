// Write a program to generate the first N Fibonacci numbers using a loop.

package main

import (
    "fmt"
)

func main() {
	var n int
	fmt.Print("Enter value of n: ")
	fmt.Scanln(&n)
	if n<=0 {
		fmt.Println("Enter a valid number")
		return
	}

	if n == 1{
		fmt.Print("0")
	} else if n == 2 {
		fmt.Print("0 1")
	} else {
		fmt.Print("0 1 ")
		a, b := 0, 1
		var c int
		for i:=2; i<n; i++ {
            c = a + b
            fmt.Print(c, " ")
            a, b = b, c
        }
	}
	
}