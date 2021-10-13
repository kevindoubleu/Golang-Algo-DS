package main

import "fmt"

type Graph struct {
	// map of ints (vertices) that maps to
	// other vertices (keys) with their respective weights (values)
	adjList map[int]map[int]int
}

func newGraph() *Graph {
	return &Graph{adjList: make(map[int]map[int]int)}
}

func (g Graph) addVertex(value int) {
	if g.adjList[value] == nil {
		g.adjList[value] = make(map[int]int)
	}
}

func (g Graph) addEdge(v1, v2, weight int) {
	if g.adjList[v1] != nil && g.adjList[v2] != nil {
		g.adjList[v1][v2] = weight // edge from v1 to v2
		// directed, so just dont add the other edge
		// g.adjList[v2][v1] = weight
	}
}

func main() {
	g1 := newGraph()

	g1.addVertex(10)
	g1.addVertex(20)
	g1.addVertex(30)

	g1.addEdge(10, 20, 50)
	g1.addEdge(10, 30, 100)
	g1.addEdge(11, 22, 1000)

	fmt.Println(g1)

	g1.addEdge(20, 10, 5)
	g1.addEdge(30, 10, 10)

	fmt.Println(g1)
}