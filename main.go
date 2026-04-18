/*
This package helps understand a Ren'Py source code by drawing a graph from the source code
*/
package main

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"Flowcharter/parser"
)

func main() {
	path, options := PlugCLI()

	content := parser.GetRenpyContent(path, options)

	graph, err := parser.Graph(content, options)
	if err != nil {
		parser.DocumentIssue(err)
	}

	err = graph.CreateFile("Flowchart.dot")
	if err != nil {
		parser.DocumentIssue(err)
	}

	drawGraph("Flowchart.dot", "Flowchart.png")

	if graph.Options.OpenFile {
		err = open.Run("Flowchart.png")
		if err != nil {
			fmt.Println("A Flowchart.png image file has been created, but couldn't be open. Please open it manually.")
		}
	}
}
