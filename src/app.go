package main

import "github.com/rivo/tview"

type App struct {
	tview.Application

	graph *Graph
}

func NewApp() *App {
	a := App{
		graph: NewGraph(),
	}

	return &a
}

