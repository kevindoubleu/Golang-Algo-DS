package main

import "fmt"

type Node struct {
	data int
	left *Node
	right *Node
}

func (n *Node) insert(data int) *Node {
	// if empty spot
	if n == nil {
		return &Node{data: data}

	// bigger go right
	} else if data > n.data {
		n.right = n.right.insert(data)

	// smaller go left
	} else if data < n.data {
		n.left = n.left.insert(data)
	}
	
	// BSTs dont accept duplicates
	return n
}

// wrapper func for recursive printer function _print
func (n *Node) print() {
	fmt.Println("")
	n._print("")
	fmt.Println("")
}
func (n *Node) _print(space string) {
	if n == nil {
		return
	}
	space += "   "
	
	n.right._print(space)
	fmt.Printf("%s%d\n", space, n.data)
	n.left._print(space)
}

func (n *Node) preorder() {
	if n == nil {
		return
	}
	fmt.Print(n.data, " ")
	n.left.preorder()
	n.right.preorder()
}

func (n *Node) inorder() {
	if n == nil {
		return
	}
	n.left.inorder()
	fmt.Print(n.data, " ")
	n.right.inorder()
}

func (n *Node) postorder() {
	if n == nil {
		return
	}
	n.left.postorder()
	n.right.postorder()
	fmt.Print(n.data, " ")
}

// wrapper class for handling BST initialization
type BST struct {
	root *Node
}

func (t *BST) insert(data int) {
	if t.root == nil {
		t.root = &Node{data: data}
	} else {
		t.root.insert(data)
	}
}

// wrapper func for BST struct
func (t *BST) print() {
	fmt.Println("")
	t.root._print("")
	fmt.Println("")
}