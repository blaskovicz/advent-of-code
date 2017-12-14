package main

import (
	"fmt"
	"strconv"
	"strings"
)

// depth: range

/*var data = `
0: 3
1: 2
4: 4
6: 4
`*/

var data = `
0: 5
1: 2
2: 3
4: 4
6: 8
8: 4
10: 6
12: 6
14: 8
16: 6
18: 6
20: 12
22: 14
24: 8
26: 8
28: 9
30: 8
32: 8
34: 12
36: 10
38: 12
40: 12
44: 14
46: 12
48: 10
50: 12
52: 12
54: 12
56: 14
58: 12
60: 14
62: 14
64: 14
66: 14
68: 17
70: 12
72: 14
76: 14
78: 14
80: 14
82: 18
84: 14
88: 20
`

type scanner struct {
	Range int
}

// func (s *scanner) caughtAtTime(t int) bool {

// }

// 1-based index where 1 == top row, and range == last row
// assumes that we start at position 1.
// returns the direction (true for inc, false for dec)
func (s *scanner) positionAtTime(t int) (int, bool) {
	if s.Range < 1 {
		panic(fmt.Errorf("invalid scanner Range %d", s.Range))
	}
	pos := 1
	inc := true

	for i := 0; i <= t; i++ {
		if inc {
			pos++
			if pos == s.Range {
				inc = false
			}
		} else {
			pos--
			if pos == 1 {
				inc = true
			}
		}
	}
	return pos, inc
}

func main() {
	// The severity of getting caught on a layer is equal to its depth multiplied by its range.

	// If a scanner moves into the top of its layer while you are there, you are not caught:
	// it doesn't have time to notice you before you leave.

	// The packet will travel along the top of each layer, and it moves at one layer per picosecond.
	// Each picosecond, the packet moves one layer forward (its first move takes it into layer 0)
	scanners := map[int]*scanner{}
	for _, scannerLine := range strings.Split(strings.TrimSpace(data), "\n") {
		parts := strings.Split(scannerLine, ": ")

		d, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Errorf("Bad depth for %s: %s", parts[0], err))
		}

		r, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(fmt.Errorf("Bad range for %s: %s", parts[1], err))
		}

		s := &scanner{r}
		scanners[d] = s
	}

	var tOffset int
	for {
		var gotem bool
		for d, s := range scanners {
			if (d+tOffset)%(2*s.Range-2) == 0 {
				gotem = true
				break
			}
		}

		if gotem {
			tOffset++
			//fmt.Printf("Offset is now %d\n", tOffset)
		} else {
			break
		}
	}

	fmt.Printf("offset time is %d\n", tOffset)
}
