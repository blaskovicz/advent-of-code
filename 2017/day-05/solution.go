package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	list := []int{}
	for s.Scan() {
		nextInt, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(fmt.Errorf("Failed to parse %s: %s", s.Text(), err))
		}
		list = append(list, nextInt)
	}

	var moves uint64
	for index := 0; index < len(list) && index > -1; {
		atIndex := list[index]
		if atIndex >= 3 {
			// part 2
			list[index]--
		} else {
			// part 1
			list[index]++
		}
		index += atIndex
		moves++
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("moves: %d\n", moves)
}
