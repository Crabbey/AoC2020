package puzzle10

import (
	"strconv"
	"fmt"
	"sort"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

type Puzzle10 struct {

}

type Adaptor struct {
	Value int
	NextAdaptors []*Adaptor
	IsEnd bool
	CachedCount int
}

func (a *Adaptor) CountWaysToEnd() int {
	if a.IsEnd {
		return 1
	}
	if a.CachedCount > 0{
		return a.CachedCount
	}
	ret := 0
	for _, c := range a.NextAdaptors {
		ret += c.CountWaysToEnd()
	}
	a.CachedCount = ret
	return ret
}

func (a *Adaptor) FindNextAdaptors(adaptors []*Adaptor) {
	for _, x := range adaptors {
		if x.Value > a.Value {
			if x.Value <= a.Value + 3 {
				a.NextAdaptors = append(a.NextAdaptors, x)
			}
		}
	}
}

func (p Puzzle10) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	var ints []int
	for _, x := range i {
		conv, _ := strconv.Atoi(x)
		ints = append(ints, conv)
	}
	sort.Ints(ints)
	oneless := 1
	threeless := 1
	for p, x := range ints {
		if p == 0 {
			continue
		}
		if x - 1 == ints[p-1] {
			oneless++
		}
		if x - 3 == ints[p-1] {
			threeless++
		}
	}
	output.Text = fmt.Sprintf("%v * %v = %v", oneless, threeless, oneless*threeless)
	return output, nil
}

func (p Puzzle10) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	var adaptors []*Adaptor
	firstAdaptor := Adaptor{Value: 0}
	adaptors = append(adaptors, &firstAdaptor)
	max := 0
	for _, x := range i {
		conv, _ := strconv.Atoi(x)
		if conv > max {
			max = conv
		}
		adaptor := Adaptor{
			Value: conv,
		}
		adaptors = append(adaptors, &adaptor)
	}
	adaptors = append(adaptors, &Adaptor{Value: max + 3, IsEnd: true})
	for _, a := range adaptors {
		a.FindNextAdaptors(adaptors)
	}

	total := firstAdaptor.CountWaysToEnd()
	output.Text = fmt.Sprintf("Found %v pathways", total)
	return output, nil
}