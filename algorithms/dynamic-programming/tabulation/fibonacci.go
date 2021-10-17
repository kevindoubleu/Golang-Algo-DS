package main

import "fmt"

func fib(n int) int {
	table := []int{0, 1}
	for i := 2; i <= n; i++ {
		table = append(table, table[i-1] + table[i-2])
	}

	return table[n]
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fib(i))
	}
}