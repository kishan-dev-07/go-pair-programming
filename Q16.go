package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var countVow, countCons int
	fmt.Println("Enter a sentence:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Split(input, " ")
	fmt.Println("No. of words in this sentence are:", len(words))

	input = strings.ToLower(input)

	for _, char := range input {
		if strings.ContainsRune("aeiou", char) {
			countVow++
		} else if char >= 'a' && char <= 'z' {
			countCons++
		}
	}

	fmt.Println("No. of vowels are:", countVow)
	fmt.Println("No. of consonants are:", countCons)
}
