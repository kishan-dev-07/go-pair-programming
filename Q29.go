// Develop a Go program that takes a list of words and counts how many times each word appears.

package main

import (
    "fmt"
    "strings"
	"bufio"
	"os"
)

func main() {
	words := []string{}
    wordCount := make(map[string]int)

	fmt.Println("Enter a list of words separated by spaces:")

	// Use bufio to read the entire input line
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim the newline character
	input = strings.TrimSpace(input)

	// fmt.Println("You entered:", input)
	words = strings.Fields(input)

    for _, word := range words {
        wordCount[word]++
    }

    fmt.Println("Word Count:")
    for word, count := range wordCount {
        fmt.Printf("%s: %d\n", word, count)
    }
}