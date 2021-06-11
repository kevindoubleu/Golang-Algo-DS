package main

import "fmt"

type Node struct {
	// the payload of the queue
	name string
	order string
}

// this queue will only hold items of type Node
type Queue []Node

// enqueue
func (q *Queue) push(n *Node) {
	*q = append(*q, *n)
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

// debug
func (q Queue) print(prefix ...interface{}) {
	fmt.Print(prefix...)
	fmt.Println(q)
	fmt.Println("------")
}

func main() {
	var line Queue

	line.push(&Node{"andy", "burger"})
	line.push(&Node{"brody", "fries"})
	line.push(&Node{"cindy", "ice cream"})
	line.print("after insertions: ")

	popped := line.pop()
	fmt.Println("popped", popped)
	line.print("after deletion: ")

	peeked := line.peek()
	fmt.Println("peeked at", peeked)
	line.print("after peek: ")

	looked := line.lookup("cindy")
	fmt.Println("looked up", looked)
	line.print("after lookup: ")
}