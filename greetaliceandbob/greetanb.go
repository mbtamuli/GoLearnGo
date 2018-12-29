package main

import "fmt"

// Write a program that asks the user for their name and greets them with their
// name - Only if their names are Alice or Bob
func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)
	if name == "Alice" || name == "Bob" {
		fmt.Printf("Hello, %s", name)
	}
}
