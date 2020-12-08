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
	VM := common.NewVM(i)
IncLoop:
	for repl := 0; repl <= len(i) ; repl++ {
		VM.Reset()
		VM.LoadMemspace(i)
		positions := []int{}
		if VM.MemSpace[repl].Command == "nop" {
			VM.MemSpace[repl].Command = "jmp"
		} else if VM.MemSpace[repl].Command == "jmp" {
			VM.MemSpace[repl].Command = "nop"
		}
		for !VM.Terminated {
			for _, x := range positions {
				if x == VM.Position {
					continue IncLoop
				}
			}
			positions = append(positions, VM.Position)
			VM.Step()
		}
		// This will only be reached once a VM terminates by finishing it's instruction set
		break
	}
	output.Text = fmt.Sprintf("Value %v is in the accumulator", VM.Accumulator)
	return output, nil
}