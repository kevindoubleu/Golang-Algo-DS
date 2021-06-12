package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) append(n Node) {
	// list empty. head and tail is nil
	if ll.head == nil && ll.tail == nil {
		ll.head = &n
		ll.tail = &n
	// list has 1 node, head and tail pointing to same node
	// list has >1, normal
	} else {
		n.prev = ll.tail
		ll.tail.next = &n

		ll.tail = &n
	}
}

func (ll *LinkedList) prepend(n Node) {
	if ll.head == nil && ll.tail == nil {
		ll.head = &n
		ll.tail = &n
	} else {
		n.next = ll.head
		ll.head.prev = &n

		ll.head = &n
	}
}

// returns pointer so we can edit the nodes
func (ll *LinkedList) first() *Node {
	return ll.head
}

func (ll *LinkedList) last() *Node {
	return ll.tail
}

func (ll *LinkedList) popFirst() *Node {
	tmp := ll.head
	if ll.head == ll.tail {
		ll.head, ll.tail = nil, nil
	} else {
		ll.head = ll.head.next
		ll.head.prev = nil
	}
	return tmp
}

func (ll *LinkedList) popLast() *Node {
	tmp := ll.tail
	if ll.head == ll.tail {
		ll.head, ll.tail = nil, nil
	} else {
		ll.tail = ll.tail.prev
		ll.tail.next = nil
	}
	return tmp
}

func (ll *LinkedList) clear() {
	// in golang, circular references can be garbage collected
	ll.head = nil
	ll.tail = nil
}

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
}