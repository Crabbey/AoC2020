package puzzle8

import (
	"fmt"
	
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle8 struct {

}

func (p Puzzle8) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}
	VM := common.NewVM(i)
	var positions []int
MainLoop:
	for {
		for _, x := range positions {
			if x == VM.Position {
				fmt.Printf("Position %v is in list already, exiting\n", x)
				break MainLoop
			}
		}
		positions = append(positions, VM.Position)
		VM.Step()
	}
	output.Text = fmt.Sprintf("Value %v is in the accumulator", VM.Accumulator)
	return output, nil
}

func (p Puzzle8) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}
	spew.Dump(i)
	return output, nil
}