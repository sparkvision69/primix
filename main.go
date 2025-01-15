package main

import (
	"fmt"
	"github.com/sparkvision69/primix/http"
	"github.com/sparkvision69/primix/process"
)

func main() {
	process.Print("This is a message printed from the process package!")

	err := http.StartServer("8080")
	
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
