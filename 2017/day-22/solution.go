package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type direction int

const (
	left direction = iota
	right
	up
	down

	cellClean    = '.'
	cellInfected = '#'
)

var rotateClock = map[direction]direction{
	up:    right,
	right: down,
	down:  left,
	left:  up,
}
var rotateCounterClock = map[direction]direction{
	up:    left,
	left:  down,
	down:  right,
	right: up,
}

type grid struct {
	onInfected     bool
	x              int
	y              int
	dir            direction
	infectCount    int
	trackInfection bool
	infected       map[int]map[int]interface{} // x, y
}

/*
- If the current node is infected, it turns to its right. Otherwise, it turns to its left. (Turning is done in-place; the current node does not change.)
- If the current node is clean, it becomes infected. Otherwise, it becomes cleaned. (This is done after the node is considered for the purposes of changing direction.)
- The virus carrier moves forward one node in the direction it is facing.
*/
func (g *grid) burst() {
	if g.isInfected(g.x, g.y) {
		g.dir = rotateClock[g.dir]
		g.clean(g.x, g.y)
	} else {
		g.dir = rotateCounterClock[g.dir]
		g.infect(g.x, g.y)
	}
	switch g.dir {
	case up:
		g.y++
	case down:
		g.y--
	case left:
		g.x--
	case right:
		g.x++
	}
}

func (g *grid) isInfected(x, y int) bool {
	if ys, ok := g.infected[x]; ok {
		if _, ok := ys[y]; ok {
			return true
		}
	}
	return false
}

func (g *grid) clean(x, y int) {
	if ys, ok := g.infected[x]; ok {
		if _, ok := ys[y]; ok {
			delete(ys, y)
		}
		if len(ys) == 0 {
			delete(g.infected, x)
		}
	}
}

func (g *grid) infect(x, y int) {
	//fmt.Printf("[I] (%d,%d)\n", x, y)
	var infected bool
	if ys, ok := g.infected[x]; ok {
		if _, ok := ys[y]; !ok {
			ys[y] = struct{}{}
			infected = true
		}
	} else {
		g.infected[x] = map[int]interface{}{
			y: struct{}{},
		}
		infected = true
	}
	if g.trackInfection && infected {
		g.infectCount++
	}
}

func newGrid(allData []byte) *grid {
	lines := bytes.Split(allData, []byte("\n"))
	rows := len(lines)
	cols := len(lines[0])
	if rows%2 != 1 {
		panic(fmt.Errorf("rows must be odd number"))
	} else if cols%2 != 1 {
		panic(fmt.Errorf("cols must be odd number"))
	}

	g := &grid{
		dir:      up,
		x:        0,
		y:        0,
		infected: map[int]map[int]interface{}{},
	}
	offsetY := rows / 2
	offsetX := cols / 2
	for i := 0; i < rows; i++ {
		oY := offsetY - i
		for j := 0; j < cols; j++ {
			c := lines[i][j]
			if c == cellClean {
				continue
			}
			oX := j - offsetX
			g.infect(oX, oY)
		}
	}

	return g
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	g := newGrid(b)
	g.trackInfection = true
	for i := 0; i < 10000; i++ {
		g.burst()
	}
	fmt.Printf("%#v\n", g.infectCount)
	//fmt.Printf("Path: %#v; Steps %d\n", string(g.pathLetters), g.steps)
}
