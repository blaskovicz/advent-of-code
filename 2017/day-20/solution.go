package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type particle struct {
	// position
	pX float64
	pY float64
	pZ float64

	// velocity
	vX float64
	vY float64
	vZ float64

	// acceleration
	aX float64
	aY float64
	aZ float64
}

func (p *particle) distance() float64 {
	return math.Abs(p.pX) + math.Abs(p.pY) + math.Abs(p.pZ)
}

func (p *particle) accelerate() {
	p.vX += p.aX
	p.vY += p.aY
	p.vZ += p.aZ

	p.pX += p.vX
	p.pY += p.vY
	p.pZ += p.vZ
}

func parseXYZ(s string) (x float64, y float64, z float64, err error) {
	parts := strings.SplitN(strings.TrimSpace(s), ",", 3)
	if len(parts) != 3 {
		err = fmt.Errorf("invalid particle single: %#v", parts)
		return
	}
	x, err = strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return
	}
	y, err = strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return
	}
	z, err = strconv.ParseFloat(parts[2], 64)
	return
}

func parseParticles() ([]*particle, error) {
	f, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	p := []*particle{}
	for s.Scan() {
		// p=<-886,2415,982>, v=<-129,344,142>, a=<9,-29,-12>
		parts := strings.SplitN(strings.TrimSpace(s.Text()), ", ", 3)
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid particle parts: %#v", parts)
		}
		var nextParticle particle
		for _, part := range parts {
			// p=<-886,2415,982> -> -886,2415,982
			// fmt.Printf("%s\n", part[3:len(part)-1])
			x, y, z, err := parseXYZ(part[3 : len(part)-1])
			if err != nil {
				return nil, err
			}
			switch part[0] {
			case 'p':
				nextParticle.pX = x
				nextParticle.pY = y
				nextParticle.pZ = z
			case 'v':
				nextParticle.vX = x
				nextParticle.vY = y
				nextParticle.vZ = z
			case 'a':
				nextParticle.aX = x
				nextParticle.aY = y
				nextParticle.aZ = z
			default:
				return nil, fmt.Errorf("invalid particle type: %#v", part)
			}

		}
		p = append(p, &nextParticle)
	}
	return p, s.Err()

}

func main() {
	particles, err := parseParticles()
	if err != nil {
		panic(err)
	}
	var minParticle *int
	var minParticleDistance *float64
	last := 100000
	for i := 0; i <= last; i++ {
		for j := range particles {
			p := particles[j]
			p.accelerate()
			if i == last {
				m := p.distance()
				if minParticle == nil || m < *minParticleDistance {
					temp := j
					minParticleDistance = &m
					minParticle = &temp
				}
				fmt.Printf("[%d] particle %d p(%.0f,%.0f,%.0f) v(%.0f,%.0f,%.0f) a(%.0f,%.0f,%.0f) -> %.0f\n", i, j, p.pX, p.pY, p.pZ, p.vX, p.vY, p.vZ, p.aX, p.aY, p.aZ, m)
			}
		}
	}
	fmt.Printf("min particle: %d (%.0f)\n", *minParticle, *minParticleDistance)
}
