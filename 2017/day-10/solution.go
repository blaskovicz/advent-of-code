package main

import "fmt"

type solver struct {
}

func reverse(input []int, inputLen, startIndex, count int) {
	//fmt.Printf("input.rotate=%#v, rotate=%d, index=%d\n", input, count, startIndex)
	for i := 0; i < count/2; i++ {
		left := (startIndex + i) % inputLen
		right := (startIndex + count - 1 - i) % inputLen
		input[left], input[right] = input[right], input[left]
		//fmt.Printf("\t[%d] %#v\n", i, input)
	}
}

func (s *solver) knotHash(input, lengths []int) {
	inputIndex := 0
	inputLen := len(input)
	var skip int
	for _, length := range lengths {
		reverse(input, inputLen, inputIndex, length)
		inputIndex = inputIndex + (length+skip)%inputLen
		if inputIndex >= inputLen {
			inputIndex = inputIndex % inputLen
		}

		skip++
		//fmt.Printf("\tinput.next=%#v, skip=%d, index.next=%d\n\n", input, skip, inputIndex)
	}
}

func newSolver() *solver {
	return &solver{}
}

var lengths = []int{227, 169, 3, 166, 246, 201, 0, 47, 1, 255, 2, 254, 96, 3, 97, 144}

func main() {
	list := make([]int, 256, 256)
	for i := 0; i < 256; i++ {
		list[i] = i
	}
	newSolver().knotHash(list, lengths)
	fmt.Printf("answer: %d\n", list[0]*list[1])
}
