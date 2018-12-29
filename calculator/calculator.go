package main

import "fmt"

func main() {
	var a, b int
	var c string
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d", &a)
	fmt.Printf("Enter the second number: ")
	fmt.Scanf("%d", &b)
	fmt.Printf("Enter the operator: ")
	fmt.Scanf("%s", &c)
	switch c {
	case "+":
		fmt.Printf("%d + %d = %d", a, b, (a + b))
	case "-":
		fmt.Printf("%d - %d = %d", a, b, (a - b))
	case "*":
		fmt.Printf("%d * %d = %d", a, b, (a * b))
	case "/":
		fmt.Printf("%d / %d = %d", a, b, (a / b))
	default:
		fmt.Println("Error")
	}
}
