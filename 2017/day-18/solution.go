package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

	ops := [][]string{}
	for _, op := range bytes.Split(b, []byte("\n")) {
		ops = append(ops, strings.Split(string(op), " "))
	}

	opc := len(ops)

	regs0 := map[string]int{"p": 0}
	regs1 := map[string]int{"p": 1}
	bus0 := make(chan int)
	bus1 := make(chan int)
	err0 := make(chan error)
	err1 := make(chan error)
	go runProg(ops, opc, regs0, bus1, bus0, err0)
	go runProg(ops, opc, regs1, bus0, bus1, err1)

	err = <-err0
	fmt.Printf("0 exitted (%v)\n", err)
	err = <-err1
	fmt.Printf("1 exitted (%v)\n", err)
}

func regValueOrRawInt(regs map[string]int, regOrVal string) int {
	val, err := strconv.Atoi(regOrVal)
	if err != nil {
		// letter -> register value
		return regs[regOrVal]
	}
	// number -> raw value
	return val
}

func runProg(ops [][]string, opc int, regs map[string]int, rcv <-chan int, snd chan<- int, errChan chan<- error) {
	progID := regs["p"]
	var sendCount int

	i := 0
	var sendLock sync.Mutex
	song := []int{}
	go func() {
		for {
			sendLock.Lock()
			if songc := len(song); songc != 0 {
				s := song[0]
				if songc > 1 {
					song = song[1:]
				} else {
					song = []int{}
				}
				sendCount++
				//fmt.Printf("[%d] snd %#v (x%d)\n", progID, s, sendCount)
				sendLock.Unlock()
				snd <- s
			} else {
				sendLock.Unlock()
			}
			time.Sleep(1 * time.Millisecond)
		}
	}()

	for i < opc && i > -1 {
		o := ops[i]
		//fmt.Printf("[%d] (%#v)\n", progID, o)
		cmd := o[0]
		reg := o[1]
		var value *int
		if len(o) > 2 && o[2] != "" {
			temp := regValueOrRawInt(regs, o[2])
			value = &temp
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
			valToSend := regValueOrRawInt(regs, reg)
			sendLock.Lock()
			fmt.Printf("[%d] snd[%s] %#v\n", progID, reg, valToSend)
			song = append(song, valToSend)
			sendLock.Unlock()
		case "rcv":
			regs[reg] = <-rcv
			fmt.Printf("[%d] rcv[%s] %#v\n", progID, reg, regs[reg])
			// go func() {
			// 	errChan <- nil
			// }()
			// return
		case "jgz":
			regVal := regValueOrRawInt(regs, reg)
			if regVal != 0 {
				i += *value
				continue
			}
		default:
			panic(fmt.Errorf("Unknown operation %#v", o))
		}
		i++
	}
}
