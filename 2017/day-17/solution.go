package main

import "fmt"

type node struct {
	data int
	next *node
}

var steps = 356

//var steps = 3

func main() {
	target := 2017
	head := &node{0, nil}
	current := head
	nextNum := 1
	for nextNum <= target {
		// walk forward n steps
		walk := steps % nextNum
		for walk != 0 {
			if current.next == nil {
				current = head
			} else {
				current = current.next
			}
			walk--
		}

		// insert new node after current
		newNode := &node{data: nextNum, next: current.next}
		current.next = newNode

		// update insert position to node
		current = newNode
		nextNum++
	}

	// print result
	if current.next != nil {
		fmt.Printf("%d\n", current.next.data)
	} else {
		fmt.Printf("%d\n", head.data)
	}
}
