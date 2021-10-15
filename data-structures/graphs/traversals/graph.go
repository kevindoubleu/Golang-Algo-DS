package main

type Graph struct {
	// map of ints (vertices) that maps to
	// other vertices (keys) with their respective weights (values)
	adjList map[int]map[int]int
}

func newGraph() *Graph {
	return &Graph{adjList: make(map[int]map[int]int)}
}

// inclusive
func (g Graph) addVertexInRange(from, to int) {
	for i := from; i <= to; i++ {
		g.addVertex(i)
	}
}

func (g Graph) addVertex(value int) {
	if g.adjList[value] == nil {
		g.adjList[value] = make(map[int]int)
	}
}

func (g Graph) addEdge(v1, v2 int) {
	g.addEdgeWithWeight(v1, v2, 0)
}

func (g Graph) addEdgeWithWeight(v1, v2, weight int) {
	if g.adjList[v1] != nil && g.adjList[v2] != nil {
		g.adjList[v1][v2] = weight
	}
}
