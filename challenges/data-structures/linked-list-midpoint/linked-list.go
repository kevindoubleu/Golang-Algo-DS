package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// func (ll *LinkedList) get(index int) *Node {
// 	current := ll.head
// 	for i := 0; i < index; i++ {
// 		if current == nil {
// 			return nil
// 		}
// 		current = current.next
// 	}
// 	return current
// }

func (ll *LinkedList) remove(index int) *Node {
	current := ll.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return nil
		}
		current = current.next
	}

	if index == 0 {
		tmp := ll.head
		if ll.head != nil {
			ll.head = ll.head.next
		}
		return tmp
	} else if index >= ll.size() {
		current := ll.head
		for current.next.next != nil {
			current = current.next
		}
		tmp := current.next
		current.next = nil
		return tmp
	} else {
		tmp := current.next
		current.next = current.next.next
		return tmp
	}
}

func (ll *LinkedList) insert(index int, n Node) {
	if index == 0 {
		n.next = ll.head
		ll.head = &n
	} else if index >= ll.size() {
		if ll.head == nil {
			ll.head = &n
		} else {
			last := ll.last()
			last.next = &n
		}
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		n.next = current.next
		current.next = &n
	}
}

func (ll *LinkedList) first() *Node {
	return ll.head
}

func (ll *LinkedList) last() *Node {
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (ll LinkedList) size() int {
	len := 0
	for current := ll.head; current != nil; current = current.next {
		len++
	}
	return len
}

func (ll LinkedList) print() {
	var current *Node
	for current = ll.head; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.data)
	}
	fmt.Printf("%v\n", current)
}