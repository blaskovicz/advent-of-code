package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var validCount uint
	for s.Scan() {
		wordSet := map[string]interface{}{}
		parts := strings.Split(s.Text(), " ")
		valid := true
		for _, p := range parts {
			// part 1, don't anagram; part 2: sorted anagram
			p2 := anagram(p)
			if _, ok := wordSet[p2]; !ok {
				wordSet[p2] = struct{}{}
			} else {
				valid = false
				break
			}
		}
		if valid {
			validCount++
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("combos: %d\n", validCount)
}

type LexSort []byte

func (a LexSort) Len() int           { return len(a) }
func (a LexSort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a LexSort) Less(i, j int) bool { return a[i] < a[j] }

func anagram(word string) string {
	newWord := []byte{}
	for i := len(word) - 1; i >= 0; i-- {
		newWord = append(newWord, word[i])
	}
	sort.Sort(LexSort(newWord))
	return string(newWord)
}
