package main

import (
	"errors"
	"fmt"
)

type Node struct {
	// the payload of the queue
	name string
	order string
}

// this queue will only hold items of type Node
type Queue []Node

// enqueue
func (q *Queue) push(n Node) {
	*q = append(*q, n)
}

// dequeue,
// returns the dequeued node and error,
// error triggered when popping an empty queue
func (q *Queue) pop() (Node, error) {
	if len(*q) > 0 {
		tmp := (*q)[0]
		*q = (*q)[1:]
		return tmp, nil
	}

	return Node{}, errors.New("queue pop error: empty queue")
}

// debug
func (q Queue) print(prefix ...interface{}) {
	fmt.Print(prefix...)
	fmt.Println(q)
	fmt.Println("------")
}

func main() {
	var line Queue
	var popped Node
	var err error

	line.push(Node{"andy", "burger"})
	line.print("after insert: ")
	
	line.push(Node{"brody", "fries"})
	line.print("after insert: ")
	
	line.push(Node{"cindy", "ice cream"})
	line.print("after insert: ")

	for i := 0; i < 4; i++ {
		popped, err = line.pop()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("popped", popped)
		line.print("after pop: ")
	}
}