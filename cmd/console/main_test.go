package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	args     []string
	vertices int
	edges    int
}

func TestMain(t *testing.T) {
	tests := []struct {
		input        input
		want         [][]int
		panic        bool
		panicMessage string
	}{
		{input: input{
			args:     []string{},
			vertices: 10,
			edges:    10,
		}, want: nil, panic: true, panicMessage: "not enough args"},
		{input: input{
			args:     []string{"a", "b"},
			vertices: 1,
			edges:    1,
		}, want: nil, panic: true, panicMessage: `strconv.Atoi: parsing "a": invalid syntax`},
		{input: input{
			args: []string{
				"1", "0",
				"0", "1",
			},
			vertices: 2,
			edges:    2,
		}, want: [][]int{
			{1, 0},
			{0, 1},
		}, panic: false},
	}

	for _, tt := range tests {
		if tt.panic {
			assert.PanicsWithValue(t, tt.panicMessage, func() {
				readMatrix(tt.input.args, tt.input.vertices, tt.input.edges)
			})
			continue
		}

		assert.Equal(t, tt.want, readMatrix(tt.input.args, tt.input.vertices, tt.input.edges))
	}
}
