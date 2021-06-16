package main

func (ll *LinkedList) reverse() {
	var prev, curr, next *Node
	curr = ll.head
	for curr != nil {
		// get the next item and put it in next
		next = curr.next

		// actually reverse the current node
		curr.next = prev
		
		// move prev forward
		prev = curr
		// we can still move curr "forward" bcs we already have it in next
		curr = next
	}

	// curr is now nil so prev is at tail
	// and tail is the new head
	ll.head = prev
}

func main() {
	var list LinkedList
	list.prepend(Node{data: 1})
	list.prepend(Node{data: 2})
	list.prepend(Node{data: 3})
	list.prepend(Node{data: 4})
	list.prepend(Node{data: 5})
	list.prepend(Node{data: 6})
	list.prepend(Node{data: 7})
	list.prepend(Node{data: 8})
	list.prepend(Node{data: 9})
	list.prepend(Node{data: 10})
	list.print()

	list.reverse()
	list.print()
}