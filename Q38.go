//Develop a Go program that reads a list of student names and scores from a file into a
//map and allows searching by student name.

package main

import(
	"fmt"
	"os"
)

func main(){
	var data,err = os.ReadFile("Q38text.txt")
	if err!=nil{
		fmt.Println("failed to read file!")
	}
	fmt.Println("the marks are:")
	fmt.Printf("%s",data)
}