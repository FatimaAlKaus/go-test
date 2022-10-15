package main

import (
	"flag"
	"fmt"
	"graph/internal/graph"
	"log"
	"strconv"
)

var (
	_vertices = flag.Int("v", 8, "count of vertices")
	_edges    = flag.Int("e", 10, "count of edges")
)

func main() {
	flag.Parse()
	fmt.Printf("Vertices: %d\n", *_vertices)
	fmt.Printf("Edges: %d\n", *_edges)
	var matrix [][]int
	if len(flag.Args()) == 0 {
		matrix = graph.GenMatrix(*_vertices, *_edges)
	} else {
		matrix = readMatrix(flag.Args(), *_vertices, *_edges)
	}
	fmt.Printf("Matrix:\n")
	graph.ShowMatrix(matrix)
	edges, err := graph.GetEdges(matrix)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Result:\n")
	for _, e := range edges {
		fmt.Printf("%+v\n", e)
	}
}

func readMatrix(args []string, vertices, edges int) [][]int {
	if len(args) < vertices*edges {
		panic("not enough args")
	}

	count := 0
	matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		matrix[i] = make([]int, edges)
		for j := 0; j < edges; j++ {
			n, err := strconv.Atoi(args[count])
			if err != nil {
				panic(err.Error())
			}
			matrix[i][j] = n
			count++
		}
	}

	return matrix
}
