package main

import (
	"fmt"
	"os"
)

// run : go run hello.go <YourFirstName>

func main() {
	name := os.Args[1]
	fmt.Println(greeting(name))
}

func greeting(name string) string {
	if name == "" {
		name = "you!"
	}
	return fmt.Sprintf("Hey %s, welcome to the project!", name)
}
