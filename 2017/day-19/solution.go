package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"unicode"
)

type direction int

const (
	left direction = iota
	right
	up
	down
	unknown

	pathVertical   = '|'
	pathJuction    = '+'
	pathHorizontal = '-'
)

var opposite = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type grid struct {
	x           int
	y           int
	rows        int
	cols        int
	plane       [][]byte
	dir         direction
	pathLetters []byte
	steps       uint
}

func newGrid(allData []byte) *grid {
	g := &grid{
		x:           0,
		y:           0,
		plane:       bytes.Split(allData, []byte("\n")),
		pathLetters: []byte{},
		dir:         down,
		steps:       1,
	}

	// let's go ahead and assume we have at least one row
	// and we have equal length cols
	g.rows = len(g.plane)
	g.cols = len(g.plane[0])

	// find the start point
	var found bool
	for i := 0; i < g.cols; i++ {
		if g.plane[0][i] != pathVertical {
			continue
		}
		g.x = i
		found = true
	}

	if !found {
		panic("Could not find start point")
	}

	return g
}

// Only turn left or right when there's no other option.
// In addition, someone has left letters on the line; these also don't change its direction
func (g *grid) traverse() {
PATH:
	for {
		current := g.point(g.x, g.y)
		if current == nil {
			panic(fmt.Errorf("Current point isn't traversable?! (%d,%d)", g.x, g.y))
		}
		if unicode.IsLetter(rune(*current)) {
			g.pathLetters = append(g.pathLetters, *current)
		}

		// add all direction choices, preferring the current direction and
		// skipping the opposite direction (eg: if we go down, don't go up)
		choiceCount := 1
		choices := []direction{g.dir}
		for _, c := range []direction{up, down, left, right} {
			if c != g.dir && c != opposite[g.dir] {
				choices = append(choices, c)
				choiceCount++
			}
		}

	NEXT_CHOICE:
		for {
			if choiceCount == 0 {
				break PATH // nowhere left to go
			}
			c := choices[0]
			if choiceCount > 1 {
				choices = choices[1:]
			} else {
				choices = []direction{}
			}
			choiceCount--

			switch c {
			case down:
				next := g.point(g.x, g.y+1)
				if next != nil {
					current = next
					g.y++
					g.dir = down
					break NEXT_CHOICE
				}
			case up:
				next := g.point(g.x, g.y-1)
				if next != nil {
					current = next
					g.y--
					g.dir = up
					break NEXT_CHOICE
				}
			case left:
				next := g.point(g.x-1, g.y)
				if next != nil {
					current = next
					g.x--
					g.dir = left
					break NEXT_CHOICE
				}
			case right:
				next := g.point(g.x+1, g.y)
				if next != nil {
					current = next
					g.x++
					g.dir = right
					break NEXT_CHOICE
				}
			}
		}
		g.steps++
		fmt.Printf("(%d,%d) %s\n", g.x, g.y, string(*current))
	}
}

// is this a point we can even go to?
func (g *grid) point(x, y int) *byte {
	if x >= g.cols || x < 0 || y >= g.rows || y < 0 {
		return nil
	}
	point := g.plane[y][x]
	if point != pathHorizontal && point != pathVertical && point != pathJuction && !unicode.IsLetter(rune(point)) {
		return nil
	}
	return &point
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	g := newGrid(b)
	g.traverse()
	fmt.Printf("Path: %#v; Steps %d\n", string(g.pathLetters), g.steps)
}
