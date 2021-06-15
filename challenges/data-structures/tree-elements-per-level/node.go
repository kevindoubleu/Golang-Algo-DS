package main

import (
	"fmt"
)

type Node struct {
	data int
	children []*Node
}

// add node as a child of other node
func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}

func (n *Node) printDF() {
	n._printDF("")
}
func (n *Node) _printDF(space string) {
	if n != nil {
		fmt.Printf("%s%v\n", space, n)
	}
	
	for _, child := range n.children {
		child._printDF(space + "    ")
	}
}