// Write a program that finds the second-largest number in a given list of integers.

package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the number of elements in the list: ")
	fmt.Scanln(&n)
	var numbers []int

	for i:=0; i<n; i++ {
		var element int
		fmt.Printf("Enter element %d: ", i+1)
		fmt.Scanln(&element)

		numbers = append(numbers, element)
	}

	var largest, secondLargest int
	largest, secondLargest = 0, 0 

	for _, num := range numbers {
		if num > largest {
            secondLargest = largest
            largest = num
        } else if num > secondLargest && num != largest {
            secondLargest = num
        }
	}

	fmt.Printf("The second-largest number in the list is: %d\n", secondLargest)
}