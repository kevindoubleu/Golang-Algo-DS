package main

import "fmt"

func main() {
	t1 := newTrie()
	t1.insert("apple")
	t1.insert("apploo")
	t1.print()

	fmt.Println(t1.search("app"))
	fmt.Println(t1.search("opp"))
	fmt.Println(t1.search("apple"))
	fmt.Println(t1.search("apples"))
}