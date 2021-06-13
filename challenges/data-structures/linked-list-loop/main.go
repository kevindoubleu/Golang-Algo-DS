package main

import "fmt"

// we keep track of the memory addresses we've visited
// if we see the same memory again then that means we're looping
func (ll *LinkedList) isLooped() bool {
	// use map to make lookup O(1)
	addresses := make(map[*Node]bool)

	i := ll.getIterator()
	for n := i(); n != nil; n = i() {
		// fmt.Printf("at %p is %v\n", n, n)
		if !addresses[n] {
			addresses[n] = true
		} else {
			// we've already seen this address so it's a loop
			return true
		}
	}

	return false
}

// we use a slow and fast pointer
// if it's looped the slow and fast pointer will point at the same thing
// at some point while they are inside the loop
// if it's not looped then fast will hit a null pointer
func (ll *LinkedList) isLooped2() bool {
	// edge cases
	// empty list
	if ll.head == nil {
		return false
	// only 1 node
	} else if ll.head.next == nil {
		// does it point at itself?
		return ll.head.next == ll.head
	}

	slow, fast := ll.first(), ll.first().next;
	for slow != fast {
		if fast.next != nil && fast.next.next != nil {
			fast = fast.next.next
			slow = slow.next
		} else {
			// hit a null
			return false
		}
	}
	// at this point slow == fast
	return true
}

// shorter version of isLooped2
func (ll *LinkedList) isLooped3() bool {
	if ll.head == nil {
		return false
	}

	slow, fast := ll.first(), ll.first();
	for fast.next != nil && fast.next.next != nil {
		fast = fast.next.next
		slow = slow.next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	var l LinkedList
	fmt.Println("empty list isLooped: ", l.isLooped())
	fmt.Println("empty list isLooped: ", l.isLooped2())
	fmt.Println("empty list isLooped: ", l.isLooped3())

	l.insert(0, Node{data: 8})
	l.insert(0, Node{data: 7})
	l.insert(0, Node{data: 6})
	l.insert(0, Node{data: 5})
	l.insert(0, Node{data: 4})
	l.insert(0, Node{data: 3})
	l.insert(0, Node{data: 2})
	l.insert(0, Node{data: 1})
	l.print()
	fmt.Println("regular list isLooped: ", l.isLooped())
	fmt.Println("regular list isLooped: ", l.isLooped2())
	fmt.Println("regular list isLooped: ", l.isLooped3())

	n1 := Node{data: 1}
	n2 := Node{data: 2}
	n3 := Node{data: 3}
	n4 := Node{data: 4}
	
	// regular looped
	var looped LinkedList
	looped.head = &n1
	n1.next = &n2
	n2.next = &n3
	n3.next = &n4
	n4.next = &n2
	fmt.Println("loop1 isLooped: ", looped.isLooped())
	fmt.Println("loop1 isLooped: ", looped.isLooped2())
	fmt.Println("loop1 isLooped: ", looped.isLooped3())

	// 1 node pointing at itself
	var looped2 LinkedList
	looped2.head = &n1
	n1.next = &n1
	fmt.Println("loop2 isLooped: ", looped2.isLooped())
	fmt.Println("loop2 isLooped: ", looped2.isLooped2())
	fmt.Println("loop2 isLooped: ", looped2.isLooped3())

	// 2 nodes pointing at each other
	var looped3 LinkedList
	looped3.head = &n1
	n1.next = &n2
	n2.next = &n1
	fmt.Println("loop3 isLooped: ", looped3.isLooped())
	fmt.Println("loop3 isLooped: ", looped3.isLooped2())
	fmt.Println("loop3 isLooped: ", looped3.isLooped3())
}