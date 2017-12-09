package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testInput struct {
	t string
	s int
}

func TestSolution(t *testing.T) {
	for _, input := range []testInput{
		testInput{"{}", 1},                            // score of 1
		testInput{"{{{}}}", 6},                        // score of 1 + 2 + 3 = 6.
		testInput{"{{},{}}", 5},                       // score of 1 + 2 + 2 = 5.
		testInput{"{{{},{},{{}}}}", 16},               // score of 1 + 2 + 3 + 3 + 3 + 4 = 16.
		testInput{"{<a>,<a>,<a>,<a>}", 1},             // score of 1.
		testInput{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9}, // score of 1 + 2 + 2 + 2 + 2 = 9.
		testInput{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9}, //score of 1 + 2 + 2 + 2 + 2 = 9.
		testInput{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3}, //score of 1 + 2 = 3.
	} {
		p := newParser()
		require.NotNil(t, p)
		groupCount := p.tokenize(input.t)
		require.Equal(t, input.s, groupCount, "Failed to correctly count groups for %#v", input)
	}
}
