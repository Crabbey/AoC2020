package puzzle12

import (
	"strconv"
	"strings"
	"fmt"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println
var _ = strconv.Atoi
var _ = strings.Split

type Puzzle12 struct {

}

type Ship struct {
	Facing int
	Xpos int
	Ypos int
	WaypointX int
	WaypointY int
}

func (s *Ship) FixFacing() {
	if s.Facing >= 360 {
		s.Facing = s.Facing % 360
	}
	if s.Facing < 0 {
		s.Facing += 360
	}
}

func (s *Ship) Traverse(line string) {
	command := line[0:1]
	value, _ := strconv.Atoi(line[1:])
	switch (command) {
	case "N":
		s.Ypos += value
	case "S":
		s.Ypos -= value
	case "E":
		s.Xpos += value
	case "W":
		s.Xpos -= value
	case "L":
		s.Facing -= value
		s.FixFacing()
	case "R":
		s.Facing += value
		s.FixFacing()
	case "F":
		if s.Facing == 0 {
			s.Traverse(fmt.Sprintf("N%v", value))
		} else if s.Facing == 90 {
			s.Traverse(fmt.Sprintf("E%v", value))
		} else if s.Facing == 180 {
			s.Traverse(fmt.Sprintf("S%v", value))
		} else if s.Facing == 270 {
			s.Traverse(fmt.Sprintf("W%v", value))
		} else {
			spew.Dump(s)
		}
	}
}

func (s *Ship) TraverseWaypoint(line string) {
	command := line[0:1]
	value, _ := strconv.Atoi(line[1:])
	switch (command) {
	case "N":
		s.WaypointY += value
	case "S":
		s.WaypointY -= value
	case "E":
		s.WaypointX += value
	case "W":
		s.WaypointX -= value
	case "L":
		for x := 0; x < value / 90; x++ {
			xtmp := s.WaypointX
			s.WaypointX = s.WaypointY * -1
			s.WaypointY = xtmp
		}
	case "R":
		for x := 0; x < value / 90; x++ {
			xtmp := s.WaypointX
			s.WaypointX = s.WaypointY
			s.WaypointY = xtmp * -1
		}
	case "F":
		for x := 0; x < value; x++ {
			s.Xpos += s.WaypointX
			s.Ypos += s.WaypointY			
		}
	}
}

func (s *Ship) TraversePath(lines []string) {
	for _, line := range lines {
		s.Traverse(line)
	}
}

func (s *Ship) TraverseWaypointPath(lines []string) {
	for _, line := range lines {
		s.TraverseWaypoint(line)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *Ship) ManhattanDistance() int {
	return Abs(s.Xpos) + Abs(s.Ypos)
}

func (p Puzzle12) NewShip() *Ship {
	return &Ship{
		Facing: 90,
	}
}

func (p Puzzle12) Part1(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	Ship := p.NewShip()
	Ship.TraversePath(i)
	output.Text = fmt.Sprintf("MD: %v", Ship.ManhattanDistance())
	return output, nil
}

func (p Puzzle12) Part2(input common.AoCInput) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return nil, err
	}
	output := common.NewSolution(input, "")
	Ship := p.NewShip()
	Ship.WaypointX = 10
	Ship.WaypointY = 1
	Ship.TraverseWaypointPath(i)
	output.Text = fmt.Sprintf("MD: %v", Ship.ManhattanDistance())
	return output, nil
}