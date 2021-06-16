package main

import "fmt"

func (t *BST) contains(data int) bool {
	return t.root._contains(data)
}
func (n *Node) _contains(data int) bool {
	if n != nil {
		// we found it
		if n.data == data {
			return true
		// bigger go right
		} else if data > n.data {
			return n.right._contains(data)
		// smaller go left
		} else if data < n.data {
			return n.left._contains(data)
		}
	}

	// not found
	return false
}

func main() {
	var bst1 BST
	bst1.insert(10)
	bst1.insert(0)
	bst1.insert(12)
	bst1.insert(-1)
	bst1.insert(4)
	bst1.insert(11)
	bst1.insert(20)
	bst1.insert(-20)
	bst1.print()

	fmt.Println(bst1.contains(10))
	fmt.Println(bst1.contains(0))
	fmt.Println(bst1.contains(20))
	fmt.Println(bst1.contains(21))
	fmt.Println(bst1.contains(200))
}