// Write a Go program to remove duplicate elements from a slice.

package main

import (
    "fmt"
)

func main() {
	slice := []int{1, 2, 3, 2, 4, 5, 3, 6, 7, 8, 9, 10, 2, 3}

	uniqueSlice := []int{}
	uniqueMap := make(map[int]bool)

	for _, num := range slice {
		if _, value := uniqueMap[num]; !value {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}
	
    fmt.Println("Original slice:", slice)
    fmt.Println("Unique slice:", uniqueSlice)

}