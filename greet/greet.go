package main

import "fmt"

// Write a program that asks the user for their name and greets them with their
// name.
func main() {
	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("Hello, %s", name)
}
