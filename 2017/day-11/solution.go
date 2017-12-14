package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

/*

translate cube coordinates to cube distance:

https://www.redblobgames.com/grids/hexagons/

          (-z)        (+x)
           \          /
            \ n      /
          nw +------+ ne
            /        \
(+y) ------+          +-------- (-y)
            \        /
         sw  +------+ se
  	        / s      \
           /          \
          (-x)         (+z)
*/
type cubeCoordinate struct {
	x float64
	y float64
	z float64
}

func (c *cubeCoordinate) advance(d string) {
	switch d {
	case "n":
		c.z--
		c.x++
	case "ne":
		c.x++
		c.y--
	case "se":
		c.y--
		c.z++
	case "s":
		c.z++
		c.x--
	case "sw":
		c.x--
		c.y++
	case "nw":
		c.y++
		c.z--
	default:
		panic(fmt.Errorf("Invalid direction, %s", d))
	}
}

func cubeDistance(a, b *cubeCoordinate) float64 {
	return math.Max(math.Max(math.Abs(a.x-b.x), math.Abs(a.y-b.y)), math.Abs(a.z-b.z))
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to read input.txt: %s", err))
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("Failed to read input.txt contents: %s", err))
	}
	var maxDistance float64
	start := &cubeCoordinate{}
	last := &cubeCoordinate{}
	for _, p := range bytes.Split(b, []byte(",")) {
		last.advance(string(p))
		if newMax := cubeDistance(start, last); newMax > maxDistance {
			maxDistance = newMax
		}
	}
	fmt.Printf("Answer %.1f and furthest was %.1f\n", cubeDistance(start, last), maxDistance)
}
