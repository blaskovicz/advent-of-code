package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	r int  // range
	t int  // time
	d bool // direction
	p int  // position
}

func TestScanner(t *testing.T) {

	for _, input := range []tc{
		tc{r: 4, t: 0, d: true, p: 2},
		tc{r: 4, t: 1, d: true, p: 3},
		tc{r: 4, t: 2, d: false, p: 4},
		tc{r: 4, t: 3, d: false, p: 3},
		tc{r: 4, t: 4, d: false, p: 2},
		tc{r: 4, t: 5, d: true, p: 1},
		tc{r: 4, t: 6, d: true, p: 2},
		tc{r: 4, t: 7, d: true, p: 3},
		tc{r: 4, t: 8, d: false, p: 4},
		tc{r: 4, t: 9, d: false, p: 3},
		tc{r: 4, t: 10, d: false, p: 2},
		tc{r: 4, t: 11, d: true, p: 1},
	} {
		s := &scanner{input.r}
		p, d := s.positionAtTime(input.t)
		assert.Equal(t, input.p, p, "positions differ for %#v", input)
		assert.Equal(t, input.d, d, "directions differ for %#v", input)
	}
}
