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
		if visited[currVertex.(int)] {
			continue
		}

		fmt.Println("processing vertex", currVertex, g.adjList[currVertex.(int)])
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
		if visited[currVertex] {
			continue
		}

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

	for len(stack) > 0 {
		currVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[currVertex] {
			continue
		}
		
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
	// below is replicated graph from
	// https://medium.com/@crusso_22624/bfs-vs-dfs-graph-traversals-comparing-times-80729f100bf
	g1 := newGraph()
	g1.addVertex(10)
	g1.addVertex(20)
	g1.addVertex(30)
	g1.addVertex(40)
	g1.addVertex(50)
	g1.addVertex(60)
	g1.addVertex(70)
	g1.addEdge(40, 20)
	g1.addEdge(40, 10)
	g1.addEdge(10, 30)
	g1.addEdge(20, 10)
	g1.addEdge(20, 30)
	g1.addEdge(20, 60)
	g1.addEdge(20, 50)
	g1.addEdge(30, 60)
	g1.addEdge(50, 70)
	g1.addEdge(60, 70)

	// g1.BFS(40)
	// g1.DFS(40)
	g1.DFSpreorder(40)

	// the ordering is still unpredictable despite attempts to order them
	// this is due to using maps as the graph implementation
	// and map's values are unpredictable when iterated over using for loop
}
