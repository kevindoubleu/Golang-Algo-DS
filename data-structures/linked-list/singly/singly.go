package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (ll *LinkedList) append(n Node) {
	if ll.head == nil {
		ll.head = &n
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = &n
	}
}

func (ll *LinkedList) prepend(n Node) {
	n.next = ll.head
	ll.head = &n
}

// returns pointer so we can edit the nodes
func (ll *LinkedList) first() *Node {
	return ll.head
}

func (ll *LinkedList) last() *Node {
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	return current
}

func (ll *LinkedList) popFirst() *Node {
	tmp := ll.head
	if ll.head != nil {
		ll.head = ll.head.next
	}
	return tmp
}

func (ll *LinkedList) popLast() *Node {
	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	defer func() {
		current.next = nil
	}()
	return current.next
}

func (ll *LinkedList) clear() {
	// entire list will be garbage collected
	ll.head = nil
}

func (ll *LinkedList) get(index int) *Node {
	current := ll.head
	for i := 0; i < index; i++ {
		// if out of bounds
		if current == nil {
			return nil
		}
		current = current.next
	}
	return current
}

func (ll *LinkedList) remove(index int) *Node {
	current := ll.head
	for i := 0; i < index-1; i++ {
		if current == nil {
			return nil
		}
		current = current.next
	}

	// edge cases
	if index == 0 {
		return ll.popFirst()
	} else if current.next.next == nil {
		return ll.popLast()
	// the node we want to remove is current.next
	// because the loop is only until index-1
	} else {
		defer func() {
			current.next = current.next.next
		}()
		return current.next
	}
}

func (ll *LinkedList) insert(index int, n Node) {
	// edge cases
	if index == 0 {
		ll.prepend(n)
	} else if index >= ll.size() {
		ll.append(n)
	// inserting in the middle of list
	} else {
		current := ll.head
		for i := 0; i < index-1; i++ {
			current = current.next
		}
		// we want to put new node in current.next
		// because the loop is only until index-1
		n.next = current.next
		current.next = &n
	}
}

func (ll LinkedList) print() {
	var current *Node
	for current = ll.head; current != nil; current = current.next {
		fmt.Printf("%v -> ", current.data)
	}
	fmt.Printf("%v\n", current)
}

func (ll LinkedList) size() int {
	len := 0
	for current := ll.head; current != nil; current = current.next {
		len++
	}
	return len
}

// for use as a loop iterator
func (ll *LinkedList) getIterator() func() *Node {
	current := ll.head
	return func() *Node {
		if current != nil {
			defer func() {
				current = current.next
			}()
			return current
		} else {
			return nil
		}
	}
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