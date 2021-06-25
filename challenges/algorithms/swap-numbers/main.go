package main

import "fmt"

func swapInt(a, b *int) {
	*b = *b+*a
	*a = *b-*a
	*b = *b-*a
}

func main() {
	a := 4
	b := 6
	fmt.Println("a =", a, "b =", b)
	swapInt(&a,&b)
	fmt.Println("a =", a, "b =", b)
}