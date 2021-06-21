package main

import "fmt"

func incrementToN(i, n int) {
	if i > n {
		return
	} else {
		fmt.Println(i)
		incrementToN(i+1, n)
	}
}

func main() {
	incrementToN(1,100)
}