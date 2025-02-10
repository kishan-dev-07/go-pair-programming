package main

import "fmt"

func main(){
	var num1,num2,num3 int;
	fmt.Println("enter a number1:")
	fmt.Scan(&num1)
	fmt.Println("enter a number2:")
	fmt.Scan(&num2)
	fmt.Println("enter a number3:")
	fmt.Scan(&num3)

	
	if (num1>num2){
		if num1>num3{
			fmt.Printf("%d is the largest number.",num1)
		}else{
			fmt.Printf("%d is the largest number.",num3)
		}
	}else{
		if num2>num3{
            fmt.Printf("%d is the largest number.",num2)
        }else{
            fmt.Printf("%d is the largest number.",num3)
        }
	}

}