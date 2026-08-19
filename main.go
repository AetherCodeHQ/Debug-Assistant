package main

import (
	"fmt"
	"os"
)

// debug_assistant - Debug tracing utilities
func debug_assistant(path string) {
	fmt.Println("========================================")
	fmt.Println("  Debug-Assistant")
	fmt.Println("  Debug tracing utilities")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	debug_assistant(path)
}
