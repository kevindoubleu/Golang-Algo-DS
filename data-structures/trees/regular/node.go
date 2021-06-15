package main

type Node struct {
	data int
	children []*Node
}

// add node as a child of other node
func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}

// remove a node's child node with matching data
func (n *Node) removeChild(data int) {
	for i, child := range n.children {
		if child.data == data {
			// deletion
			copy(n.children[i:], n.children[i+1:])
			n.children = n.children[:len(n.children)-1]
			break
		}
	}
}