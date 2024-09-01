package main

import (
	"fmt"
	"math/rand"
)

// "gonum.org/v1/plot"
// "gonum.org/v1/plot/plotter"
// "gonum.org/v1/plot/vg"
// "image/color"
// "graph/graph"
//

// //...
// plt := plot.New()
// plt.Title.Text = ""
// plt.X.Label.Text = ""
// plt.Y.Label.Text = ""

// pts := make(plotter.XYs, liczba_punktów)

// for i := 0; i < liczba_punktów; i++ {
// 	pts[i].X = ...
// 	pts[i].Y = ...
// }

// line, _ := plotter.NewLine(pts)
// line.LineStyle.Color = color.RGBA{R: 1, G: 1, B: 1, A: 255}
// plt.Add(line)
// plt.Save(8*vg.Inch, 4*vg.Inch, name)

func main() {
	const N int = 10
	var graph Graph = generateRandomGraph(10)

	graph.print()
	fmt.Println(graph.ranks())

	// dist := graph.distanceBetweenNodes(0, 1, []int{})
	// fmt.Println(dist)
}

func generateRandomGraph(N int) Graph {
	graph := Graph{}
	for i := range N {
		graph.addNode(i)
	}

	for i := range N {
		has_in_edge := rand.Intn(N) < 5
		has_out_edge := rand.Intn(N) < 5

		target_node := i
		for target_node == i {
			target_node = rand.Intn(10)
		}

		if has_in_edge {
			graph.addEdge(target_node, graph.nodes[i].id)
		}

		if has_out_edge {
			graph.addEdge(graph.nodes[i].id, target_node)
		}
	}
	return graph
}
