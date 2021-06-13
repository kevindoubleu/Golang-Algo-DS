package main

import "fmt"

func main() {
	var list LinkedList
	list.print()

	// appends
	list.append(Node{data: 1})
	list.append(Node{data: 2})
	list.append(Node{data: 3})
	list.print()
	fmt.Println("size is", list.size())

	// prepends
	list.prepend(Node{data: 4})
	list.prepend(Node{data: 5})
	list.prepend(Node{data: 6})
	list.print()
	fmt.Println("size is", list.size())
	fmt.Println("")

	// gethead gettail wrappers
	list.first().data = 0
	list.last().data = 99
	fmt.Println("first node is now", list.first())
	fmt.Println("last node is now", list.last())
	list.print()
	fmt.Println("")

	// pops
	fmt.Println("popfirst popped", list.popFirst())
	fmt.Println("poplast popped", list.popLast())
	list.print()
	fmt.Println("")

	// index based operations
	index := 2
	fmt.Println("item at index", index, "is", list.get(index))
	fmt.Println("remove item at index", index, "which is", list.remove(index))
	item := Node{data: 69}
	fmt.Println("insert item", item,"at index", index)
	list.insert(index, item)
	item = Node{data: 999}
	index = 30
	fmt.Println("insert item", item,"at index", index)
	list.insert(index, item)
	list.print()
	fmt.Println("")

	// crude iterator
	// haven't understood go's iterator idioms yet
	i := list.getIterator()
	for n := i(); n != nil; n = i() {
		fmt.Println("modifying node", n)
		n.data += 10
	}
	list.print()
	fmt.Println("")

	list.clear()
	fmt.Println("cleared list")
	list.print()
	fmt.Println("remove item at index", 1, "which is", list.remove(1))
	fmt.Println("remove item at index", 2, "which is", list.remove(2))
	list.print()
}