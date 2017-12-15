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

func main() {
	var bitsUsed uint32
	for r := 0; r < 128; r++ {
		s := newSolver()
		base16Hash := s.denseKnotHash(fmt.Sprintf("%s-%d", data, r))
		for c := 0; c < 32; c++ {
			hexString := string(base16Hash[c])
			hexInt, err := strconv.ParseInt(hexString, 16, 8)
			if err != nil {
				panic(fmt.Errorf("failed to parse hex string %s: %s", hexString, err))
			}
			if hexInt&1 == 1 {
				bitsUsed++
			}
			if hexInt&2 == 2 {
				bitsUsed++
			}
			if hexInt&4 == 4 {
				bitsUsed++
			}
			if hexInt&8 == 8 {
				bitsUsed++
			}
		}
	}
	fmt.Printf("Bits used: %d\n", bitsUsed)
}
