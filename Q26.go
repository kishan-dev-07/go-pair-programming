//Create a program that demonstrates slicing operations (append, delete, slice a slice)
//with a user-defined slice.

package main

import "fmt"

func main(){
	var slice1=[]int{1,2,3,4}
	fmt.Println("original slice:",slice1)
	slice1 = append(slice1,5)
	fmt.Println("after appending 5 to the slice:",slice1)
	fmt.Println("a part of the slice with the firsst 2 numbers is:",slice1[:2])
	var index = 2
	slice1 = append(slice1[:index],slice1[index+1:]...)
	fmt.Println("after deleting the third number:",slice1)
}