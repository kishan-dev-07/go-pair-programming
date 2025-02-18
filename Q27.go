// Implement a Go program that counts the frequency of each character in a given string using a map.

package main

import (
    "fmt"
)

func main() {
	var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)

    charCount := make(map[rune]int)

    for _, char := range input {
        charCount[char]++
    }

    fmt.Println("Character frequencies:")
    for char, count := range charCount {
        fmt.Printf("%c: %d\n", char, count)
    }
}