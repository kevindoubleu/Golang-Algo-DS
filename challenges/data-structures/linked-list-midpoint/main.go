package main

import (
	"fmt"
)

// use 2 pointers, 1 slow, 1 fast.
// fast will iterate 2x as fast as slow.
// this way when fast reaches tail,
// slow will be halfway from tail, which is the middle.
func (ll *LinkedList) midpoint() *Node {
	// edge case empty list
	if ll.head == nil {
		return nil
	}
	
	slow, fast := ll.first(), ll.first()
	for {
		if fast.next != nil {
			fast = fast.next
		} else {
			break
		}
		// we move slow on the second "jump"
		// so on even length lists, slow will not advance
		// and therefore will take the "left" middle element
		if fast.next != nil {
			slow = slow.next
			fast = fast.next
		}
	}
	return slow
}

// same thing but shorter
func (ll *LinkedList) midpoint2() *Node {
	if ll.head == nil {
		return nil
	}
	
	slow, fast := ll.first(), ll.first()
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}

func main() {
	var l LinkedList
	// empty list
	l.print()
	fmt.Println("midpoint is", l.midpoint())

	// 1 node
	l.insert(0, Node{data: 4})
	l.print()
	fmt.Println("midpoint is", l.midpoint())

	// 2 nodes
	l.insert(0, Node{data: 3})
	l.print()
	fmt.Println("midpoint is", l.midpoint())

	// populating the list for normal tests
	l.insert(0, Node{data: 2})
	l.insert(0, Node{data: 1})
	l.insert(l.size(), Node{data: 5})
	l.insert(l.size(), Node{data: 6})
	l.insert(l.size(), Node{data: 7})
	l.insert(l.size(), Node{data: 8})
	l.remove(0)
	l.print()
	
	fmt.Println("midpoint with length odd")
	fmt.Println("midpoint is", l.midpoint())
	fmt.Println("midpoint is", l.midpoint2())

	l.insert(3, Node{data:123})
	l.print()
	fmt.Println("midpoint with length even")
	fmt.Println("midpoint is", l.midpoint())
	fmt.Println("midpoint is", l.midpoint2())
}