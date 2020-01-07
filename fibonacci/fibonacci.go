package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	if num, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Println("Expected a number. \nUsage: " + os.Args[0] + " 42")
	} else {
		i := 1
		for i < num {
			fmt.Printf("%d ", fib(i))
			i++
		}
	}
	fmt.Println()
}
