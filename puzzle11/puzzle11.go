package puzzle11

import (
	
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle11 struct {

}

func (p Puzzle11) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	spew.Dump(i)
	return output, nil
}

func (p Puzzle11) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	spew.Dump(i)
	return output, nil
}