package main

import "fmt"

// functions here can be used by both
// singly linked list and doubly linked list

func (ll *LinkedList) get(index int) *Node {
	current := ll.head
	for i := 0; i < index; i++ {
		// if out of bounds
		if current == nil {
			return nil
		}
		current = current.next
	}
	return current
}

func (ll *LinkedList) remove(index int) *Node {
	current := ll.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return nil
		}
		current = current.next
	}

	// edge cases
	if index == 0 {
		return ll.popFirst()
	} else if current.next.next == nil {
		return ll.popLast()
	// the node we want to remove is current.next
	// because the loop is only until index-1
	} else {
		defer func() {
			current.next = current.next.next
		}()
		return current.next
	}
}

func (ll *LinkedList) insert(index int, n Node) {
	// edge cases
	if index == 0 {
		ll.prepend(n)
	} else if index >= ll.size() {
		ll.append(n)
	// inserting in the middle of list
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		// we want to put new node in current.next
		// because the loop is only until index-1
		n.next = current.next
		current.next = &n
	}
}

func (ll LinkedList) print() {
	var current *Node
	for current = ll.head; current != nil; current = current.next {
		fmt.Printf("%v <-> ", current.data)
	}
	fmt.Printf("%v\n", current)
}

func (ll LinkedList) size() int {
	len := 0
	for current := ll.head; current != nil; current = current.next {
		len++
	}
	return len
}

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