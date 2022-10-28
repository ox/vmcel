package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"sort"
)

const NodePortInputFormatStr = "%s"
const NodePortOutputFormatStr = "%s"

type NodeView struct {
	*tview.Box
	node *Node
	dragNodeX, dragNodeY int
}

func calcNodeWidth(n *Node, spacing int) int {
	maxIn := 0
	for name := range n.Inputs {
		portLen := len(fmt.Sprintf(NodePortInputFormatStr, name))
		if portLen > maxIn {
			maxIn = portLen
		}
	}

	maxOut := 0
	for name := range n.Outputs {
		portLen := len(fmt.Sprintf(NodePortOutputFormatStr, name))
		if portLen > maxOut {
			maxOut = portLen
		}
	}

	internalMinimumWidth := maxIn + spacing + maxOut + 2
	if len(n.Name) > internalMinimumWidth {
		return len(n.Name) + 4 // add some space to title
	}

	return internalMinimumWidth
}

func NewNodeView(node *Node) *NodeView {
	return &NodeView{
		Box:       tview.NewBox().SetBorder(true).SetTitle(node.Name),
		node:      node,
		dragNodeX: -1,
		dragNodeY: -1,
	}
}

func (nv *NodeView) Draw(screen tcell.Screen) {
	nv.Box.DrawForSubclass(screen, nv)
	x, y, width, _ := nv.GetInnerRect()

	inputs := make([]string, 0)
	for in := range nv.node.Inputs {
		inputs = append(inputs, in)
	}
	sort.Strings(inputs)
	for i, in := range inputs {
		tview.Print(screen, fmt.Sprintf(NodePortInputFormatStr, in), x, y + i, width, tview.AlignLeft, tcell.ColorWhite)
	}

	outputs := make([]string, 0)
	for out := range nv.node.Outputs {
		outputs = append(outputs, out)
	}
	sort.Strings(outputs)
	for i, out := range outputs {
		tview.Print(screen, fmt.Sprintf(NodePortOutputFormatStr, out), x, y + i, width, tview.AlignRight, tcell.ColorWhite)
	}
}

func (nv *NodeView) Focus(delegate func(p tview.Primitive)) {
	nv.Box.Focus(delegate)
}

func (nv *NodeView) Blur() {
	nv.Box.Blur()
}

func (nv *NodeView) HasFocus() bool {
	return nv.Box.HasFocus()
}

func (nv *NodeView) MouseHandler() func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
	return nv.Box.WrapMouseHandler(func(action tview.MouseAction, event *tcell.EventMouse, setFocus func(p tview.Primitive)) (consumed bool, capture tview.Primitive) {
		if !nv.InRect(event.Position()) {
			return false, nil
		}

		if action == tview.MouseLeftDown || action == tview.MouseMiddleDown || action == tview.MouseRightDown {
			setFocus(nv)
		}

		if action == tview.MouseLeftDown {
			x, y, width, _ := nv.GetRect()
			mouseX, mouseY := event.Position()
			topEdge := mouseY == y
			if mouseX >= x && mouseX <= x+width-1 && topEdge {
				nv.dragNodeX = mouseX - x
				nv.dragNodeY = mouseY - y
			}
		}

		return true, capture
	})
}


