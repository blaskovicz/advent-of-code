package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type testInput struct {
	t string
	s int
	g int
}

func TestSolution(t *testing.T) {
	for _, input := range []testInput{
		testInput{"<>", 0, 0},                             // 0 garbage characters.
		testInput{"<random characters>", 0, 17},           // 17 garbage characters.
		testInput{"<<<<>", 0, 3},                          // 3 garbage characters.
		testInput{"<{!>}>", 0, 2},                         // 2 garbage characters.
		testInput{"<!!>", 0, 0},                           // 0 garbage characters.
		testInput{"<!!!>>", 0, 0},                         // 0 garbage characters.
		testInput{"<{o\"i!a,<{i<a>", 0, 10},               // 10 garbage characters.
		testInput{"{}", 1, 0},                             // score of 1
		testInput{"{{{}}}", 6, 0},                         // score of 1 + 2 + 3 = 6.
		testInput{"{{},{}}", 5, 0},                        // score of 1 + 2 + 2 = 5.
		testInput{"{{{},{},{{}}}}", 16, 0},                // score of 1 + 2 + 3 + 3 + 3 + 4 = 16.
		testInput{"{<a>,<a>,<a>,<a>}", 1, 4},              // score of 1.
		testInput{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},  // score of 1 + 2 + 2 + 2 + 2 = 9.
		testInput{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},  //score of 1 + 2 + 2 + 2 + 2 = 9.
		testInput{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17}, //score of 1 + 2 = 3.
	} {
		p := newParser()
		require.NotNil(t, p)
		groupCount, garbageCount := p.tokenize(input.t)
		require.Equal(t, input.s, groupCount, "Failed to correctly count groups for %#v", input.t)
		require.Equal(t, input.g, garbageCount, "Failed to correctly count garbage for %#v", input.t)
	}
}
