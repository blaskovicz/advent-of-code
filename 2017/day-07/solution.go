package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	pointingToCount := map[string]int{}
	for s.Scan() {
		// eg: ugml (68) -> gyxo, ebii, jptl
		parts := strings.Split(s.Text(), " ")
		partCount := len(parts)
		if partCount < 3 {
			continue
		}
		k := parts[0]
		//fmt.Printf("k=%s\n", k)
		if _, ok := pointingToCount[k]; !ok {
			pointingToCount[k] = 0
		}
		for i := 3; i < partCount; i++ {
			v := strings.Replace(parts[i], ",", "", -1)
			//fmt.Printf("\t-%s\n", v)
			pointingToCount[v] = 1
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	//fmt.Printf("%#v", pointingToCount)
	for k, v := range pointingToCount {
		if v == 0 {
			fmt.Printf("%s\n", k)
			break
		}
	}
}
