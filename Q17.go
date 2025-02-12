// Develop a Go program that determines whether a given ISBN number is valid or not (ISBN-10 checksum verification).

package main

import (
    "fmt"
	"strconv"
)

func main() {
	var isbn string
    fmt.Print("Enter an ISBN-10 number: ")
    fmt.Scan(&isbn)

    if len(isbn) !=10 {
        fmt.Println("Invalid ISBN-10 number")
        return
    }

    sum := 0
    for i:=0; i<9; i++ {
        digit, _ := strconv.Atoi(string(isbn[i]))
        sum += digit * (10 - i)
    }

    checkDigit, _ := strconv.Atoi(string(isbn[9]))
    if sum%11 !=checkDigit {
        fmt.Println("Invalid ISBN-10 number")
    } else {
        fmt.Println("Valid ISBN-10 number")
    }
}