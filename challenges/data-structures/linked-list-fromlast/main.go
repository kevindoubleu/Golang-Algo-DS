package main

import "fmt"

// we'll use 2 pointers again, left and right
// we "increment" right x times so that
// the distance between left and right is x
// then we move them together until right reaches the end of list
// and left will be pointing at index x from last
func (ll *LinkedList) fromLast(index int) *Node {
	// edge cases
	// empty list
	if ll.head == nil {
		return nil
	}
	// negative indexes arent accepted
	if index < 0 {
		return nil
	}

	left, right := ll.first(), ll.first()
	for i := 0; i < index; i++ {
		if right.next != nil {
			right = right.next
		} else {
			// if out of bounds
			return nil
		}
	}
	
	for right.next != nil {
		left = left.next
		right = right.next
	}

	return left
}

func main() {
	var l LinkedList

	// empty list test
	index := 0
	fmt.Println("index", index, "fromlast is", l.fromLast(index))

	l.insert(0, Node{data: 8})
	l.insert(0, Node{data: 7})
	l.insert(0, Node{data: 6})
	l.insert(0, Node{data: 5})
	l.insert(0, Node{data: 4})
	l.insert(0, Node{data: 3})
	l.insert(0, Node{data: 2})
	l.insert(0, Node{data: 1})
	l.print()

	// normal
	index = 0
	fmt.Println("index", index, "fromlast is", l.fromLast(index))
	index = 1
	fmt.Println("index", index, "fromlast is", l.fromLast(index))
	index = 2
	fmt.Println("index", index, "fromlast is", l.fromLast(index))
	index = 3
	fmt.Println("index", index, "fromlast is", l.fromLast(index))

	// negatives
	index = -100
	fmt.Println("index", index, "fromlast is", l.fromLast(index))

	// out of bounds
	index = 100
	fmt.Println("index", index, "fromlast is", l.fromLast(index))
	index = 8
	fmt.Println("index", index, "fromlast is", l.fromLast(index))
}