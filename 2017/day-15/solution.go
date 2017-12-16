package main

import (
	"fmt"
)

type generator struct {
	seed           int64
	multiplyFactor int64
	modCheck       int64
}

func (g *generator) next() int64 {
	for {
		g.seed = (g.seed * g.multiplyFactor) % 2147483647
		if g.seed%g.modCheck != 0 {
			continue
		}
		break
	}
	return g.seed
}

const seedGenA = 516
const seedGenB = 190

//const seedGenA = 65
//const seedGenB = 8921

func main() {
	genA := generator{seed: seedGenA, multiplyFactor: 16807, modCheck: 4}
	genB := generator{seed: seedGenB, multiplyFactor: 48271, modCheck: 8}
	var matchedTimes uint64
	for i := 0; i < 5000000; i++ {
		a := genA.next()
		b := genB.next()

		// check that the same lower 16 bits are set
		if (a & 0xFFFF) == (b & 0xFFFF) {
			matchedTimes++
		}
	}
	fmt.Printf("Matched times: %d", matchedTimes)
}
