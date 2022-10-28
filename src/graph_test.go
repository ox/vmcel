package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph()
	assert.Equal(t, 0, len(g.edges))
	assert.Equal(t, 0, len(g.nodes))
}

func TestGraph_AddNodeAndEdge(t *testing.T) {
	g := NewGraph()
	from, fromErr := NewNode("from", noports, outports, "outputs['out'] = 1")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", inports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)

	g.AddNode(from)
	g.AddNode(to)
	g.AddEdge(e)
}

func TestGraph_Step(t *testing.T) {
	g := NewGraph()
	from, fromErr := NewNode("from", noports, outports, "outputs['out'] = 1")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", inports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)

	g.AddNode(from)
	g.AddNode(to)
	g.AddEdge(e)

	assert.Nil(t, g.Step())
	assert.Equal(t, int64(1), from.Outputs["out"])
	assert.Equal(t, int64(1), to.Inputs["in"])
}


func TestGraph_StepCycle(t *testing.T) {
	g := NewGraph()

	a, fromErr := NewNode("a", inports, outports, "outputs['out'] = inputs['in'] + 1")
	assert.Nil(t, fromErr)
	g.AddNode(a)

	b, toErr := NewNode("b", inports, outports, "outputs['out'] = inputs['in'] + 1")
	assert.Nil(t, toErr)
	g.AddNode(b)

	c, toErr := NewNode("c", inports, outports, "outputs['out'] = inputs['in'] + 1")
	assert.Nil(t, toErr)
	g.AddNode(c)

	ab := NewEdge(a, "out", b, "in")
	g.AddEdge(ab)
	bc := NewEdge(b, "out", c, "in")
	g.AddEdge(bc)
	ca := NewEdge(c, "out", a, "in")
	g.AddEdge(ca)

	// Run the graph N times and make sure that all inputs and outputs are mutated 
	for i := 0; i < 5; i++ {
		assert.Nil(t, g.Step())
		for _, node := range g.nodes {
			assert.Equal(t, int64(i+1), node.Outputs["out"])
			assert.Equal(t, int64(i+1), node.Inputs["in"])
		}
	}
}

