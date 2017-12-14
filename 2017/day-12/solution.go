package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const delim = " <-> "
const delim2 = ", "

type node struct {
	data int
	conn map[*node]interface{}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := bufio.NewScanner(f)
	var tree *node
	nodes := map[int]*node{}
	for r.Scan() {
		parts := strings.SplitN(r.Text(), delim, 2)
		parent := parts[0]
		children := strings.Split(parts[1], delim2)
		// parent

		nP, err := strconv.Atoi(parent)
		if err != nil {
			panic(err)
		}
		vP, ok := nodes[nP]
		if !ok {
			vP = &node{nP, map[*node]interface{}{}}
			nodes[nP] = vP
			if tree == nil {
				tree = vP
			}
		}

		// conns
		{
			for _, c := range children {
				n, err := strconv.Atoi(c)
				if err != nil {
					panic(err)
				}
				v, ok := nodes[n]
				if !ok {
					v = &node{n, map[*node]interface{}{vP: struct{}{}}}
					nodes[n] = v
				} else {
					if _, ok := v.conn[vP]; !ok {
						v.conn[vP] = struct{}{}
					}
				}

				if _, ok := vP.conn[v]; !ok {
					vP.conn[v] = struct{}{}
				}
			}

		}
	}

	// traverse the tree, keeping track of visitted
	var visittedCount int
	visitted := map[int]interface{}{}
	stack := []*node{tree}
	var groupCount int
	for {
		groupCount++
		for {
			if len(stack) == 0 {
				break
			}
			next := stack[0]
			stack = stack[1:]
			if _, ok := visitted[next.data]; ok {
				continue
			}
			if groupCount == 1 {
				visittedCount++
			}
			visitted[next.data] = struct{}{}
			delete(nodes, next.data)
			// visit all neighbors
			for k, _ := range next.conn {
				stack = append(stack, k)
			}
		}
		if len(nodes) == 0 {
			break
		} else {
			for _, v := range nodes {
				stack = append(stack, v)
				break
			}
		}
	}
	fmt.Printf("visitted %d in group 1 and %d groups total\n", visittedCount, groupCount)
}
