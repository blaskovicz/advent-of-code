package main

import "fmt"

type node struct {
	data int
	next *node
}

var steps = 356

//var steps = 3

func main() {
	target := 50000000
	currentPosition := 0
	afterZero := -1
	for i := 1; i <= target; i++ {
		// walk forward n steps from current position, normalized to array len
		walk := (currentPosition + steps) % i

		// inserting at index 0
		if walk == 0 {
			afterZero = i
		}

		// update current position to newly inserted node index
		currentPosition = walk + 1
	}
	fmt.Printf("After zero: %d\n", afterZero)
}

func main1() {
	//target := 2017
	target := 50000000
	zero := &node{0, nil}
	lastZero := 0
	head := zero
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
		// if lastZero != zero.next.data {
		// 	fmt.Printf("[%d] -> %d\n", nextNum-1, zero.next.data)
		// 	lastZero = zero.next.data
		// }
	}

	// print result
	if zero.next != nil {
		fmt.Printf("%d\n", zero.next.data)
	} else {
		fmt.Printf("%d\n", head.data)
	}
}
