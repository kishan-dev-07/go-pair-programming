//Create a Go program that reads a string input and determines if it is a palindrome.
package main

import (
	"fmt"
)

func main(){
	var string1 string
	fmt.Printf("Enter the string:")
	fmt.Scan(&string1)

	runes:=[]rune(string1)
	for i :=0 ; i<len(runes)/2; i++ {
		if runes[i]!= runes[len(runes)-i-1] {
			fmt.Printf("not a palindrome")
			return 
		}
	}
	fmt.Printf("Palindrome")
	
}