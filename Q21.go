// Implement a program to store and display a list of students’ names and marks using arrays.

package main

import "fmt"

func main() {
    const students = 5
    var names [students]string
    var marks [students]int

    fmt.Println("Enter 5 student names:")
    for i := 0; i < students; i++ {
        fmt.Scanln(&names[i])
    }

    fmt.Println("\nEnter thier marks in Maths:")
    for i := 0; i < students; i++ {
        fmt.Scan(&marks[i])
    }

    fmt.Println("\nStudent Name\tMarks")
    for i := 0; i < students; i++ {
        fmt.Printf("%s\t\t%d\n", names[i], marks[i])
    }
}