package main

import (
	"github.com/pkg/errors"
	"github.com/dop251/goja"
)

type NodePort struct {
	Name string
	DefaultValue any
}

type Node struct {
	Name string
	Inputs map[string]any
	Outputs map[string]any
	OutputsUpdated bool
	program        *goja.Program
	runtime        *goja.Runtime
}

func NewNode(name string, inputs, outputs []NodePort, src string) (*Node, error) {
	p, perr := goja.Compile(name, src, true)
	if perr != nil {
		return nil, errors.Wrapf(perr, "error compiling src for node %s", name)
	}

	c := Node{
		Name:           name,
		Inputs:         make(map[string]any),
		Outputs:        make(map[string]any),
		OutputsUpdated: false,
		program:        p,
		runtime:        goja.New(),
	}

	for _, k := range inputs {
		c.Inputs[k.Name] = k.DefaultValue
	}

	for _, k := range outputs {
		c.Outputs[k.Name] = k.DefaultValue
	}

	return &c, nil
}

func (c *Node) getInput(name string) any {
	if v, ok := c.Inputs[name]; ok {
		return v
	}
	return nil
}

func (c *Node) setOutput(name string, value any) {
	c.Outputs[name] = value
}

func (c *Node) Step() error {
	if err := c.runtime.Set("inputs", c.Inputs); err != nil {
		return errors.Wrap(err, "could not set inputs object")
	}

	if err := c.runtime.Set("outputs", c.Outputs); err != nil {
		return errors.Wrap(err, "could not set inputs object")
	}

	if _, err := c.runtime.RunProgram(c.program); err != nil {
		return errors.Wrap(err, "could not run program")
	}

	outs := c.runtime.Get("outputs").ToObject(c.runtime)
	for k := range c.Outputs {
		c.Outputs[k] = outs.Get(k).Export()
	}

	return nil
}
