//Write a Go program that simulates rolling two dice and prints the outcome, ensuring
//fair randomness.
package main

import (
	"fmt"
	"math/rand"
		)
func main(){
	var dice1,dice2	int
	dice1=rand.Intn(6)+1
	dice2=rand.Intn(6)+1
	fmt.Println("Dice 1:",dice1)
	fmt.Println("Dice 2:",dice2)
}