package main

import "fmt"

// Modify the previous program such that only multiples of three or five are
// considered in the sum, e.g. 3, 5, 6, 9, 10, 12, 15 for n=17
func main() {
	num, sum := 0, 0
	fmt.Printf("Enter a number: ")
	fmt.Scanf("%d", &num)
	for i := 1; i <= num; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("The sum of the numbers from 1 to %d, which are multiples"+
		" of three or five = %d", num, sum)
}
