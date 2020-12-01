package main

import (
	"fmt"
	"sort"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/crabbey/aoc2020/puzzle1"
	// "github.com/crabbey/aoc2020/puzzle2"
	// "github.com/crabbey/aoc2020/puzzle3"
	// "github.com/crabbey/aoc2020/puzzle4"
	// "github.com/crabbey/aoc2020/puzzle5"
	// "github.com/crabbey/aoc2020/puzzle6"
	// "github.com/crabbey/aoc2020/puzzle7"
	// "github.com/crabbey/aoc2020/puzzle8"
	// "github.com/crabbey/aoc2020/puzzle9"
	// "github.com/crabbey/aoc2020/puzzle10"
	// "github.com/crabbey/aoc2020/puzzle11"
	// "github.com/crabbey/aoc2020/puzzle12"
	// "github.com/crabbey/aoc2020/puzzle13"
	// "github.com/crabbey/aoc2020/puzzle14"
	// "github.com/crabbey/aoc2020/puzzle15"
	// "github.com/crabbey/aoc2020/puzzle16"
	// "github.com/crabbey/aoc2020/puzzle17"
	// "github.com/crabbey/aoc2020/puzzle18"
	// "github.com/crabbey/aoc2020/puzzle19"
	// "github.com/crabbey/aoc2020/puzzle20"
	// "github.com/crabbey/aoc2020/puzzle21"
	// "github.com/crabbey/aoc2020/puzzle22"
	// "github.com/crabbey/aoc2020/puzzle23"
	// "github.com/crabbey/aoc2020/puzzle23"
	// "github.com/crabbey/aoc2020/puzzle24"
	// "github.com/crabbey/aoc2020/puzzle25"
	"github.com/crabbey/aoc2020/common"

	"github.com/urfave/cli/v2"
)

var _ = spew.Dump

const implemented = 1

func main() {
	app := cli.NewApp()
	app.Name = "AoC2020 Runner"
	app.EnableBashCompletion = true
	app.CommandNotFound = func(context *cli.Context, cmd string) {
		fmt.Printf("ERROR: Unknown command '%s'\n", cmd)
	}

	app.Commands = []*cli.Command{
		&cmdAllPuzzles,
		&cmdSinglePuzzle,
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

var cmdAllPuzzles = cli.Command{
	Name:  "all",
	Action: func(c *cli.Context) error {
		for i := 1; i <= implemented; i++ {
			puzzleid := strconv.Itoa(i)
			x := CallPuzzle(c, puzzleid)
			if x != nil {
				return x
			}
		}
		return nil
	},
}

var cmdSinglePuzzle = cli.Command{
	Name:  "puzzle",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "inputname",
        	Aliases: []string{"i"},
			Usage:   "Input file for puzzle",
			EnvVars: []string{"inputname"},
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "part, p",
			Usage:   "Puzzle part",
			EnvVars: []string{"part"},
			Value:   "",
		},
	},
	Action: func(c *cli.Context) error {
		puzzleid := c.Args().Get(0)
		partid := c.String("partid")
		if partid == "" {			
			return CallPuzzle(c, puzzleid)
		}
		CallPuzzlePart(c, puzzleid, partid)
		return nil
	},
}

func GetInput(c *cli.Context, puzzleid string) common.AoCInput {
	iname := c.String("inputname")
	ret := common.AoCInput{
		Path: "puzzle"+puzzleid,
		InputFile: iname,
	}
	if iname == "" {
		ret.InputFile = "input.txt"
	}
	return ret
}

func CallPuzzlePart(c *cli.Context, puzzleid string, partid string) (*common.AoCSolution, error) {
	input := GetInput(c, puzzleid)
	solution := common.NewSolution(puzzleid, partid)

	var puzzle common.AoCPuzzle
	switch puzzleid {
	case "1":
		puzzle = puzzle1.Puzzle1{}
	default:
		return nil, fmt.Errorf("Unknown puzzle %v", puzzleid)
	}
	switch partid {
	case "1":
		return puzzle.Part1(input, solution)
	case "2":
		return puzzle.Part2(input, solution)
	default:
		return nil, fmt.Errorf("Unknown part id %v", partid)
	}
	return nil, nil
}

func CallPuzzle(c *cli.Context, puzzleid string) error {
	solution, err := CallPuzzlePart(c, puzzleid, "1")
	if err != nil {
		return err
	}
	solution.Print()
	solution2, err := CallPuzzlePart(c, puzzleid, "2")
	if err != nil {
		return err
	}
	solution2.Print()
	return nil
}