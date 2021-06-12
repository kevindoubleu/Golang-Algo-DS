package main

import "fmt"

// there are 2 ways to do it
// 1. using top of stack as the "front" of queue
//     push O(n), pop O(1), peek O(1), lookup O(n)
// 2. using bottom of stack as the "front" of queue
//     push O(1), pop O(n), peek O(n), lookup O(n)
// i'm using top of stack as the front because it's faster

func (s *Stack) pushQ(n Node) {
	// move all items to another stack, the items are now inverted
	var tempStack Stack
	for s.peek() != nil {
		tempStack.push(*s.pop())
	}
	// push into the original stack
	s.push(n)
	// move back all the things to the original stack
	for tempStack.peek() != nil {
		s.push(*tempStack.pop())
	}
}

// pop is the same
// because top of stack is also top of queue
func (s *Stack) popQ() *Node {
	return s.pop()
}

// peek is the same
// because top of stack is also top of queue
func (s *Stack) peekQ() *Node {
	return s.peek()
}

// lookup is the same
// doesn't matter where the top of queue is
func (s *Stack) lookupQ(meal string) *Node {
	return s.lookup(meal)
}

func main() {
	var fakeStack Stack

	fakeStack.pushQ(Node{"breakfast", 1})
	fmt.Printf("%v\n", fakeStack)
	fakeStack.pushQ(Node{"lunch", 2})
	fmt.Printf("%v\n", fakeStack)
	fakeStack.pushQ(Node{"dinner", 3})
	fmt.Printf("%v\n", fakeStack)
	fakeStack.pushQ(Node{"midnight", 4})
	fmt.Printf("after insertions: %v\n\n", fakeStack)
	
	popped := fakeStack.popQ()
	fmt.Println("popped", popped)
	fmt.Printf("after deletion: %v\n\n", fakeStack)

	peeked := fakeStack.peekQ()
	fmt.Println("peeked at", peeked)
	fmt.Printf("after peek: %v\n\n", fakeStack)

	looked := fakeStack.lookupQ("dinner")
	fmt.Println("looked up", looked)
	fmt.Printf("after lookup: %v\n\n", fakeStack)

	for i := 0; i < 4; i++ {
		popped := fakeStack.popQ()
		fmt.Println("popped", popped)
		fmt.Printf("after deletion: %v\n\n", fakeStack)
	}
}