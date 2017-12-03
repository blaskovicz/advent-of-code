package main

import (
	"fmt"
	"math"
)

var data = 265149

//var data = 11

func main() {
	partOne()
}

func partOne() {
	width := 3
	min := 1
	for {
		max := int(math.Pow(float64(width), 2))
		if data > min && data <= max {
			// at this point, we are width+offset away from 1
			// since we've calculated the bounds of our outer ring
			fmt.Printf("width=%d (%d, %d]\n", width, min, max)
			var i int
			for cell := min + 1; cell <= max; cell++ {
				if cell == data {
					for i := 0; i < 4; i++ {
						rowStart := min + 1 + (width-1)*i
						rowEnd := rowStart + width - 2
						if cell < rowStart || cell > rowEnd {
							continue
						}
						mid := (rowEnd + rowStart) / 2
						fmt.Printf("[%d < %d < %d]\n", rowStart, mid, rowEnd)
						inRingDistance := int(math.Abs(float64(data - mid)))
						overRingDistance := width / 2

						/*
							width=515 (263169, 265225]
							[264712 < 264968 < 265225]
							Distance: 438
						*/
						fmt.Printf("Distance: %d\n", inRingDistance+overRingDistance)
						break
					}
					break
				}
				i++
			}
			break
		}
		min = max
		width += 2
	}
}
