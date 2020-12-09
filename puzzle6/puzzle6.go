package puzzle6

import (
	"strings"
	"fmt"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle6 struct {

}

type Group struct {
	Answers map[string]int
	NumPeople int
}

func (p Puzzle6) LoadResponses(input []string) []Group {
	g := Group{
		Answers: make(map[string]int),
	}
	var ret []Group
	for _, line := range input {
		if line == "" {
			ret = append(ret, g)
			g = Group{
				Answers: make(map[string]int),
			}
			continue
		}
		g.NumPeople++
		chars := strings.Split(line, "")
		for _, char := range chars {
			g.Answers[char]++
		}
	}
	ret = append(ret, g)
	return ret
}

func (p Puzzle6) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	groups := p.LoadResponses(i)
	count := 0
	for _, group := range groups {
		for _, _ = range group.Answers {
			count++
		}
	}
	output.Text = fmt.Sprintf("Count: %v\n", count)
	return output, nil
}

func (p Puzzle6) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	groups := p.LoadResponses(i)
	count := 0
	for _, group := range groups {
		for _, total := range group.Answers {
			if total == group.NumPeople {
				count++
			}
		}
	}
	output.Text = fmt.Sprintf("Count: %v\n", count)
	return output, nil
}