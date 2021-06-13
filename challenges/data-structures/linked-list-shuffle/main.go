package main

import (
	"math/rand"
	"time"
)

func main() {
	var list LinkedList
	list.append(Node{data: 1})
	list.append(Node{data: 2})
	list.append(Node{data: 3})
	list.append(Node{data: 4})
	list.append(Node{data: 5})
	list.append(Node{data: 6})
	list.print()
	
	// shuffle procedure
	// 1. get slice of values from linked list
	// 2. shuffle the slice
	// 3. assign the shuffled slice values into the linked list
	
	// 1. get slice of values from linked list
	var vals []int
	i := list.getIterator()
	for n := i(); n != nil; n = i() {
		vals = append(vals, n.data)
	}
	
	// 2. shuffle the slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(vals), func(i, j int) {
		vals[i], vals[j] = vals[j], vals[i]
	})

	// 3. assign the shuffled slice values into the linked list
	i = list.getIterator()
	for _, v := range vals {
		n := i()
		n.data = v
	}

	list.print()
}