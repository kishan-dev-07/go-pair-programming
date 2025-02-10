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