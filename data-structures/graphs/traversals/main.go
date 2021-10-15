package main

import (
	"container/list"
	"fmt"
)

func (g Graph) BFS(rootVertex int) {
	fmt.Println("running breadth first search")
	// edge case vertex doesnt exist
	if g.adjList[rootVertex] == nil {
		fmt.Println("vertex not found")
		return
	}

	queue := list.New().Init()
	visited := make(map[int]bool)

	queue.PushBack(rootVertex)
	for queue.Len() > 0 {
		currVertex := queue.Remove(queue.Front())

		fmt.Println("processing vertex", currVertex)
		visited[currVertex.(int)] = true
		// fmt.Println("visited", visited)

		for neighbour := range g.adjList[currVertex.(int)] {
			if !visited[neighbour] {
				queue.PushBack(neighbour)
			}
		}
	}
}

func (g Graph) DFS(rootVertex int) {
	fmt.Println("running depth first search (postorder)")
	// edge case vertex doesnt exist
	if g.adjList[rootVertex] == nil {
		fmt.Println("vertex not found")
		return
	}

	stack := []int{rootVertex}
	visited := make(map[int]bool)

	for len(stack) > 0 {
		currVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println("processing vertex", currVertex)
		visited[currVertex] = true
		
		for neighbour := range g.adjList[currVertex] {
			if !visited[neighbour] {
				stack = append(stack, neighbour)
			}
		}
	}
}

func (g Graph) DFSpreorder(rootVertex int) {
	fmt.Println("running depth first search (preorder)")
	// edge case vertex doesnt exist
	if g.adjList[rootVertex] == nil {
		fmt.Println("vertex not found")
		return
	}

	stack := []int{rootVertex}
	visited := make(map[int]bool)

	visited[rootVertex] = true
	for len(stack) > 0 {
		currVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		fmt.Println("processing vertex", currVertex)
		visited[currVertex] = true
		
		// to reverse the children's order
		tmpStack := []int{}

		for neighbour := range g.adjList[currVertex] {
			if !visited[neighbour] {
				tmpStack = append(tmpStack, neighbour)
			}
		}

		// reverse the tmp stack then put the items into the real stack
		for i, j := 0, len(tmpStack)-1; i < j; i, j = i+1, j-1 {
			tmpStack[i], tmpStack[j] = tmpStack[j], tmpStack[i]
		}
		stack = append(stack, tmpStack...)
	}
}

func main() {
	// below is replicated graph from this wiki page
	// https://en.wikipedia.org/wiki/Breadth-first_search
	bfsGraph := newGraph()
	bfsGraph.addVertexInRange(1, 12)
	bfsGraph.addEdge(1,2)
	bfsGraph.addEdge(1,3)
	bfsGraph.addEdge(1,4)
	bfsGraph.addEdge(2,5)
	bfsGraph.addEdge(2,6)
	bfsGraph.addEdge(5,9)
	bfsGraph.addEdge(5,10)
	bfsGraph.addEdge(4,7)
	bfsGraph.addEdge(4,8)
	bfsGraph.addEdge(7,11)
	bfsGraph.addEdge(7,12)

	bfsGraph.BFS(1)


	// below is replicated graph from this wiki page
	// https://en.wikipedia.org/wiki/Depth-first_search
	dfsGraph := newGraph()
	dfsGraph.addVertexInRange(1, 12)
	dfsGraph.addEdge(1,2)
	dfsGraph.addEdge(1,7)
	dfsGraph.addEdge(1,8)
	dfsGraph.addEdge(2,3)
	dfsGraph.addEdge(2,6)
	dfsGraph.addEdge(3,4)
	dfsGraph.addEdge(3,5)
	dfsGraph.addEdge(8,9)
	dfsGraph.addEdge(8,12)
	dfsGraph.addEdge(9,10)
	dfsGraph.addEdge(9,11)

	dfsGraph.DFS(1)
	dfsGraph.DFSpreorder(1)

	// the ordering is still unpredictable despite attempts to order them
	// this is due to using maps as the graph implementation
	// and map's values are unpredictable when iterated over using for loop
}
