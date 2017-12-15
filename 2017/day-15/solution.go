package main

import (
	"fmt"
	"strconv"
)

type generator struct {
	seed           int64
	multiplyFactor int64
}

func (g *generator) next() int64 {
	g.seed = (g.seed * g.multiplyFactor) % 2147483647
	return g.seed
}

const seedGenA = 516
const seedGenB = 190

//const seedGenA = 65
//const seedGenB = 8921

func main() {
	genA := generator{seed: seedGenA, multiplyFactor: 16807}
	genB := generator{seed: seedGenB, multiplyFactor: 48271}
	var matchedTimes uint64
	for i := 0; i < 40000000; i++ {
		a := strconv.FormatInt(genA.next(), 2)
		aLen := len(a) - 1
		b := strconv.FormatInt(genB.next(), 2)
		bLen := len(b) - 1
		matched := true
		for j := 0; j < 16; j++ {
			if a[aLen-j] != b[bLen-j] {
				matched = false
				break
			}
		}
		if matched {
			//fmt.Printf("[%d] == %v | %v ==\n", i, a, b)
			matchedTimes++
		}
	}
	fmt.Printf("Matched times: %d", matchedTimes)
}
