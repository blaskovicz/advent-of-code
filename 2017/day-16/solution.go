package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func spin(data []rune, dataLen int, index map[rune]int, count int) {
	// if we spin all items, we're back where we started
	count = count % dataLen
	temp := append([]rune{}, data[dataLen-count:dataLen]...)
	temp = append(temp, data[:dataLen-count]...)
	for i, d := range temp {
		data[i] = d
	}

	// rebuild the index
	for i, r := range data {
		index[r] = i
	}
}
func exchange(data []rune, dataLen int, index map[rune]int, a, b int) {
	dataLen--
	if a > dataLen || b > dataLen || a < 0 || b < 0 {
		panic(fmt.Errorf("Invalid exchange request based on state: %d<->%d (%#v)", a, b, data))
	}
	aRune := data[a]
	bRune := data[b]
	data[a], data[b] = data[b], data[a]
	index[aRune], index[bRune] = b, a
}
func partner(data []rune, dataLen int, index map[rune]int, a, b rune) {
	aIndex, aOk := index[a]
	bIndex, bOk := index[b]
	if !aOk || !bOk {
		panic(fmt.Errorf("Invalid partner request based on state: %s<->%s (%#v)", string(a), string(b), index))
	}
	data[aIndex], data[bIndex] = data[bIndex], data[aIndex]
	index[a], index[b] = index[b], index[a]
}

func parseSlashArgNums(s string) (int, int) {
	slash := strings.IndexByte(s, '/')
	if slash == -1 || slash == 0 {
		panic("Invalid slash args")
	}
	xA, err := strconv.Atoi(string(s[1:slash]))
	if err != nil {
		panic(fmt.Errorf("Invalid slash arg A: %s", err))
	}
	xB, err := strconv.Atoi(string(s[slash+1:]))
	if err != nil {
		panic(fmt.Errorf("Invalid slash arg B: %s", err))
	}
	return xA, xB
}

func parseSlashArgs(s string) (rune, rune) {
	slash := strings.IndexByte(s, '/')
	if slash == -1 || slash == 0 {
		panic("Invalid slash args")
	}
	// if we need to support a range here, switch to strings or array
	xA := rune(s[1])
	xB := rune(s[slash+1])
	return xA, xB
}

const delim = ','

func initialState() ([]rune, int, map[rune]int, io.ReadCloser) {
	index := map[rune]int{}
	state := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}
	l := len(state)
	// build initial index
	for i, r := range state {
		index[r] = i
	}

	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %s", err))
	}
	return state, l, index, f
}

// func initialState() ([]rune, int, map[rune]int, io.ReadCloser) {
// 	return []rune{'a', 'b', 'c', 'd', 'e'}, 5, map[rune]int{'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 5}, ioutil.NopCloser(strings.NewReader("s1,x3/4,pe/b"))
// }

func main() {
	state, l, index, f := initialState()
	defer f.Close()
	r := bufio.NewReader(f)
	ops := []string{}
	for {
		// get next operation
		s, err := r.ReadString(delim)
		sLast := len(s) - 1
		if s[sLast] == ',' {
			s = string(s[:sLast])
		}
		eof := err == io.EOF
		if err != nil && !eof {
			panic(err)
		} else if sLast == 0 {
			break
		}
		ops = append(ops, s)
		if eof {
			break
		}
	}
	var previousStates = map[string]int{}
	var ignoreLoop bool
	for i := 0; i < 1000000000; i++ {
		for _, s := range ops {
			// do operation
			switch s[0] {
			case 's':
				spinCount, err := strconv.Atoi(string(s[1:]))
				if err != nil {
					panic(fmt.Errorf("Bad spin count: %s", err))
				}
				spin(state, l, index, spinCount)
			case 'x':
				xA, xB := parseSlashArgNums(s)
				exchange(state, l, index, xA, xB)
			case 'p':
				xA, xB := parseSlashArgs(s)
				partner(state, l, index, xA, xB)
			default:
				panic(fmt.Errorf("Unknown directive: %s", s))
			}
			//fmt.Printf("[%s] -> %s\n", s, string(state))
		}
		if !ignoreLoop {
			ss := string(state)
			if after, ok := previousStates[ss]; ok {
				fmt.Printf("Saw %s again after %d (%d)\n", ss, i, after)
				// we're now looping, so short circuit us
				loop := i - after
				for i+loop < 1000000000 {
					i += loop
				}
				ignoreLoop = true
				//os.Exit(1)
			} else {
				previousStates[ss] = i
			}
		}
	}
	fmt.Printf("%#v\n", string(state))
}
