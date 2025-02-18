//Create a Go program that reverses an array without using built-in functions.

package main

import "fmt"

func main(){
	var string1 string
	fmt.Println("Enter the string:")
	fmt.Scan(&string1)
	runes:=[]rune(string1)
	fmt.Println("the reverse string is:")
	for i:=range runes{
		fmt.Printf("%c",runes[len(runes)-i-1])
	}
}