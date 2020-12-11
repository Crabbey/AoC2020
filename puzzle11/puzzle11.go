package puzzle11

import (
	"strings"
	"fmt"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump

type Puzzle11 struct {

}

type Grid struct {
	Row map[int]GridRow
	Seats []*Position
}

type GridRow struct {
	Column map[int]*Position
}

type Position struct {
	IsSeat bool
	Xpos int
	Ypos int
	DirectlyAdjacent []*Position
	NearestAdjacent []*Position
	IsSeatOccupied bool
	IsSeatChanging bool
	Tolerance int
}

func (p *Position) FindNearest(grid Grid, x, y int) {
	relx := 0
	rely := 0
	for {
		relx += x
		rely += y
		if _, ok := grid.Row[p.Ypos + rely]; !ok {
			break
		}
		if _, ok2 := grid.Row[p.Ypos + rely].Column[p.Xpos + relx]; !ok2 {
			break
		}
		if grid.Row[p.Ypos + rely].Column[p.Xpos + relx].IsSeat {
			p.NearestAdjacent = append(p.NearestAdjacent, grid.Row[p.Ypos + rely].Column[p.Xpos + relx])
			break
		}
	}
}

func (p *Position) FindAdjacent(grid Grid) {
	if _, ok := grid.Row[p.Ypos-1]; ok {
		if topleft, ok2 := grid.Row[p.Ypos-1].Column[p.Xpos-1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, topleft)
		}
		if top, ok2 := grid.Row[p.Ypos-1].Column[p.Xpos]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, top)
		}
		if topright, ok2 := grid.Row[p.Ypos-1].Column[p.Xpos+1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, topright)
		}
	}
	if _, ok := grid.Row[p.Ypos]; ok {
		if left, ok2 := grid.Row[p.Ypos].Column[p.Xpos-1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, left)
		}
		if topright, ok2 := grid.Row[p.Ypos].Column[p.Xpos+1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, topright)
		}
	}
	if _, ok := grid.Row[p.Ypos+1]; ok {
		if bottomleft, ok2 := grid.Row[p.Ypos+1].Column[p.Xpos-1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, bottomleft)
		}
		if bottom, ok2 := grid.Row[p.Ypos+1].Column[p.Xpos]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, bottom)
		}
		if bottomright, ok2 := grid.Row[p.Ypos+1].Column[p.Xpos+1]; ok2 {
			p.DirectlyAdjacent = append(p.DirectlyAdjacent, bottomright)
		}
	}
	p.FindNearest(grid, -1, -1)
	p.FindNearest(grid, -1, 0)
	p.FindNearest(grid, -1, 1)
	p.FindNearest(grid, 0, -1)
	p.FindNearest(grid, 0, 1)
	p.FindNearest(grid, 1, -1)
	p.FindNearest(grid, 1, 0)
	p.FindNearest(grid, 1, 1)
}

func (p *Position) Iterate(mode string) bool {
	stable := true
	countAdjacentOccupied := 0
	adjacents := p.DirectlyAdjacent
	if (mode == "near") {
		adjacents = p.NearestAdjacent
	}
	for _, adj := range adjacents {
		if (adj.IsSeatOccupied && !adj.IsSeatChanging) || (!adj.IsSeatOccupied && adj.IsSeatChanging) {
			countAdjacentOccupied++
		}
	}
	if !p.IsSeatOccupied && countAdjacentOccupied == 0 {
		p.IsSeatOccupied = true
		p.IsSeatChanging = true
		stable = false
	} else if p.IsSeatOccupied && countAdjacentOccupied >= p.Tolerance {
		p.IsSeatOccupied = false
		p.IsSeatChanging = true
		stable = false
	}
	return stable
}

func (g *Grid) Print() {
	for y := 0; y < len(g.Row); y++ {
		row := g.Row[y]	
		for x := 0; x < len(row.Column)	; x++ {
			pos := row.Column[x]
			if !pos.IsSeat {
				fmt.Print(".")
				continue
			}
			if pos.IsSeatOccupied {
				fmt.Print("#")
				continue
			}
			fmt.Print("L")
		}
		fmt.Println("")
	}
}

func (g *Grid) CountOccupiedSeats() int {
	count := 0
	for _, s := range g.Seats {
		if s.IsSeatOccupied {
			count++
		}
	}
	return count
}

func (g *Grid) IterateOnce(mode string) bool {
	stable := true
	for _, seat := range g.Seats {
		stability := seat.Iterate(mode)
		if stable && !stability {
			stable = false
		}
	}
	for _, seat := range g.Seats {
		seat.IsSeatChanging = false
	}
	return stable
}

func (g *Grid) Iterate(mode string) int {
	count := 0
	for {
		if g.IterateOnce(mode) {
			break
		}
		count++
	}
	return count
}

func NewGrid(input []string, tolerance int) Grid {
	grid := Grid{
		Row: make(map[int]GridRow),
	}
	for y, line := range input {
		parts := strings.Split(line, "")
		for x, seat := range parts {
			if _, ok := grid.Row[y]; !ok {
				grid.Row[y] = GridRow{
					Column: make(map[int]*Position),
				}
			}

			position := Position{
				Xpos: x,
				Ypos: y,
				Tolerance: tolerance,
			}
			if seat == "L" {
				position.IsSeat = true
				grid.Seats = append(grid.Seats, &position)
			}
			grid.Row[y].Column[x] = &position
		}
	}
	for _, seat := range grid.Seats {
		seat.FindAdjacent(grid)
	}
	return grid
}

func (p Puzzle11) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	grid := NewGrid(i, 4)
	runs := grid.Iterate("")
	seats := grid.CountOccupiedSeats()
	output.Text = fmt.Sprintf("%v runs produces %v occupied seats to stability", runs, seats)
	return output, nil
}

func (p Puzzle11) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	grid := NewGrid(i, 5)
	runs := grid.Iterate("near")
	seats := grid.CountOccupiedSeats()
	output.Text = fmt.Sprintf("%v runs produces %v occupied seats to stability", runs, seats)
	return output, nil
}