package graph

import (
	"errors"
)

type vertex struct {
	idx   int
	edges []*edge
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

func contains(edges []*edge, searchIdx int) bool {

	for _, edge := range edges {
		if edge.neighbor.idx == searchIdx {
			return true
		}
	}

	return false
}
