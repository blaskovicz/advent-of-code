package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	name     string
	weight   int
	children map[string]*node
}

func newNode(name string, weight int) *node {
	return &node{name: name, weight: weight, children: map[string]*node{}}
}

func weight(w string) int {
	i, err := strconv.Atoi(strings.Replace(strings.Replace(w, "(", "", -1), ")", "", -1))
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	pointingToCount := map[string]int{}
	nodes := map[string]*node{}
	for s.Scan() {
		// eg: ugml (68) -> gyxo, ebii, jptl
		parts := strings.Split(s.Text(), " ")
		partCount := len(parts)

		k := parts[0]
		w := weight(parts[1])
		if n, ok := nodes[k]; !ok {
			nodes[k] = newNode(k, w)
		} else if n.weight != w {
			n.weight = w
		}
		if partCount < 3 {
			continue
		}
		//fmt.Printf("k=%s\n", k)
		if _, ok := pointingToCount[k]; !ok {
			pointingToCount[k] = 0
		}
		for i := 3; i < partCount; i++ {
			v := strings.Replace(parts[i], ",", "", -1)
			//fmt.Printf("\t-%s\n", v)
			pointingToCount[v] = 1
			if _, ok := nodes[v]; !ok {
				nodes[v] = newNode(k, -1)
			}
			if _, ok := nodes[k].children[v]; !ok {
				nodes[k].children[v] = nodes[v]
			}
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	var head *node
	for k, v := range pointingToCount {
		if v == 0 {
			head = nodes[k]
			break
		}
	}
	if head == nil {
		panic("headless horseman")
	}
	//fmt.Printf("%#v", head)
	checkBalance(head)
}

type pair struct {
	n     *node
	count int
}

func checkBalance(n *node) int {
	var count int
	var t int
	towers := []pair{}
	for i, c := range n.children {
		t++
		b := checkBalance(c)
		count += b
		towers = append(towers, pair{n.children[i], b})
	}

	if t != 0 {
		if t > 1 {
			sort.Sort(TowerSort(towers))
			t0 := towers[0]
			tN := towers[t-1]
			if t0.count != tN.count {
				adj := -1 * (tN.count - t0.count)
				panic(fmt.Errorf("Adjust %s(%d) by %d to %d\n", tN.n.name, tN.n.weight, adj, tN.n.weight+adj))
			}
		}

		return n.weight + count
	}
	return n.weight
}

type TowerSort []pair

func (a TowerSort) Len() int           { return len(a) }
func (a TowerSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a TowerSort) Less(i, j int) bool { return a[i].count < a[j].count }
