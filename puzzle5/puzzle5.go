package puzzle5

import (
	"strings"
	"math"
	"fmt"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

type Puzzle5 struct {

}

type Seat struct {
	Row int
	Col int
	Uid int
}

func (p Puzzle5) LoadSeats(input []string) []Seat {
	var ret []Seat
	rowMax := 127
	colMax := 7
	for _, line := range input {
		seat := Seat{}
		char := strings.Split(line, "")
		for x := 0; x < 7; x++ {
			if string(char[x]) == "B" {
				plusone := float64(x + 1)
				seat.Row += ((rowMax + 1) / int(math.Pow(2, plusone))) 
			}
			continue
		}
		for x := 7; x < 10; x++ {
			if string(char[x]) == "R" {
				// -6 because we start at 7
				plusone := float64(x - 6)
				seat.Col += ((colMax + 1) / int(math.Pow(2, plusone))) 
			}
			continue
		}
		seat.Uid = seat.Row * 8 + seat.Col
		ret = append(ret, seat)
	}
	return ret
}

func (p Puzzle5) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}
	seats := p.LoadSeats(i)
	highest := 0
	for _, seat := range seats {
		if seat.Uid > highest {
			highest = seat.Uid
		}
	}
	output.Text = fmt.Sprintf("Seat UID %v is the highest", highest)
	return output, nil
}

func (p Puzzle5) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}

	seats := p.LoadSeats(i)
	var seatUids []int
	seatUidMap := make(map[int]Seat)
	for _, seat := range seats {
		seatUids = append(seatUids, seat.Uid)
		seatUidMap[seat.Uid] = seat
	}
	for _, uid := range seatUids {
		// If the next one doesn't exist (is vacant) but the one after that exists, it's yours
		if _, ok := seatUidMap[uid + 1] ; !ok {
			if _, ok2 := seatUidMap[uid + 2] ; ok2 {
				output.Text = fmt.Sprintf("Seat UID %v is yours", uid+1)
				return output, nil
			}
		}
	}
	return output, nil
}