package puzzle9

import (
	"strconv"
	"fmt"
	
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle9 struct {

}

func (p Puzzle9) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	lines, err := input.Read()
	if err != nil {
		spew.Dump(lines)
		return output, err
	}
	var ints []int
	for _, i := range lines {
		x, _ := strconv.Atoi(i)
		ints = append(ints, x)
	}
	var broken int
	MainLoop:
	for x, i := range ints {
		if x < 25 {
			// 25 line preamble
			continue;
		}
		for y := x - 25; y < x; y++ {
			for z := x - 25; z < x; z++ {
				if ints[y] + ints[z] == i {
					// fmt.Printf("%v + %v = %v \n", ints[y], ints[x], i)
					continue MainLoop
				}
			}
		}
		fmt.Printf("Could not find a sum for %v", i)
		broken = i
		break
	}
	output.Text = fmt.Sprintf("%v is the first broken number", broken)
	return output, nil
}

func (p Puzzle9) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	lines, err := input.Read()
	if err != nil {
		spew.Dump(lines)
		return output, err
	}
	sumTo := 26134589 // Answer from part 1
	var ints []int
	for _, i := range lines {
		x, _ := strconv.Atoi(i)
		ints = append(ints, x)
	}
	var firstPos int
	var lastPos int
	MainLoop:
	for x, i := range ints {
		y := x + 1
		total := i
		for {
			total += ints[y]
			if total > sumTo {
				break;
			}
			if total == sumTo {
				firstPos = x
				lastPos = y
				break MainLoop
			}
			y++
		}
	}
	var high int
	var low = ints[firstPos]
	for x := firstPos; x <= lastPos; x++ {
		if ints[x] < low {
			low = ints[x]
		}
		if ints[x] > high {
			high = ints[x]
		}
	}
	output.Text = fmt.Sprintf("%v is the sum of the highest and lowest values in the broken set", high + low)
	return output, nil
}