//Write a Go program to calculate and print the sum of all even numbers from 1 to N
using a loop.
package main

import "fmt"

func main(){
	var num int
	fmt.Print("enter a number:")
	fmt.Scan(&num)
	var sum int
	for i:=0 ; i<= num ; i++{
		if i%2==0{
			sum=sum+i
		}else{
			continue
		}
	}
	fmt.Printf("the sum of all the even numbers untill %d are %d",num,sum)
}