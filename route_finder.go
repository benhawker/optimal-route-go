package main

import "fmt"
import "math/rand"


type Solution [][2]int

func main() {
	graph := build()

	for i := 0; i < 10; i++ {
		fmt.Println(generateRoute(graph))
	}

	// Still to complete - currently just printing some potential routes.
}

func generateRoute(graph Graph) Solution {
	number_of_nodes := numNodes() // Number of coords in the csv
	nodes_to_explore := []int{}

	for i := 1; i <= number_of_nodes; i++ {
		nodes_to_explore = append(nodes_to_explore, i)
	}

	edges := Solution{}
	last_node := 0

	for i := 1; i <= number_of_nodes; i++ {
		// Pick a node at random to explore. When there are no more, revert to 0 (back to start).
		node := pickRandNode(nodes_to_explore)

		// Append the `origin` and `destination` node.
		coordsToAdd := [2]int{last_node, node}
		edges = append(edges, coordsToAdd)

		// We have now traversed this so delete it from nodes_to_explore.
		for i := 0; i < len(nodes_to_explore); i++ {
			if (nodes_to_explore[i] == node) {
				// Delete the node from the nodes to explore slice.
				nodes_to_explore = append(nodes_to_explore[:i], nodes_to_explore[(i+1):]...)
			}
		}
		// Set last last_node to the current node and continue.
		last_node = node
	}
	return edges
}	

func pickRandNode(nodes []int) int {
	max := len(nodes)
	randomIndex := rand.Intn(max)
	
	// If the element exists...
	if len(nodes) > randomIndex{
		return randomIndex
	} else {
		return 0
	}
}