package main

type Node struct {
	children map[string]*Node
	isEnd bool
}

func newNode() *Node {
	n := Node{}
	n.children = make(map[string]*Node)
	return &n
}

func (n *Node) addChild(value string) {
	child := newNode()
	n.children[value] = child
}
