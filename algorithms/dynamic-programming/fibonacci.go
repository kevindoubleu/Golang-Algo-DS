package main

import "fmt"

func fib(n int) int {
	fmt.Println("executing fib(", n, ")")
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fib(i))
	}
}