package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testInput struct {
	input   []int
	lengths []int
	count   int
}

type denseTestInput struct {
	input  string
	result string
}

func TestSolution(t *testing.T) {
	for _, ti := range []testInput{
		testInput{[]int{0, 1, 2, 3, 4}, []int{3, 4, 1, 5}, 12},
		//testInput{[]int{0, 1, 2, 3, 4, 5}, []int{3, 4, 1, 5}, 12},
	} {
		s := newSolver()
		require.NotNil(t, s, "solver didn't get instantiated")
		s.knotHash(ti.input, ti.lengths)
		actualCount := ti.input[0] * ti.input[1]
		require.Equal(t, ti.count, actualCount, "Multiplication of first 2 result numbers incorrect")
	}

	for _, ti := range []denseTestInput{
		denseTestInput{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		denseTestInput{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		denseTestInput{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		denseTestInput{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	} {
		s := newSolver()
		require.NotNil(t, s, "solver didn't get instantiated")
		actual := s.denseKnotHash(ti.input)
		require.Equal(t, ti.result, actual, "Dense hash incorrect")
	}
}
