package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var data = `4	1	15	12	0	9	9	5	5	8	7	3	14	5	12	3`

//var data = `0 2 7 0`

func main() {
	spaceRe := regexp.MustCompile("\\s+")
	parts := spaceRe.Split(data, -1)
	v := []int{}
	// ideally, use a heap to keep track of the highest count index
	for _, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil {
			panic(err)
		}
		v = append(v, val)
	}
	//fmt.Printf("%v\n", v)

	valCount := len(v)
	seen := map[string]interface{}{}
	var iterations uint64
	var target string
	for {
		// find max
		var max int
		var index int
		for i, val := range v {
			if val > max {
				max = val
				index = i
			}
		}

		// distribute max's blocks in a clockwise fashion
		v[index] = 0
		for ; max > 0; max-- {
			index = (index + 1) % valCount
			v[index]++
		}

		// serialize our iteration's state or bail
		s := []string{}
		for _, v := range v {
			s = append(s, strconv.Itoa(v))
		}
		s2 := strings.Join(s, ",")
		//fmt.Printf("%s\n", s2)

		if _, ok := seen[s2]; ok {
			iterations++
			if target == "" {
				target = s2
				iterations = 0
			} else if s2 == target {
				break
			}
		} else {
			seen[s2] = struct{}{}
			iterations++
		}
	}

	// 2392
	fmt.Printf("Iterations: %d\n", iterations)
}
