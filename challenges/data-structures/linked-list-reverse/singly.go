package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) prepend(n Node) {
	n.next = ll.head
	ll.head = &n
}

func (ll LinkedList) print() {
	var current *Node
	for current = ll.head; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.data)
	}
	fmt.Printf("%v\n", current)
}
