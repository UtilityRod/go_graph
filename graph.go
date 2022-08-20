package graph

import (
	"errors"
	"fmt"
)

type vertex struct {
	idx       uint64
	neighbors []*vertex
}

type Graph struct {
	v []*vertex
}

func (g *Graph) AddVertex() {
	newVertex := new(vertex)
	newVertex.idx = uint64(len(g.v))
	g.v = append(g.v, newVertex)
}

func (g *Graph) AddEdge(vOneIdx, vTwoIdx uint64) error {
	vertexOne := g.v[vOneIdx]
	vertexTwo := g.v[vTwoIdx]

	if contains(vertexOne.neighbors, vTwoIdx) || contains(vertexTwo.neighbors, vOneIdx) {
		return errors.New("vertecies already contain edge to each other")
	}

	vertexOne.neighbors = append(vertexOne.neighbors, vertexTwo)
	vertexTwo.neighbors = append(vertexTwo.neighbors, vertexOne)

	return nil
}

func (g *Graph) PrintNeighbhors(vIdx uint64) {
	vertex := g.v[vIdx]
	fmt.Printf("Neighbors for vertex with index %d\n", vIdx)
	for _, neighbor := range vertex.neighbors {
		fmt.Println(neighbor.idx)
	}
}

func contains(neighbhors []*vertex, searchIdx uint64) bool {

	for _, v := range neighbhors {
		if v.idx == searchIdx {
			return true
		}
	}

	return false
}
