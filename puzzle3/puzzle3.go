package puzzle3

import (
	"strings"
	"strconv"
	"fmt"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

const TreeChar = "#"

type Puzzle3 struct {}

type Map struct {
	Rows []MapRow
	Height int
	Width int
}

type MapRow struct {
	Columns []MapEntry
}

type MapEntry struct {
	Char string
}

func (e *MapEntry) IsTree() bool {
	if e.Char == TreeChar {
		return true
	}
	return false
}

func (m *Map) GetPosition(x int, y int) MapEntry {
	x = x % m.Width
	return m.Rows[y].Columns[x]
}

func (m *Map) CountTrees(iterX int, iterY int) int {
	y := 0
	countTrees := 0
	for x := 0; x < 1000000; x += iterX {
		pos := m.GetPosition(x, y)
		if pos.IsTree() {
			countTrees++
		}
		y += iterY
		if (y >= m.Height) {
			break
		}
	}
	return countTrees
}

func (p Puzzle3) ParseMap(input []string) (*Map) {
	ret := Map{}
	for _, line := range input {
		ret.Height++
		row := MapRow{}
		chars := strings.Split(line, "")
		ret.Width = len(chars)
		for _, char := range chars {
			e := MapEntry{
				Char: char,
			}
			row.Columns = append(row.Columns, e)
		}
		ret.Rows = append(ret.Rows, row)
	}
	return &ret
}

func (p Puzzle3) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return output, err
	}
	inputMap := p.ParseMap(i)
	countTrees := inputMap.CountTrees(3, 1)
	output.Text = "There are " + strconv.Itoa(countTrees) + " trees"
	return output, nil
}

func (p Puzzle3) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		return output, err
	}
	inputMap := p.ParseMap(i)
	total := inputMap.CountTrees(1, 1)
	total *= inputMap.CountTrees(3, 1)
	total *= inputMap.CountTrees(5, 1)
	total *= inputMap.CountTrees(7, 1)
	total *= inputMap.CountTrees(1, 2)
	output.Text = "Multiplication of total # of trees " + strconv.Itoa(total)
	return output, nil
}