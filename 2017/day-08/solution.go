package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var max int
	var maxKey string
	registers := map[string]int{}
	for s.Scan() {
		// eg: c dec -10 if a >= 1
		//<=
		//!=
		parts := strings.Split(s.Text(), " ")
		key := parts[0]
		op := parts[1]
		v := registers[key]
		num, err := strconv.Atoi(parts[2])
		if err != nil {
			panic(fmt.Errorf("failed to parse num %s: %s", parts[2], err))
		}

		condKey := parts[4]
		condV := registers[condKey]
		condOp := parts[5]
		condNumRaw, err := strconv.Atoi(parts[6])
		if err != nil {
			panic(fmt.Errorf("failed to parse condNum %s: %s", parts[6], err))
		}

		//fmt.Printf("%d %s %d?\n", condV, condOp, condNumRaw)
		// this could be cleaned up into an OpFactory(condOp) -> func opFunc(left, right) bool
		if !((condOp == ">=" && condV >= condNumRaw) ||
			(condOp == "<=" && condV <= condNumRaw) ||
			(condOp == "!=" && condV != condNumRaw) ||
			(condOp == "<" && condV < condNumRaw) ||
			(condOp == ">" && condV > condNumRaw) ||
			(condOp == "==" && condV == condNumRaw)) {
			//fmt.Printf("\tno\n")

			continue
		}
		//fmt.Printf("\tyes\n")
		if op == "inc" {
			v += num
		} else if op == "dec" {
			v -= num
		} else {
			panic(fmt.Errorf("Unknown op %s", op))
		}
		//fmt.Printf("%s is now %d\n", key, v)
		registers[key] = v

		if v > max {
			max = v
			maxKey = key
		}
	}

	fmt.Printf("Max EVER was %s with %d\n", maxKey, max)
	max = 0
	maxKey = ""
	for k, v := range registers {
		if v > max {
			max = v
			maxKey = k
		}
	}
	fmt.Printf("Max was %s with %d\n", maxKey, max)
}
