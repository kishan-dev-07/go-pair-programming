// Develop a Go program that simulates a simple ATM where the user can deposit, withdraw, or check their balance using loops and conditionals.

package main

import (
    "fmt"
)

func main() {
	var choice int
    var balance float64 = 0.0

	for {
		fmt.Println("\nATM Menu:")
        fmt.Println("1. Deposit")
        fmt.Println("2. Withdraw")
        fmt.Println("3. Check Balance")
        fmt.Println("4. Exit")
        fmt.Print("Enter your choice: ")
        fmt.Scan(&choice)
		
        switch choice {
			case 1: fmt.Println("Enter the amount to deposit: ")
					var deposit float64
					fmt.Scan(&deposit)
					balance += deposit
					fmt.Println("Amount deposited successfully.")
					fmt.Println("Current balance: ", balance)
			case 2: fmt.Println("Enter the amount to withdraw: ")
					var withdraw float64
					fmt.Scan(&withdraw)
					if withdraw > balance {
						fmt.Println("Insufficient balance.")
                    } else {
						balance -= withdraw
						fmt.Println("Amount withdrawn successfully.")
						fmt.Println("Current balance: ", balance)
					}
			case 3: fmt.Println("Current balance: ", balance)
			case 4: fmt.Println("Thank you for using our ATM. Goodbye!")
					return
			default: fmt.Println("Invalid choice. Please try again.")
		}
	}
}
