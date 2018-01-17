// Accepts a CSV file and outputs a map showing the distance between all nodes.
//
// Example CSV:
// Row 1: 1, 2 (Node 0)
// Row 2: 2, 4 (Node 1)
// Row 3: 1, 3 (Node 2)
//
// These values represents points a grid we will create. They are coordinates (x, y).
//
// Imagine a (zero indexed) 4 x 4 grid:
//
//        0     1      2      3
//      ----------------------------
//   0  | x     Node0  Node2  x
//   1  | x     x      x      Node1
//   2  | x     x      x      x
//   3  | x     x      x      x
//
// Output: {
//           [0, 0]: 0, [0, 1]: 2.24, [0, 2]: 1.0,
//           [1, 0]: 2.24, [1, 1]: 0, [1, 2]: 1.41,
//           [2, 0]: 1.0, [2, 1]: 1.41, [2, 2]: 0
//         }
//
// Show thats the straightline distance from:
// =>  Node 0 (0, 1) to Node 1 (1, 3) is 2.24.
// =>  Node 0 (0, 1) to Node 2 (0, 2) is 1.0. (I.e. just adjacent horizontally.)
//
// =>  Node 2 (0, 2) to Node 1 (1, 3) is 1.41.

package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"math"
)

type Graph map[[2]int]float64

type CoordPair struct {
	X int `csv:"x"`
	Y int `csv:"y"`
}

func build() Graph {
	numNodes()
	// Slices cannot be used as keys as they do not have equality defined.
	// Arrays however can be used as they do.
	output := make(Graph)

	for i, coord := range csv() {
		for inner_i, inner_coord := range csv() {
			if i == inner_i {
				output[[2]int{i, inner_i}] = 0
			} else {
				output[[2]int{i, inner_i}] = distanceCalc(coord, inner_coord)
			}
		}
	}

	fmt.Println("Graph Built with straight line distances:", output)
	return output
}


func distanceCalc(inner_coords *CoordPair, outer_coords *CoordPair) float64 {
	var sum_of_squares float64 = 0

	inner := convertStructToSlice(inner_coords)
	outer := convertStructToSlice(outer_coords)

	for i := 0; i < len(inner); i++ {
		sum_of_squares += math.Pow(float64(inner[i] - outer[i]), float64(2))
	}

	returnVal := math.Sqrt(sum_of_squares)
	return returnVal
}

func convertStructToSlice(inputStruct *CoordPair) []int {
	output := make([]int, 2)
	output = append(output, inputStruct.X)
	output = append(output, inputStruct.Y)
	return output
}

func numNodes() int {
	return len(csv())
}

func csv() []*CoordPair {
	coordsFile, err := os.OpenFile("./sample.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer coordsFile.Close()

	coords := []*CoordPair{}

	if err := gocsv.UnmarshalFile(coordsFile, &coords); err != nil { // Load coords from file
		panic(err)
	}

	return coords
}