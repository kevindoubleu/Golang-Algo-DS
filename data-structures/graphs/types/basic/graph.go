package main

import "fmt"

type Graph struct {
	adjList map[int][]int
}

func newGraph() *Graph {
	return &Graph{adjList: make(map[int][]int)}
}

func (g Graph) addVertex(value int) {
	if g.adjList[value] == nil {
		g.adjList[value] = make([]int, 0)
	}
}

func (g Graph) addEdge(v1, v2 int) {
	if g.adjList[v1] != nil && g.adjList[v2] != nil {
		g.adjList[v1] = append(g.adjList[v1], v2)
		g.adjList[v2] = append(g.adjList[v2], v1)
	}
}

func main() {
	g1 := newGraph()

	g1.addVertex(10)
	g1.addVertex(20)
	g1.addVertex(30)

	g1.addEdge(10, 20)
	g1.addEdge(10, 30)
	g1.addEdge(11, 22)

	fmt.Println(g1)
}