package puzzle1

import (
	"strconv"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle1 struct {

}

func (p Puzzle1) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	for a, input1 := range i {
		x, _ := strconv.Atoi(input1)
		for b, input2 := range i {
			y, _ := strconv.Atoi(input2)
			if x + y == 2020 && a != b {
				output.Text = strconv.Itoa(x * y)
				return output, nil
			}
		}
	}
	return output, nil
}

func (p Puzzle1) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	for a, input1 := range i {
		x, _ := strconv.Atoi(input1)
		for b, input2 := range i {
			y, _ := strconv.Atoi(input2)
			for c, input3 := range i {
				z, _ := strconv.Atoi(input3)
				if x + y + z == 2020 && a != b && a != c && b != c {
					output.Text = strconv.Itoa(x * y * z)
					return output, nil
				}

			}
		}
	}
	return output, nil
}