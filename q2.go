// Implement a number guessing game where the program generates a random number, and the user has to guess it with hints (higher/lower).

package main
 
import (
    "fmt"
    "math/rand"
)

func  main() {
	target := rand.Intn(100)
	fmt.Println("I've chosen a random number between 1 and 100")
	// fmt.Println("target: ", target)
	fmt.Println("Guess the number!")
	for {
		var guess int
		fmt.Scanf("%d", &guess)
		if guess < target {
			fmt.Println("Give a Higher Try")
		} else if guess > target {
			fmt.Println("Give a Lower Try")
		} else {
			fmt.Println("You got it!")
			break
		}
	}
}