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
//           [0, 0]=>0, [0, 1]=>2.24, [0, 2]=>1.0,
//           [1, 0]=>2.24, [1, 1]=>0, [1, 2]=>1.41,
//           [2, 0]=>1.0, [2, 1]=>1.41, [2, 2]=>0
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
)

type Coords struct {
	X int `csv:"x"`
	Y int `csv:"y"`
}

func main() {
	data := csv()

	for _, coord := range data {
		fmt.Println(coord.X, coord.Y)
	}
}

func csv() []*Coords {
	coordsFile, err := os.OpenFile("./sample.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)

	if err != nil {
		panic(err)
	}
	defer coordsFile.Close()

	coords := []*Coords{}

	if err := gocsv.UnmarshalFile(coordsFile, &coords); err != nil { // Load coords from file
		panic(err)
	}

	return coords
}