package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEdges(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []Edge
		err   error
	}{
		{input: [][]int{}, want: nil, err: ErrEmptyMatrix},
		{input: [][]int{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		}, want: nil, err: ErrManyVertices},
		{input: [][]int{
			{0, 1, 1, 1, 1, 0, 0, 0, 0, 1},
			{0, 1, 0, 0, 1, 1, 0, 1, 0, 1},
			{1, 0, 1, 1, 0, 1, 0, 0, 0, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{1, 0, 0, 1, 0, 1, 1, 0, 0, 1},
			{1, 1, 1, 0, 1, 1, 1, 1, 0, 0},
			{0, 1, 0, 0, 0, 1, 1, 1, 1, 0},
		}, want: []Edge{
			{[]int{2, 4, 5}},
			{[]int{0, 1, 3, 5, 6}},
			{[]int{0, 2, 3, 5}},
			{[]int{0, 2, 3, 4}},
			{[]int{0, 1, 3, 5}},
			{[]int{1, 2, 3, 4, 5, 6}},
			{[]int{3, 4, 5, 6}},
			{[]int{1, 3, 5, 6}},
			{[]int{6}},
			{[]int{0, 1, 2, 4}},
		}, err: nil}}

	for _, tt := range tests {
		actual, err := GetEdges(tt.input)
		assert.ErrorIs(t, err, tt.err)
		assert.Equal(t, tt.want, actual)
	}
}
