package main

import (
	"container/list"
	"fmt"
)

func (n *Node) elementsCount() []int {
	// edge case nil node reciever
	if n == nil {
		return []int{}
	}

	// make 2 queues
	// 1 for the current level
	// 1 for the next element
	// for every item in the currentlevel array
	// 		push their children into the next element array
	// 		get the arr len and append it to result arr
	// now that we're done with curr level
	// we use next lvl as curr lvl and clear the next lvl arr
	// do this while there are items in the array
	
	currLv := list.New().Init()
	nextLv := list.New().Init()
	
	currLv.PushBack(n)
	result := []int{1}

	for {
		// for every item in the currentlevel array
		for currLv.Len() > 0 {
			curr := currLv.Front()
	
			// push their children into the next element array
			for _, child := range curr.Value.(*Node).children {
				// fmt.Println("adding", child, "to level", len(result))
				nextLv.PushBack(child)
			}

			currLv.Remove(curr)
		}

		// get the arr len and append it to result arr
		// do this while there are items in the array
		if nextLv.Len() > 0 {
			result = append(result, nextLv.Len())
		} else {
			break
		}

		// now that we're done with curr level
		// we use next lvl as curr lvl and clear the next lvl arr
		currLv.PushBackList(nextLv)
		nextLv.Init()
	}

	return result
}

func main() {
	n11 := Node{data: 5}
	n21 := Node{data: 1}
	n22 := Node{data: 3}
	n23 := Node{data: 9}
	n31 := Node{data: 2}
	n32 := Node{data: 6}
	n33 := Node{data: 4}
	n34 := Node{data: 8}

	n11.addChild(&n21)
	n11.addChild(&n22)
	n11.addChild(&n23)
	n21.addChild(&n31)
	n21.addChild(&n32)
	n22.addChild(&n33)
	n23.addChild(&n34)

	n11.printDF()
	fmt.Println("the element counts per level are:", n11.elementsCount())
	fmt.Println("")

	n21.printDF()
	fmt.Println("the element counts per level are:", n21.elementsCount())
	fmt.Println("")

	n22.printDF()
	fmt.Println("the element counts per level are:", n22.elementsCount())
	fmt.Println("")

	n31.printDF()
	fmt.Println("the element counts per level are:", n31.elementsCount())
	fmt.Println("")
}