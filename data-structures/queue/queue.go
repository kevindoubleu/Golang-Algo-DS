package main

import "fmt"

type Node struct {
	// example node representing people ordering foods in a line
	name string
	order string
}

type Queue []Node

// enqueue
func (q *Queue) push(n Node) {
	*q = append(*q, n)
}

// dequeue,
// returns pointer to dequeued node or
// nil if queue is empty
func (q *Queue) pop() *Node {
	if len(*q) > 0 {
		tmp := (*q)[0]
		*q = (*q)[1:]
		return &tmp
	}

	return nil
}

// returns pointer to first node or
// nil if queue is empty
func (q Queue) peek() *Node {
	if len(q) > 0 {
		return q.pop()
	}

	return nil
}

// returns pointer to node with node's name name or
// nil if not found
func (q Queue) lookup(name string) *Node {
	for q.peek() != nil {
		if q.peek().name == name {
			return q.pop()
		}
		q.pop()
	}

	return nil
}

func main() {
	var line Queue

	line.push(Node{"andy", "burger"})
	line.push(Node{"brody", "fries"})
	line.push(Node{"cindy", "ice cream"})
	fmt.Printf("after insertions: %v\n\n", line)

	popped := line.pop()
	fmt.Println("popped", popped)
	fmt.Printf("after deletion: %v\n\n", line)

	peeked := line.peek()
	fmt.Println("peeked at", peeked)
	fmt.Printf("after peek: %v\n\n", line)

	looked := line.lookup("cindy")
	fmt.Println("looked up", looked)
	fmt.Printf("after lookup: %v\n\n", line)
}