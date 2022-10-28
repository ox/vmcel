package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var noports = make([]NodePort, 0)
var inports = []NodePort{{Name: "in", DefaultValue: 0}}
var outports = []NodePort{{Name: "out", DefaultValue: 0}}

func TestNewNode(t *testing.T) {
	name := "foo"
	n,err := NewNode(name, inports, outports, "code")
	assert.Equal(t, nil, err)
	assert.Equal(t, name, n.Name)
	assert.Equal(t, 1, len(n.Inputs))
	assert.Equal(t, 1, len(n.Outputs))
}

func TestNewNode_CompError(t *testing.T) {
	n,err := NewNode("foo", noports, noports, "nonsense, program;,c,d")
	assert.Nil(t, n)
	assert.Error(t, err)
}

func TestNode_Step(t *testing.T) {
	n, err := NewNode("foo", noports, noports, "1+1")
	assert.Equal(t, nil, err)
	assert.Equal(t, nil, n.Step())
}

func TestNode_Step_ReadInputs(t *testing.T) {
	n, err := NewNode("foo", inports, outports, "inputs['in'] + 1")
	assert.Equal(t, nil, err)

	n.Inputs["in"] = 2
	assert.Equal(t, nil, n.Step())
}

func TestNode_Step_WriteOutputs(t *testing.T) {
	n, err := NewNode("foo", inports, outports, "outputs['out'] = inputs['in'] + 1")
	assert.Equal(t, nil, err)

	n.Inputs["in"] = 2
	n.Outputs["out"] = 0
	assert.Equal(t, nil, n.Step())

	assert.Equal(t, int64(3), n.Outputs["out"])

}