package main

import "fmt"

func memoize(f func(param int) int) func(param int) int {
	cache := make(map[int]int)
	return func (param int) int {
		_, exist := cache[param]
		if !exist {
			fmt.Println("creating cache for param", param)
			cache[param] = f(param)
		}
		return cache[param]
	}
}

func fib(n int) int {
	fmt.Println("executing fib(", n, ")")
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	memoizedFib := memoize(fib)
	
	for i := 0; i < 5; i++ {
		fmt.Println(memoizedFib(i))
	}
	for i := 0; i < 5; i++ {
		fmt.Println(memoizedFib(i))
	}
}