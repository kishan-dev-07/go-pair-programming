//Implement a Go program that simulates a traffic light system, where the user enters a
//color ("Red", "Yellow", "Green"), and the program outputs the appropriate action.

package main

import (
    "fmt"
)

func main(){
	var color string
	fmt.Println("enter the color to simulate the traffic system(Red/Green/Yellow):")
	fmt.Scan(&color)
	if(color == "Red"){
		fmt.Println("Stop!")
	}else if(color == "Yellow"){
		fmt.Println("Warning!")
    }else if(color == "Green"){
		fmt.Println("Go!")
	}else{
		fmt.Println("enter a valid color!")
	}
}