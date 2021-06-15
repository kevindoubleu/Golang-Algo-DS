package main

import "fmt"

func main() {
	// indentation for illustration
	n1 := Node{data:1}
		n2 := Node{data:2}
		n3 := Node{data:3}
			n6 := Node{data:31}
			n7 := Node{data:32}
		n4 := Node{data:4}
		n5 := Node{data:5}

	n1.addChild(&n2)
	n1.addChild(&n3)
	n3.addChild(&n6)
	n3.addChild(&n7)
	n1.addChild(&n4)
	n1.addChild(&n5)
		
	t := Tree{root: &n1}
	n1.removeChild(5)

	// test for breadth print
	fmt.Println("print breadth first")
	t.traverseBF(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")

	// test execute function on every node
	fmt.Println("add 10 breadth first")
	t.traverseBF(func(n *Node) {
		(*n).data += 10
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")

	// test for depth first
	fmt.Println("print depth first")
	t.traverseDF(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")
	fmt.Println("print depth first recursive")
	t.traverseDFrec(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")

	// test execute function on every node
	fmt.Println("minus 10 depth first")
	t.traverseDF(func(n *Node) {
		(*n).data -= 10
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")

	var emptyTree Tree
	// test empty tree breadth first
	fmt.Println("print breadth first on empty tree")
	emptyTree.traverseBF(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")

	// test empty tree depth first
	fmt.Println("print depth first on empty tree")
	emptyTree.traverseBF(func(n *Node) {
		fmt.Printf("%v\n", n)
	})
	fmt.Println("")
}