package main

import "github.com/pkg/errors"

type Graph struct {
	nodes map[string]*Node
	edges []*Edge
}

func NewGraph() *Graph {
	g := Graph{
		nodes: make(map[string]*Node),
		edges: make([]*Edge, 0),
	}
	return &g
}

func (g *Graph) AddNode(n *Node) {
	g.nodes[n.Name] = n
}

func (g *Graph) AddEdge(e *Edge) {
	g.edges = append(g.edges, e)
}

func (g *Graph) Step() error {
	// First step every node to load up their outputs
	for name, node := range g.nodes {
		if err := node.Step(); err != nil {
			return errors.Wrapf(err, "error evaluating node %s", name)
		}
	}

	// Then sync every edge to move the new outputs to their appropriate inputs
	for _, edge := range g.edges {
		if err := edge.Sync(); err != nil {
			return errors.Wrapf(err, "error in edge from %s[%s] to %s[%s]", edge.From.Name, edge.FromOutput, edge.To.Name, edge.ToInput)
		}
	}

	return nil
}