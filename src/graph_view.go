package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"math"
)

type GraphView struct {
	*tview.Box
	graph *Graph
	nodeViews []*NodeView
	draws int
}

func NewGraphView(g *Graph) *GraphView {
	nodeViews := make([]*NodeView, 0)
	x := 0
	for _, node := range g.nodes {
		nv := NewNodeView(node)
		width := calcNodeWidth(node, 1)
		nv.SetRect(4 + x, 5, width, int(math.Max(float64(len(node.Inputs)), float64(len(node.Outputs)))) + 2)
		x += width
		nodeViews = append(nodeViews, nv)
	}

	return &GraphView{
		Box: tview.NewBox(),
		graph: g,
		nodeViews: nodeViews,
		draws: 0,
	}
}

func (gv *GraphView) Draw(screen tcell.Screen) {
	gv.Box.DrawForSubclass(screen, gv)
	x, y, width, height := gv.GetInnerRect()

	gv.draws += 1

	line := fmt.Sprintf("%d nodes and %d edges, draws %d", len(gv.graph.nodes), len(gv.graph.edges), gv.draws)
	tview.Print(screen, line, x, y + height / 2, width, tview.AlignCenter, tcell.ColorYellow)

	for _, nv := range gv.nodeViews {

		nv.Draw(screen)
	}
}

func (gv *GraphView) Focus(delegate func(p tview.Primitive))  {
	if len(gv.nodeViews) == 0 {
		return
	}
	// Not sure why we need this
	gv.nodeViews[len(gv.nodeViews) - 1].Focus(delegate)
}

func (gv *GraphView) HasFocus() bool {
	for _, node := range gv.nodeViews {
		if node.HasFocus() {
			return true
		}
	}
	return false
}

func (gv *GraphView) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return gv.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	})
}

func (gv *GraphView) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return gv.Box.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		if !gv.InRect(event.Position()) {
			return false, nil
		}

		mouseX, mouseY := event.Position()

		if action == tview.MouseMove {
			for _, nv := range gv.nodeViews {
				x, y, width, height := nv.GetRect()

				if nv.dragNodeX != -1 || nv.dragNodeY != -1 {
					offsetX := x - mouseX
					offsetY := y - mouseY
					x -= offsetX + nv.dragNodeX
					y -= offsetY + nv.dragNodeY
					nv.SetRect(x, y, width, height)
					consumed = true
				}
			}
		} else if action == tview.MouseLeftUp {
			for _, nv := range gv.nodeViews {
				nv.dragNodeX, nv.dragNodeY = -1, -1
			}
		}

		var focusNv *NodeView
		for _, nv := range gv.nodeViews {
			if nv.InRect(mouseX, mouseY) {
				focusNv = nv
			}
		}

		if focusNv != nil {
			return focusNv.MouseHandler()(action, event, setFocus)
		}

		return consumed, capture
	})
}