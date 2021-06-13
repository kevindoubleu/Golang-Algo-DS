package main

import "fmt"

type Node struct {
	prev *Node
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) append(n Node) {
	// list empty. head and tail is nil
	if ll.head == nil && ll.tail == nil {
		ll.head = &n
		ll.tail = &n
	// list has 1 node, head and tail pointing to same node
	// list has >1, normal
	} else {
		n.prev = ll.tail
		ll.tail.next = &n

		ll.tail = &n
	}
}

// func (ll *LinkedList) prepend(n Node) {
// 	if ll.head == nil && ll.tail == nil {
// 		ll.head = &n
// 		ll.tail = &n
// 	} else {
// 		n.next = ll.head
// 		ll.head.prev = &n

// 		ll.head = &n
// 	}
// }

// // returns pointer so we can edit the nodes
// func (ll *LinkedList) first() *Node {
// 	return ll.head
// }

// func (ll *LinkedList) last() *Node {
// 	return ll.tail
// }

// func (ll *LinkedList) popFirst() *Node {
// 	tmp := ll.head
// 	if ll.head == ll.tail {
// 		ll.head, ll.tail = nil, nil
// 	} else {
// 		ll.head = ll.head.next
// 		ll.head.prev = nil
// 	}
// 	return tmp
// }

// func (ll *LinkedList) popLast() *Node {
// 	tmp := ll.tail
// 	if ll.head == ll.tail {
// 		ll.head, ll.tail = nil, nil
// 	} else {
// 		ll.tail = ll.tail.prev
// 		ll.tail.next = nil
// 	}
// 	return tmp
// }

// func (ll *LinkedList) clear() {
// 	// in golang, circular references can be garbage collected
// 	ll.head = nil
// 	ll.tail = nil
// }

// func (ll *LinkedList) get(index int) *Node {
// 	current := ll.head
// 	for i := 0; i < index; i++ {
// 		// if out of bounds
// 		if current == nil {
// 			return nil
// 		}
// 		current = current.next
// 	}
// 	return current
// }

// func (ll *LinkedList) remove(index int) *Node {
// 	current := ll.head
// 	for i := 0; i < index-1; i++ {
// 		if current == nil {
// 			return nil
// 		}
// 		current = current.next
// 	}

// 	// edge cases
// 	if index == 0 {
// 		return ll.popFirst()
// 	} else if current.next.next == nil {
// 		return ll.popLast()
// 	// the node we want to remove is current.next
// 	// because the loop is only until index-1
// 	} else {
// 		tmp := current.next

// 		current.next.next.prev = current
// 		current.next = current.next.next

// 		return tmp
// 	}
// }

// func (ll *LinkedList) insert(index int, n Node) {
// 	// edge cases
// 	if index == 0 {
// 		ll.prepend(n)
// 	} else if index >= ll.size() {
// 		ll.append(n)
// 	// inserting in the middle of list
// 	} else {
// 		current := ll.head
// 		for i := 0; i < index-1; i++ {
// 			current = current.next
// 		}
// 		// we want to put new node in current.next
// 		// because the loop is only until index-1
// 		n.next = current.next
// 		n.prev = current
// 		current.next.prev = &n
// 		current.next = &n
// 	}
// }

func (ll LinkedList) print() {
	var current *Node
	for current = ll.head; current != nil; current = current.next {
		fmt.Printf("%p <- %p(%v) -> %p\n", current.prev, current, current.data, current.next)
	}
}

// func (ll LinkedList) size() int {
// 	len := 0
// 	for current := ll.head; current != nil; current = current.next {
// 		len++
// 	}
// 	return len
// }

// for use as a loop iterator
func (ll *LinkedList) getIterator() func() *Node {
	current := ll.head
	return func() *Node {
		if current != nil {
			defer func() {
				current = current.next
			}()
			return current
		} else {
			return nil
		}
	}
}
