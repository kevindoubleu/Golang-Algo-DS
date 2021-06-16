package main

import "fmt"

func (t *BST) preorder(f func(n *Node)) {
	fmt.Println("running preorder traversal")
	t.root._preorder(f)
}
func (n *Node) _preorder(f func(n *Node)) {
	if n != nil {
		f(n)
		n.left._preorder(f)
		n.right._preorder(f)
	}
}

func (t *BST) inorder(f func(n *Node)) {
	fmt.Println("running inorder traversal")
	t.root._inorder(f)
}
func (n *Node) _inorder(f func(n *Node)) {
	if n != nil {
		n.left._inorder(f)
		f(n)
		n.right._inorder(f)
	}
}

func (t *BST) postorder(f func(n *Node)) {
	fmt.Println("running postorder traversal")
	t.root._postorder(f)
}
func (n *Node) _postorder(f func(n *Node)) {
	if n != nil {
		n.left._postorder(f)
		n.right._postorder(f)
		f(n)
	}
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

	bst1.preorder(func(n *Node) {fmt.Println(n.data)})

	bst1.inorder(func(n *Node) {
		n.data += 10
		fmt.Println(n.data)
	})

	bst1.postorder(func(n *Node) {
		n.data -= 10
		fmt.Println(n.data)
	})

	bst1.print()
}