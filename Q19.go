package main

import (
	"fmt"
	"time"
)

func main() {
	var command string
	var elapsed time.Duration
	var running bool
	var start time.Time

	fmt.Println("⏱️  Basic Stopwatch (Type 'start', 'stop', 'reset', 'exit')")

	for {
		fmt.Print("\nEnter command: ")
		fmt.Scan(&command)

		switch command {
		case "start":
			if !running {
				start = time.Now()
				running = true
				fmt.Println("▶ Stopwatch started...")
			} else {
				fmt.Println("⚠️ Stopwatch is already running!")
			}

		case "stop":
			if running {
				elapsed += time.Since(start)
				running = false
				fmt.Printf("⏸️ Stopwatch stopped. Elapsed time: %.2f seconds\n", elapsed.Seconds())
			} else {
				fmt.Println("⚠️ Stopwatch is not running!")
			}

		case "reset":
			elapsed = 0
			running = false
			fmt.Println("🔄 Stopwatch reset!")

		case "exit":
			fmt.Println("👋 Exiting stopwatch.")
			return

		default:
			fmt.Println("❌ Invalid command! Use 'start', 'stop', 'reset', or 'exit'.")
		}

		// If running, show the current elapsed time
		if running {
			fmt.Printf("⏳ Current time: %.2f seconds\n", time.Since(start).Seconds()+elapsed.Seconds())
		}
	}
}
