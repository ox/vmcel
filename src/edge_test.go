package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewEdge(t *testing.T) {
	from, fromErr := NewNode("from", noports, outports, "")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", inports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)
}

func TestEdge_Sync(t *testing.T) {
	from, fromErr := NewNode("from", noports, outports, "outputs['out'] = 1")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", inports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)

	assert.Nil(t, from.Step())
	assert.Nil(t, to.Step())
	assert.Nil(t, e.Sync())
	assert.Equal(t, int64(1), to.Inputs["in"])
}

func TestEdge_Sync_OutputDoesntExist(t *testing.T) {
	from, fromErr := NewNode("from", noports, noports, "outputs['out'] = 1")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", inports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)

	assert.Error(t, e.Sync())
}


func TestEdge_Sync_InputDoesntExist(t *testing.T) {
	from, fromErr := NewNode("from", noports, outports, "outputs['out'] = 1")
	assert.Nil(t, fromErr)
	to, toErr := NewNode("to", noports, noports, "")
	assert.Nil(t, toErr)
	e := NewEdge(from, "out", to, "in")
	assert.NotNil(t, e)

	assert.Error(t, e.Sync())
}