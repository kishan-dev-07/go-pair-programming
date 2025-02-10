package main

import (
    "fmt"
)

func main() {
	var n int 
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	var digitSum int
	fmt.Print("Sum of digits of ", n, " = ")
	for n > 0 {
		digitSum += n % 10
        n /= 10
	}
	fmt.Println(digitSum)
}	
