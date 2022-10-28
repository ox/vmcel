package main

import (
	"github.com/rivo/tview"
)

func main() {
	appRoot := tview.NewFlex()

	g := NewGraph()

	a, _ := NewNode("a", []NodePort{{Name: "in", DefaultValue: 0}, {Name: "foo", DefaultValue: 0}}, []NodePort{{Name: "out", DefaultValue: 0}}, "outputs['out'] = inputs['in'] + 1")
	g.AddNode(a)
	b, _ := NewNode("Node balacazam", []NodePort{{Name: "in", DefaultValue: 0}}, []NodePort{{Name: "out", DefaultValue: 0}}, "outputs['out'] = inputs['in'] + 1")
	g.AddNode(b)
	c, _ := NewNode("Node cadabra", []NodePort{{Name: "in", DefaultValue: 0}}, []NodePort{{Name: "out", DefaultValue: 0}}, "outputs['out'] = inputs['in'] + 1")
	g.AddNode(c)
	ab := NewEdge(a, "out", b, "in")
	g.AddEdge(ab)
	bc := NewEdge(b, "out", c, "in")
	g.AddEdge(bc)
	ca := NewEdge(c, "out", a, "in")
	g.AddEdge(ca)

	gview := NewGraphView(g)
	gview.SetBorder(true).SetTitle("Graph View")
	appRoot.AddItem(gview, 0, 2, false)
	box := tview.NewBox().SetTitle("Hello").SetBorder(true)
	appRoot.AddItem(box, 0, 1, false)

	if err := tview.NewApplication().SetRoot(appRoot, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
