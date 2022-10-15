package graph

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ErrEmptyMatrix  = errors.New("passed empty matrix")
	ErrManyVertices = errors.New("count of vertices must be lower then 8")
)

type Edge struct {
	Vertices []int
}

// Gets incidence matrix and returns list of edges
func GetEdges(matrix [][]int) ([]Edge, error) {
	if len(matrix) == 0 {
		return nil, ErrEmptyMatrix
	}
	if len(matrix) > 8 {
		return nil, ErrManyVertices
	}

	edges := make([]Edge, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != 0 {
				edges[j].Vertices = append(edges[j].Vertices, i)
			}
		}
	}

	return edges, nil
}

// Generate incidence matrix where columns = edges and rows = vertices
func GenMatrix(vertices, edges int) [][]int {
	matrix := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		matrix[i] = make([]int, edges)
		for j := 0; j < edges; j++ {
			matrix[i][j] = rand.Int() % 2
		}
	}

	return matrix
}

func ShowMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
