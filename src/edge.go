package main

import (
	"errors"
	"fmt"
)

type Edge struct {
	From *Node
	FromOutput string
	To *Node
	ToInput string
}

func NewEdge(from *Node, fromOutput string, to *Node, toInput string) *Edge {
	e := Edge{
		From: from,
		FromOutput: fromOutput,
		To: to,
		ToInput: toInput,
	}
	return &e
}

func (e *Edge) Sync() error {
	if val, okInput := e.From.Outputs[e.FromOutput]; okInput {
		_, okInput := e.To.Inputs[e.ToInput]
		if okInput {
			e.To.Inputs[e.ToInput] = val
			return nil
		} else {
			return errors.New(fmt.Sprintf("Input named %s does not exist on cell %s", e.ToInput, e.To.Name))
		}
	}

	return errors.New(fmt.Sprintf("Node %s does not have an output named %s: %v", e.From.Name, e.FromOutput, e.From))
}