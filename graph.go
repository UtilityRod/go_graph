package graph

import (
	"errors"
	"fmt"
)

type vertex struct {
	idx     int
	edges   []*edge
	visited bool
}

type edge struct {
	neighbor *vertex
	weight   int
}

type Graph struct {
	v []*vertex
}

func (g *Graph) AddVertex() {
	newVertex := new(vertex)
	newVertex.idx = int(len(g.v))
	newVertex.visited = false
	g.v = append(g.v, newVertex)
}

func (g *Graph) AddEdge(vOneIdx, vTwoIdx, weight int) error {
	vertexOne := g.v[vOneIdx]
	vertexTwo := g.v[vTwoIdx]

	if contains(vertexOne.edges, vTwoIdx) || contains(vertexTwo.edges, vOneIdx) {
		return errors.New("vertecies already contain edge to each other")
	}

	edgeOne := new(edge)
	edgeOne.weight = weight
	edgeOne.neighbor = vertexTwo
	vertexOne.edges = append(vertexOne.edges, edgeOne)

	edgeTwo := new(edge)
	edgeTwo.weight = weight
	edgeTwo.neighbor = vertexOne
	vertexTwo.edges = append(vertexTwo.edges, edgeTwo)

	return nil
}

func (g *Graph) PrintNeighbhors(vertexIdx int) {
	fmt.Printf("Neighbors for vertex with index %d\n", vertexIdx)
	for _, edge := range g.v[vertexIdx].edges {
		fmt.Printf("Weight: %d Vertex Idx: %d\n", edge.weight, edge.neighbor.idx)
	}
}

func (g *Graph) Dijkstra(start int, end int) []int {
	var prev []*vertex
	var dist []int
	var queue []*vertex

	// Set prev and dist slices to null values
	for i := 0; i < len(g.v); i++ {
		prev = append(prev, nil)
		dist = append(dist, 100000)
		queue = append(queue, g.v[i])
	}

	dist[start] = 0

	for len(queue) != 0 {
		currentIdx := -1
		removeIdx := -1

		for idx, v := range queue {
			if currentIdx == -1 || dist[currentIdx] > dist[v.idx] {
				removeIdx = idx
				currentIdx = v.idx
			}
		}

		queue = remove(queue, removeIdx)

		for _, edge := range g.v[currentIdx].edges {
			if !containsVertex(queue, edge.neighbor.idx) {
				continue
			}

			alt := dist[currentIdx] + edge.weight
			if alt < dist[edge.neighbor.idx] {
				dist[edge.neighbor.idx] = alt
				prev[edge.neighbor.idx] = g.v[currentIdx]
			}
		}
	}

	currentVertex := g.v[end]
	var path []int
	for currentVertex != g.v[start] {
		path = append(path, currentVertex.idx)
		currentVertex = prev[currentVertex.idx]
	}

	return path
}

func remove(slice []*vertex, s int) []*vertex {
	return append(slice[:s], slice[s+1:]...)
}

func containsVertex(slice []*vertex, s int) bool {
	for _, vertex := range slice {
		if vertex.idx == s {
			return true
		}
	}

	return false
}

func contains(edges []*edge, searchIdx int) bool {

	for _, edge := range edges {
		if edge.neighbor.idx == searchIdx {
			return true
		}
	}

	return false
}
