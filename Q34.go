//34. Write a program to demonstrate the use of multidimensional slices to store student
//marks across different subjects.

package main

import "fmt"

func main(){
	marks:=[][]int{
		{85,90,88},
	}

	fmt.Println("the marks for science subject are:",marks[0][0])
	fmt.Println("the marks for english subject are:",marks[0][1])
	fmt.Println("the marks for go subject are:",marks[0][2])

}