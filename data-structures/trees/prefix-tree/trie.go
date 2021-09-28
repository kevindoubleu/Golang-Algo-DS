package main

import (
	"container/list"
	"fmt"
)

type Trie struct {
	root *Node
}

func newTrie() *Trie {
	return &Trie{root: newNode("")}
}

// breadth first
func (t *Trie) print() {
	if t == nil { return }

	queue := list.New().Init()
	queue.PushBack(t.root)

	for queue.Len() > 0 {
		head := queue.Front()

		// process and pop front of queue
		fmt.Println(head.Value)
		queue.Remove(head)

		// queue the children
		for _, child := range head.Value.(*Node).children {
			queue.PushBack(child)
		}

	}
}

func (t *Trie) insert(str string) {
	current := t.root
	for _, c := range str {
		char := string(c)
		if current.children[char] == nil {
			current.addChild(char)
		}
		current = current.children[char]
	}
}

func (t *Trie) search(str string) bool {
	current := t.root
	for _, c := range str {
		char := string(c)
		if current.children[char] == nil {
			return false
		}
		current = current.children[char]
	}
	return true
}
