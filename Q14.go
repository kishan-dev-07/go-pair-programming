//Implement a grading system that takes marks as input and assigns grades based on predefined categories using conditionals.

package main

import "fmt"

func main(){
	var num int
	fmt.Println("welcome to this grading system!")
	fmt.Println("Please enter your number:")
	fmt.Scanln(&num)
	if num == 0 || num >100{
		fmt.Println("enter a valid number!")
		return
	}else if num<40{
		fmt.Println("You have failed!")
	}else if num >50 && num < 60{
		fmt.Println("grade achieved - D")
	}else if num >61 && num < 70{
		fmt.Println("grade achieved - C")
	}else if num >71 && num < 80{
		fmt.Println("grade achieved - B")
	}else if num >81{
		fmt.Println("grade achieved - A")
	}
}