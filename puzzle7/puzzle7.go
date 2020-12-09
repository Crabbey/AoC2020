package puzzle7

import (
	"fmt"
	"strings"
	"strconv"
	"regexp"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle7 struct {

}

type Bag struct {
	Name string
	Contents []string
	Parents []*Bag
	Children []ChildBag
}

type ChildBag struct {
	Count int
	Bag *Bag
}

var allNumbers = regexp.MustCompile(`\d+ ?`)

func ParseInput(input []string) map[string]*Bag {
	bags := make(map[string]*Bag)
	for _, line := range input {
		bag := Bag{}
		line = strings.ReplaceAll(line, "bags", "bag")
		line = strings.ReplaceAll(line, ".", "")
		line = strings.ReplaceAll(line, ".", "")
		typeOfBag := strings.Split(line, " contain ")
		bag.Name = typeOfBag[0]
		childBags := strings.Split(typeOfBag[1], ", ")
		for _, childBag := range childBags {
			if childBag == "no other bag" {
				break
			}
			bag.Contents = append(bag.Contents, childBag)
		}
		bags[bag.Name] = &bag
	}
	for _, bag := range bags {
		for _, childBag := range bag.Contents {
			childBagName := allNumbers.ReplaceAllString(childBag, "")
			x := bags[childBagName]
			x.Parents = append(x.Parents, bag)
			bags[childBagName] = x
			bagCount := 1
			testCount, err := strconv.Atoi(childBag[0:1])
			if (err == nil) {
				bagCount = testCount
			}
			bag.Children = append(bag.Children, ChildBag{Count: bagCount, Bag: x})
		}
	}
	return bags
}

func (b *Bag) CountChildrenRecursive() (int) {
	total := 0
	for _, child := range b.Children {
		total += child.Count
		total += child.Bag.CountChildrenRecursive() * child.Count
	}
	return total
}

func (b *Bag) CountParentsRecursive(countedBags []string) (int, []string) {
	total := 0
ParentLoop:
	for _, parent := range b.Parents {
		for _, cb := range countedBags {
			if cb == parent.Name {
				continue ParentLoop
			}
		}
		total += 1
		countedBags = append(countedBags, parent.Name)
		x, newCountedBags := parent.CountParentsRecursive(countedBags)
		total += x
		countedBags = newCountedBags
	}
	return total, countedBags
}

func (p Puzzle7) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	bags := ParseInput(i)
	var emptyBags []string
	total, _ := bags["shiny gold bag"].CountParentsRecursive(emptyBags)
	output.Text = fmt.Sprintf("There are %v bags that can contain a shiny gold bag eventually", total)
	return output, nil
}

func (p Puzzle7) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	bags := ParseInput(i)
	total := bags["shiny gold bag"].CountChildrenRecursive()
	output.Text = fmt.Sprintf("There are %v bags contained within your shiny gold bag", total)

	return output, nil
}