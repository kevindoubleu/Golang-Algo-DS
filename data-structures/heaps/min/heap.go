package main

import (
	"fmt"
)

func parentOf(idx int) int {
	var parent int
	// odd is left child
	if idx % 2 == 1 {
		parent = (idx-1) / 2
	} else {
		// even is right child
		parent = (idx-2) / 2
	}

	// index out of bounds safety
	if parent < 0 {
		return 0
	}
	return parent
}

func leftChild(idx int) int {
	return idx * 2 + 1
}

func rightChild(idx int) int {
	return idx * 2 + 2
}

type MinHeap []int

// heapify from newest leaf to root
func (h MinHeap) heapifyUp() {
	last := len(h)-1
	parent := parentOf(last)

	for h[last] < h[parent] {
		// swap if current is lesser than parent (minheap)
		h[last], h[parent] = h[parent], h[last]
		// set parent as current for further swaps
		last = parent
		parent = parentOf(last)
	}
}

func (h *MinHeap) insert(value int) {
	*h = append(*h, value)
	h.heapifyUp()
}

func (h *MinHeap) insertMany(values []int) {
	for _, v := range values {
		h.insert(v)
	}
}


// heapify from root to leaf
func (h MinHeap) heapifyDown() {
	heapLen := len(h)

	if heapLen == 0 {
		return
	}

	curr := 0
	for rightChild(curr) < heapLen {
		// find smaller child
		var childIdx int
		if h[leftChild(curr)] < h[rightChild(curr)] {
			childIdx = leftChild(curr)
		} else {
			childIdx = rightChild(curr)
		}

		// swap if child is smaller
		if h[childIdx] < h[curr] {
			h[childIdx], h[curr] = h[curr], h[childIdx]
		}

		// traverse down to child
		curr = childIdx
	}
}

func (h *MinHeap) extractMax() int {
	if len(*h) == 0 {
		return 0
	}

	maxValue := (*h)[0]

	// replace root with last leaf then heapify
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	h.heapifyDown()

	return maxValue
}

func main() {
	h1 := MinHeap{}
	h1.insertMany([]int{4, 10, 3, 5, 1})
	fmt.Println(h1)

	h2 := MinHeap{}
	h2.insertMany([]int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17})
	fmt.Println(h2)

	h2.extractMax()
	fmt.Println(h2)
}