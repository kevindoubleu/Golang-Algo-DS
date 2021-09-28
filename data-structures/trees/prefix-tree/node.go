package main

type Node struct {
	children map[string]*Node
	value string
}

func newNode(value string) *Node {
	n := Node{value: value}
	n.children = make(map[string]*Node)
	return &n
}

func (n *Node) addChild(value string) {
	child := newNode(value)
	n.children[child.value] = child
}
