package main

func main() {
	// with BST wrapper struct
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

	// without, notice the first node is inserted differently
	bst2 := Node{data: 10}
	bst2.insert(0)
	bst2.insert(12)
	bst2.insert(-1)
	bst2.insert(4)
	bst2.insert(11)
	bst2.insert(20)
	bst2.insert(-20)
	bst2.print()
}