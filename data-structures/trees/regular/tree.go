package main

import (
	"container/list"
)

type Tree struct {
	root *Node
}

func (t *Tree) traverseBF(f func(n *Node)) {
	// edge case empty tree
	if t.root == nil {
		return
	}

	queue := list.New().Init()
	queue.PushBack(t.root)

	for queue.Len() > 0 {
		// we're going to process the head
		head := queue.Front()

		// enqueue all the head's children
		for _, child := range head.Value.(*Node).children {
			queue.PushBack(child)
		}

		// process the head
		f(head.Value.(*Node))

		// we're done, remove the head
		queue.Remove(head)
	}
}

func (t *Tree) traverseDF(f func(n *Node)) {
	// edge case empty tree
	if t.root == nil {
		return
	}

	// we'll use a doubly link list as a stack
	// with the list head as the top of stack
	stack := list.New().Init()
	stack.PushFront(t.root)

	for stack.Len() > 0 {
		// we're going to process the top of stack
		top := stack.Front()

		// push all the top's children to top of stack
		// first we convert the children into a list
		// then we combine that list
		// we do it this way to maintain the order of pushed children
		childrenList := list.New().Init()
		for _, child := range top.Value.(*Node).children {
			childrenList.PushBack(child)
		}
		stack.PushFrontList(childrenList)

		// process the top
		f(top.Value.(*Node))

		// we're done, remove the top
		stack.Remove(top)
	}
}

// recursive depth first traversal
func (t *Tree) traverseDFrec(f func(n *Node)) {
	if t.root != nil {
		f(t.root)
	}
	
	for _, child := range t.root.children {
		subTree := Tree{root: child}
		subTree.traverseDFrec(f)
	}
}