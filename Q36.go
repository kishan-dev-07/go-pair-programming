//Implement a Go program that merges two sorted slices into a single sorted slice
//without using built-in sorting functions.

package main

import "fmt"

func main(){

	slice1:=[]int{1,5,10}
	slice2:=[]int{2,4,6}

	var slice3 = append(slice1,slice2...)
	fmt.Println(slice3)
	
	for i:=0;i<len(slice3);i++{
		for j:=i+1;j<len(slice3);j++{
			if slice3[i]>slice3[j]{
				slice3[i],slice3[j]=slice3[j],slice3[i]
            }
	    }			
	}

	fmt.Println(slice3)
}