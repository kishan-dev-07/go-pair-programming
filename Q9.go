//Implement a basic calculator that takes two numbers and an operator (+, -, *, /) and
//performs the operation using a switch statement.
package main

import (
    "fmt"
)

func main(){
	var num1,num2 int
	fmt.Println("Enter number1")
	fmt.Scan(&num1)
	fmt.Println("Enter number2")
	fmt.Scan(&num2)
	fmt.Println("1. + ")
	fmt.Println("2. - ")
	fmt.Println("3. * ")
	fmt.Println("4. / ")
	var choice int
	fmt.Println("enter your choice(1/2/3/4):")
	fmt.Scan(&choice)
	switch choice {
		case 1: fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
		case 2: fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
		case 3: fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
		case 4: if num2!=0 {
                fmt.Printf("%d / %d = %f\n", num1, num2, float64(num1)/float64(num2))
            } else {
                fmt.Println("Cannot divide by zero")
            }
    }
}