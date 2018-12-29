package main

import "fmt"

// Write a program that asks the user for a number n and prints the sum of the
// numbers 1 to n
func main() {
	num, sum := 0, 0
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		sum += i
	}
	fmt.Printf("The sum of the numbers from 1 to %d = %d", num, sum)
}
