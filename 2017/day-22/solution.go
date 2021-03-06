package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type direction int
type infectionState int

const (
	left direction = iota
	right
	up
	down

	clean infectionState = iota
	weakened
	infected
	flagged

	cellClean    = '.'
	cellInfected = '#'
)

var nextInfectState = map[infectionState]infectionState{
	clean:    weakened,
	weakened: infected,
	infected: flagged,
	flagged:  clean,
}

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
var opposite = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type grid struct {
	x              int
	y              int
	dir            direction
	infectCount    int
	trackInfection bool
	infected       map[int]map[int]infectionState // x, y
}

/*
- If the current node is infected, it turns to its right. Otherwise, it turns to its left. (Turning is done in-place; the current node does not change.)
- If the current node is clean, it becomes infected. Otherwise, it becomes cleaned. (This is done after the node is considered for the purposes of changing direction.)
- The virus carrier moves forward one node in the direction it is facing.
*/
func (g *grid) burst() {
	currentState := g.infectionState(g.x, g.y)
	switch currentState {
	case clean:
		g.dir = rotateCounterClock[g.dir]
	case weakened:
		// no turn
	case infected:
		g.dir = rotateClock[g.dir]
	case flagged:
		g.dir = opposite[g.dir]
	}

	g.infect(g.x, g.y, nextInfectState[currentState])

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

func (g *grid) infectionState(x, y int) infectionState {
	if ys, ok := g.infected[x]; ok {
		if v, ok := ys[y]; ok {
			return v
		}
	}
	return clean
}

func (g *grid) infect(x, y int, state infectionState) {
	//fmt.Printf("[I] (%d,%d)\n", x, y)
	var gotInfected bool
	if state == clean {
		delete(g.infected[x], y)
		if len(g.infected[x]) == 0 {
			delete(g.infected, x)
		}
	} else {
		if ys, ok := g.infected[x]; ok {
			if v, ok := ys[y]; !ok || v != state {
				ys[y] = state
				if state == infected {
					gotInfected = true
				}
			}
		} else {
			g.infected[x] = map[int]infectionState{
				y: state,
			}
			if state == infected {
				gotInfected = true
			}

		}
	}
	if gotInfected && g.trackInfection {
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
		infected: map[int]map[int]infectionState{},
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
			g.infect(oX, oY, infected)
		}
	}
	g.trackInfection = true
	return g
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	g := newGrid(b)
	for i := 0; i < 10000000; i++ {
		g.burst()
	}
	fmt.Printf("%#v\n", g.infectCount)
	//fmt.Printf("Path: %#v; Steps %d\n", string(g.pathLetters), g.steps)
}
