package main

import (
	"fmt"
	"strconv"
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
		a := strconv.FormatInt(genA.next(), 2)
		aLen := len(a) - 1
		b := strconv.FormatInt(genB.next(), 2)
		bLen := len(b) - 1
		matched := true
		if a != b {
			for j := 0; j < 16; j++ {
				var aVal byte
				var bVal byte
				aIndex := aLen - j
				bIndex := bLen - j
				if aIndex < 0 {
					aVal = '0'
				} else {
					aVal = a[aIndex]
				}
				if bIndex < 0 {
					bVal = '0'
				} else {
					bVal = b[bIndex]
				}
				if aVal != bVal {
					matched = false
					break
				}
			}
		}
		if matched {
			//fmt.Printf("[%d] == %v | %v ==\n", i, a, b)
			matchedTimes++
		}
	}
	fmt.Printf("Matched times: %d", matchedTimes)
}
