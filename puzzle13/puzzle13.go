package puzzle13

import (
	
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle13 struct {

}

func (p Puzzle13) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
	return output, err
	}
	spew.Dump(i)
	return output, nil
}

func (p Puzzle13) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
	return output, err
	}
	spew.Dump(i)
	return output, nil
}