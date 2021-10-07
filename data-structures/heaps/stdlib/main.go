package main

import (
	"container/heap"
	"fmt"
)

type MyHeap []int

func (a MyHeap) Len() int           { return len(a) }
func (a MyHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// just change this function to change between maxheap and minheap
func (a MyHeap) Less(i, j int) bool { return a[i] < a[j] }

func (h *MyHeap) Push(x interface{}) {
	// as in the documentation says
	// add x as element Len()
	*h = append(*h, x.(int))
}
func (h *MyHeap) Pop() interface{} {
	// as in the documentation says too
	// remove and return element Len() - 1.

	retval := (*h)[len(*h)-1]

	// the actual removal from the heap
	*h = (*h)[0 : len(*h)-1]

	return retval
}

func main() {
	h1 := MyHeap{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17}

	heap.Init(&h1)
	fmt.Println(h1)

	for range h1 {
		fmt.Println("popping", heap.Pop(&h1))
		fmt.Println(h1)
	}
}