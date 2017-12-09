package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type parser struct {
	tokens     []byte
	groupDepth int
}

func newParser() *parser {
	return &parser{tokens: []byte{}}
}

func (p *parser) push(token byte) {
	if token == groupOpen {
		p.groupDepth++
	}
	p.tokens = append(p.tokens, token)
}

func (p *parser) peek() *byte {
	tokenCount := len(p.tokens)
	if tokenCount == 0 {
		return nil
	}
	return &p.tokens[tokenCount-1]

}

func (p *parser) pop() *byte {
	tokenCount := len(p.tokens)
	if tokenCount == 0 {
		return nil
	}
	t := p.tokens[tokenCount-1]
	if tokenCount > 1 {
		p.tokens = p.tokens[:tokenCount-1]
	} else {
		p.tokens = []byte{}
	}

	if t == groupOpen {
		p.groupDepth--
	}

	return &t
}

const (
	stateReset = iota
	stateGarbage
	stateIgnore
	stateNormal

	garbageOpen  = '<'
	garbageClose = '>'
	groupOpen    = '{'
	groupClose   = '}'
	ignore       = '!'
)

func (p *parser) tokenize(s string) (count, garbagePile int) {
	state := stateReset
	for _, t := range s {
		if state == stateReset {
			last := p.peek()
			if last == nil || *last == groupOpen {
				state = stateNormal
			} else if *last == garbageOpen {
				state = stateGarbage
			}
		}

		if state == stateIgnore {
			state = stateReset
			continue
		} else if state == stateGarbage {
			if t == garbageClose {
				p.pop()
				state = stateReset
			} else if t == ignore {
				state = stateIgnore
			} else {
				garbagePile++
			}
		} else if state == stateNormal {
			if t == garbageOpen {
				p.push(garbageOpen)
				state = stateGarbage
			} else if t == groupOpen {
				p.push(groupOpen)
			} else if t == groupClose {
				p.pop()
				count += p.groupDepth + 1
			}
			//
		} else {
			panic("rekt")
		}

	}
	return
}

func main() {
	f, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Errorf("Failed to open input: %s", err))
	}
	defer f.Close()
	input, err := ioutil.ReadAll(f)
	if err != nil {
		panic(fmt.Errorf("Failed to parse input: %s", err))
	}
	groupCount, garbageCount := newParser().tokenize(string(input))
	fmt.Printf("%d group count\n", groupCount)
	fmt.Printf("%d garbage count\n", garbageCount)
}
