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
}
