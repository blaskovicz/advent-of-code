package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	regs := map[string]int{}
	ops := [][]string{}
	for _, op := range bytes.Split(b, []byte("\n")) {
		ops = append(ops, strings.Split(string(op), " "))
	}

	opc := len(ops)
	i := 0
	song := []int{}
	for i < opc && i > -1 {
		o := ops[i]
		cmd := o[0]
		reg := o[1]
		var value *int
		if len(o) > 2 && o[2] != "" {
			regOrVal := o[2]
			if val, err := strconv.Atoi(regOrVal); err != nil {
				// letter -> register value
				temp := regs[regOrVal]
				value = &temp
			} else {
				// number -> raw value
				value = &val
			}
		}
		switch cmd {
		case "set":
			regs[reg] = *value
		case "add":
			regs[reg] += *value
		case "mul":
			regs[reg] *= *value
		case "mod":
			regs[reg] %= *value
		case "snd":
			song = append(song, regs[reg])
		case "rcv":
			regVal := regs[reg]
			if songc := len(song); songc != 0 && regVal != 0 {
				// potential disco
				regs[reg] = song[songc-1]
				fmt.Printf("rcv %#v\n", regs[reg])
				os.Exit(0)
			}
		case "jgz":
			regVal := regs[reg]
			if regVal != 0 {
				i += *value
				continue
			}
		default:
			panic(fmt.Errorf("Unknown operation %#v", o))
		}
		i++
	}

	fmt.Printf("%s", string(b))
}
