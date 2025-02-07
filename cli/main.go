package main

import (
	"fmt"
	"os"
)

const version = "1.0.0"

func main() {
	// Check if an argument is passed
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-v", "--version":
			fmt.Println("Primix version 1")
			return
		default:
			fmt.Println("Unknown command:", os.Args[1])
			fmt.Println("Usage: primix -v")
			os.Exit(1)
		}
	}

	fmt.Println("Usage: primix -v")
}
