package main

import (
	"fmt"
	"strconv"
)

// copied from day 10: "knot hash"

type solver struct {
	inputIndex int
	skip       int
}

func reverse(input []int, inputLen, startIndex, count int) {
	for i := 0; i < count/2; i++ {
		left := (startIndex + i) % inputLen
		right := (startIndex + count - 1 - i) % inputLen
		input[left], input[right] = input[right], input[left]
	}
}

func (s *solver) denseKnotHash(data string) string {
	lengths := asciiData(data)

	input := make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		input[i] = i
	}
	for i := 0; i < 64; i++ {
		s.knotHash(input, lengths)
	}

	var denseHash string
	for i := 0; i < 16; i++ {
		offset := i * 16
		end := offset + 16
		denseHash += fmt.Sprintf("%02x", xorBlock(input[offset:end]))
	}
	return denseHash
}

func (s *solver) knotHash(input, lengths []int) {
	inputLen := len(input)
	for _, length := range lengths {
		reverse(input, inputLen, s.inputIndex, length)
		s.inputIndex = s.inputIndex + (length+s.skip)%inputLen
		if s.inputIndex >= inputLen {
			s.inputIndex = s.inputIndex % inputLen
		}
		s.skip++
	}
}

func xorBlock(input []int) int {
	result := input[0]
	for i := 1; i < len(input); i++ {
		result ^= input[i]
	}
	return result
}

func newSolver() *solver {
	return &solver{}
}

func asciiData(s string) []int {
	lengths := []int{}
	for _, b := range []byte(s) {
		lengths = append(lengths, int(b))
	}
	lengths = append(lengths, 17, 31, 73, 47, 23)
	return lengths
}

// end copied from day 10: "knot hash"
//var data = "flqrgnkx"
var data = "hfdlxzhv"

// rows[row][col]
func countGroups(rows map[int]map[int]interface{}) uint32 {
	//fmt.Printf("%#v\n", rows)
	var groupCount uint32
	cellsToProcess := [][]int{}

	getNextCell := func() []int {
		for r, cols := range rows {
			for c := range cols {
				return []int{r, c}
			}
		}
		return nil
	}

	isCellSet := func(r, c int) bool {
		if cols, ok := rows[r]; ok {
			if _, ok := cols[c]; ok {
				return true
			}
		}
		return false
	}

	for {
		// no more cells to process
		if len(cellsToProcess) == 0 {
			// if the set has more items, start a new group
			if c := getNextCell(); c != nil {
				//fmt.Println("Starting new cell group")
				groupCount++
				cellsToProcess = append(cellsToProcess, c)
			} else {
				// otherwise we're done
				//fmt.Println("Done")
				break
			}
		}
		cell := cellsToProcess[0]
		if len(cellsToProcess) == 1 {
			cellsToProcess = [][]int{}
		} else {
			cellsToProcess = cellsToProcess[1:]
		}

		r := cell[0]
		c := cell[1]
		//fmt.Printf("processing neighbor row=%d,col=%d\n", r, c)
		delete(rows[r], c)
		if len(rows[r]) == 0 {
			delete(rows, r)
		}
		// add cell neighbors to process if they are 1s

		// up
		if isCellSet(r-1, c) {
			cellsToProcess = append(cellsToProcess, []int{r - 1, c})
		}
		// down
		if isCellSet(r+1, c) {
			cellsToProcess = append(cellsToProcess, []int{r + 1, c})
		}
		// left
		if isCellSet(r, c-1) {
			cellsToProcess = append(cellsToProcess, []int{r, c - 1})
		}
		// right
		if isCellSet(r, c+1) {
			cellsToProcess = append(cellsToProcess, []int{r, c + 1})
		}
	}

	return groupCount
}

func main() {
	var bitsUsed uint32
	rows := map[int]map[int]interface{}{}

	bits := []int64{8, 4, 2, 1}
	for r := 0; r < 128; r++ {
		s := newSolver()
		base16Hash := s.denseKnotHash(fmt.Sprintf("%s-%d", data, r))
		for c := 0; c < 32; c++ {
			hexString := string(base16Hash[c])
			hexInt, err := strconv.ParseInt(hexString, 16, 8)
			if err != nil {
				panic(fmt.Errorf("failed to parse hex string %s: %s", hexString, err))
			}
			i := c * 4
			for _, b := range bits {
				if hexInt&b == b {
					bitsUsed++
					c, ok := rows[r]
					if !ok {
						c = map[int]interface{}{}
						rows[r] = c
					}
					c[i] = struct{}{}
				}
				i++
			}
		}
	}

	// another, more complicated, way to do this would have been to keep
	// a rollowing state of the rows but it's more difficult to calculate than a traditional
	// visitted search

	fmt.Printf("Bits used: %d; Groups found: %d\n", bitsUsed, countGroups(rows))
}
