package main

import (
	"fmt"
	"os"
	"strconv"
)

func fib(n int) int {
	if (n == 0) || (n == 1) {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func printFib(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	printFib(n)
}
