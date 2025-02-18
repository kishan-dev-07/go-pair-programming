//Implement a function to merge two slices into a single slice in sorted order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{1, 4, 3, 4, 5}
	slice2 := []int{6, 7, 8, 0, 10}


	for i := 0; i < len(slice2); i++ {
		slice1 = append(slice1, slice2[i]) 
	}
	sort.Ints(slice1)
	fmt.Println("sorted slices:",slice1)
}
