package puzzle2

import (
	"strconv"
	"strings"
	// "regexp"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle2 struct {

}

type Puz2Entry struct {
	MaxCount int
	MinCount int
	Letter string
	Input string
}

func (e *Puz2Entry) Part1Test() bool {
	count := strings.Count(e.Input, e.Letter)
	if count >= e.MinCount && count <= e.MaxCount {
		return true
	}
	return false
}

func (e *Puz2Entry) Part2Test() bool {
	bytes := []byte(e.Input)
	bytes = append([]byte("="), bytes...)
	letters := []byte(e.Letter)
	letter := letters[0]
	matchcount := 0
	if bytes[e.MinCount] == letter {
		matchcount++
	}
	if bytes[e.MaxCount] == letter {
		matchcount++
	}
	if matchcount == 1 {
		return true
	}
	return false
}

func (p Puzzle2) Parse(input string) *Puz2Entry {
	entry := Puz2Entry{}
	parts := strings.Split(input, " ")
	counts := strings.Split(parts[0], "-")
	entry.MinCount, _ = strconv.Atoi(counts[0])
	entry.MaxCount, _ = strconv.Atoi(counts[1])
	entry.Letter = strings.Replace(parts[1], ":", "", -1)
	entry.Input = parts[2]
	return &entry
}

func (p Puzzle2) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	count := 0
	for _, line := range i {
		entry := p.Parse(line)
		if entry.Part1Test() {
			count++
		}
	}
	output.Text = strconv.Itoa(count)
	return output, nil
}

func (p Puzzle2) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return nil, err
	}
	output := common.NewSolution(input, "")
	count := 0
	for _, line := range i {
		entry := p.Parse(line)
		if entry.Part2Test() {
			count++
		}
	}
	output.Text = strconv.Itoa(count)
	return output, nil
}