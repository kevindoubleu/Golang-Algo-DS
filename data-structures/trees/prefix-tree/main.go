package main

import "fmt"

func main() {
	t1 := newTrie()
	words := []string{"flower","flow","flight"}
	for _, v := range words {
		t1.insert(v)
	}
	t1.print()

	fmt.Println(t1.search("flo")) // false
	fmt.Println(t1.search("flow")) // true
	fmt.Println(t1.search("flowe")) // false

	fmt.Println(t1.startsWith("flo")) // true
	fmt.Println(t1.startsWith("flow")) // true
	fmt.Println(t1.startsWith("flowers")) // false
}