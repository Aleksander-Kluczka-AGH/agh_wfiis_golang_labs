package main

import (
	"fmt"
	"math"
)

type Node struct {
	id       int
	inEdges  []int
	outEdges []int
}

type Graph struct {
	nodes []Node
}

func (graph Graph) hasNode(node_id int) bool {
	for i := range graph.nodes {
		if graph.nodes[i].id == node_id {
			return true
		}
	}
	return false
}

func (graph *Graph) addNode(node_id int) {

	if graph.hasNode(node_id) {
		fmt.Println("Can't add an existing node to the graph")
		return
	}

	node := Node{node_id, []int{}, []int{}}
	graph.nodes = append(graph.nodes, node)
}

func (graph *Graph) addEdge(start int, end int) {
	if !graph.hasNode(start) {
		fmt.Println("Graph does not have a Node with id = ", start)
		return
	}
	if !graph.hasNode(end) {
		fmt.Println("Graph does not have a Node with id = ", end)
		return
	}

	for i := range graph.nodes {
		if graph.nodes[i].id == start {
			graph.nodes[i].outEdges = append(graph.nodes[i].outEdges, end)
		}
		if graph.nodes[i].id == end {
			graph.nodes[i].inEdges = append(graph.nodes[i].inEdges, start)
		}
	}
}

func (graph Graph) print() {
	for i := range graph.nodes {
		var node *Node = &graph.nodes[i]

		fmt.Println("Node", node.id, ", in:", node.inEdges, ", out:", node.outEdges)
	}
}

func (graph Graph) ranks() (map[int]int, map[int]int) {
	ranks_in := map[int]int{}
	ranks_out := map[int]int{}
	for i := range len(graph.nodes) - 1 {
		ranks_in[i+1] = 0
		ranks_out[i+1] = 0
	}

	for i := range graph.nodes {
		var node *Node = &graph.nodes[i]
		ranks_in[len(node.inEdges)]++
		ranks_out[len(node.outEdges)]++
	}
	return ranks_in, ranks_out
}

func (graph Graph) getNode(id int) *Node {
	for i := range graph.nodes {
		if graph.nodes[i].id == id {
			return &graph.nodes[i]
		}
	}
	return nil
}

func (graph Graph) hasNeighbor(target int, neighbor int) bool {
	node := graph.getNode(target)

	if node == nil {
		return false
	}

	for i := range node.outEdges {
		if node.outEdges[i] == neighbor {
			return true
		}
	}
	return false
}

func (graph Graph) distanceBetweenNodes(start_id int, end_id int, visited_nodes []int) int {
	if start_id == end_id {
		return 0
	}

	if graph.hasNeighbor(start_id, end_id) {
		return 1
	}

	node := graph.getNode(start_id)

	visited_nodes = append(visited_nodes, start_id)
	if len(node.outEdges) == 0 {
		return 0
	}

	distance := 1
	for i := range node.outEdges {
		if graph.hasNeighbor(node.outEdges[i], end_id) {
			return 2
		} else {
			is_visited := false
			for j := range visited_nodes {
				if visited_nodes[j] == node.outEdges[i] {
					is_visited = true
					break
				}
			}
			if is_visited {
				continue
			}
			distance += graph.distanceBetweenNodes(node.outEdges[i], end_id, visited_nodes)
		}
	}
	return distance
}

func (graph Graph) shortestPath() {
	size := len(graph.nodes)
	dist_map := make([][]int, size)
	for i := range dist_map {
		dist_map[i] = make([]int, size)
	}

}

func floydWarshall(dist [][]int) {
	for k := 0; k < 4; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				dist[i][j] = int(math.Min(float64(dist[i][j]), float64(dist[i][k]+dist[k][j])))
			}
		}
	}
}
