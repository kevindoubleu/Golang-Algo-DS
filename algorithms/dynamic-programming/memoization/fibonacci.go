package main

import "fmt"

var cache = make(map[int]int)
func fib(n int) int {
	if n <= 1 {
		return n
	}

	_, exist := cache[n]
	if !exist {
		fmt.Println("executing fib(", n, ")")
		cache[n] = fib(n-1) + fib(n-2)
	}

	return cache[n]
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fib(i))
	}
}