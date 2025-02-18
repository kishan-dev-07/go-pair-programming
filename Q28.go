//Write a Go program to implement a phone directory using a map where names are
//mapped to phone numbers.

package main

import "fmt"

func main(){
	var directory = make(map[string]int32)
	directory["tushar"]=954321723
	directory["kishan"]=952133213
	directory["nitish"]=954312323
	directory["sid"]=91233213
	fmt.Println("phone directory:")
	for name, number:= range(directory){
		fmt.Println(name,":",number)
	}

}