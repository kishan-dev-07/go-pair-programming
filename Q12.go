//Create a program to calculate the power of a number using a loop (without using the math package).
package main

import "fmt"

func main(){
	var num, pow int
	fmt.Println("enter a number:")
	fmt.Scan(&num)
	fmt.Println("enter the power:")
	fmt.Scan(&pow)
	var result int 
	result = num
	for i:= 0; i< pow-1; i++{
	    result = result * num
	}
	fmt.Println("power :",result)
}