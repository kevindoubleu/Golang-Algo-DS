package main

import "fmt"

type Node struct {
	// example node representing dishes to be washed
	meal string
	count int
}

type Stack []Node

func (s *Stack) push(n Node) {
	*s = append(*s, n)
}

func (s *Stack) pop() *Node {
	if len(*s) > 0 {
		tmp := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return &tmp
	}

	return nil
}

func (s Stack) peek() *Node {
	if len(s) > 0 {
		return s.pop()
	}

	return nil
}

func (s Stack) lookup(meal string) *Node {
	for s.peek() != nil {
		if s.peek().meal == meal {
			return s.pop()
		}
		s.pop()
	}

	return nil
}

func main() {
	var dishes Stack

	dishes.push(Node{"breakfast", 2})
	dishes.push(Node{"snack", 1})
	dishes.push(Node{"lunch", 3})
	dishes.push(Node{"brunch", 2})
	dishes.push(Node{"dinner", 3})
	fmt.Printf("after insertion: %v\n\n", dishes)

	popped := dishes.pop()
	fmt.Println("popped", popped)
	fmt.Printf("after pop: %v\n\n", dishes)

	peeked := dishes.peek()
	fmt.Println("peeked at", peeked)
	fmt.Printf("after peek: %v\n\n", dishes)

	looked := dishes.lookup("breakfast")
	fmt.Println("looked up", looked)
	fmt.Printf("after lookup: %v\n\n", dishes)

	// tests for empty stack
	for i := 0; i < 5; i++ {
		popped = dishes.pop()
		fmt.Println("popped", popped)
	}
	peeked = dishes.peek()
	fmt.Println("peeked at", peeked)

	fmt.Printf("after pops: %v\n\n", dishes)
}