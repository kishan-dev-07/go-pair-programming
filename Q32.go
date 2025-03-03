//Create a Go program that uses an anonymous struct to store temporary user
// information and display it.

package main

import "fmt"

func main(){

	employee := struct{
		name string
		age int
		fees int
	}{
		name:"Kishan",
		age:25,
        fees:50000,
	}

	fmt.Printf("Employee Name: %s\n", employee.name)
	fmt.Printf("Employee Age: %d\n", employee.age)
	fmt.Printf("Employee Fees: %d\n", employee.fees)
}