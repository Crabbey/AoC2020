package main

import (
	"fmt"
	"sort"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/crabbey/aoc2020/puzzle1"
	"github.com/crabbey/aoc2020/puzzle2"
	"github.com/crabbey/aoc2020/puzzle3"
	"github.com/crabbey/aoc2020/puzzle4"
	"github.com/crabbey/aoc2020/puzzle5"
	"github.com/crabbey/aoc2020/puzzle6"
	"github.com/crabbey/aoc2020/puzzle7"
	"github.com/crabbey/aoc2020/puzzle8"
	"github.com/crabbey/aoc2020/puzzle9"
	"github.com/crabbey/aoc2020/puzzle10"
	"github.com/crabbey/aoc2020/puzzle11"
	"github.com/crabbey/aoc2020/puzzle12"
	"github.com/crabbey/aoc2020/puzzle13"
	"github.com/crabbey/aoc2020/puzzle14"
	"github.com/crabbey/aoc2020/puzzle15"
	"github.com/crabbey/aoc2020/puzzle16"
	"github.com/crabbey/aoc2020/puzzle17"
	"github.com/crabbey/aoc2020/puzzle18"
	"github.com/crabbey/aoc2020/puzzle19"
	"github.com/crabbey/aoc2020/puzzle20"
	"github.com/crabbey/aoc2020/puzzle21"
	"github.com/crabbey/aoc2020/puzzle22"
	"github.com/crabbey/aoc2020/puzzle23"
	"github.com/crabbey/aoc2020/puzzle24"
	"github.com/crabbey/aoc2020/puzzle25"
	"github.com/crabbey/aoc2020/common"

	"github.com/urfave/cli/v2"
)

var _ = spew.Dump

const implemented = 25

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
			Name:    "part",
        	Aliases: []string{"p"},
			Usage:   "Puzzle part",
			EnvVars: []string{"part"},
			Value:   "",
		},
	},
	Action: func(c *cli.Context) error {
		puzzleid := c.Args().Get(0)
		partid := c.String("part")
		if partid == "" {
			return CallPuzzle(c, puzzleid)
		}
		solution, err := CallPuzzlePart(c, puzzleid, partid)
		if err != nil {
			return err
		}
		solution.Print()
		return nil
	},
}

func GetInput(c *cli.Context, puzzleid, partid string) common.AoCInput {
	iname := c.String("inputname")
	ret := common.AoCInput{
		Path: "puzzle"+puzzleid,
		InputFile: iname,
		Puzzle: puzzleid,
		Part: partid,
	}
	if iname == "" {
		ret.InputFile = "input.txt"
	}
	return ret
}

func CallPuzzlePart(c *cli.Context, puzzleid string, partid string) (*common.AoCSolution, error) {
	input := GetInput(c, puzzleid, partid)
	var puzzle common.AoCPuzzle
	switch puzzleid {
	case "1":
		puzzle = puzzle1.Puzzle1{}
	case "2":
		puzzle = puzzle2.Puzzle2{}
	case "3":
		puzzle = puzzle3.Puzzle3{}
	case "4":
		puzzle = puzzle4.Puzzle4{}
	case "5":
		puzzle = puzzle5.Puzzle5{}
	case "6":
		puzzle = puzzle6.Puzzle6{}
	case "7":
		puzzle = puzzle7.Puzzle7{}
	case "8":
		puzzle = puzzle8.Puzzle8{}
	case "9":
		puzzle = puzzle9.Puzzle9{}
	case "10":
		puzzle = puzzle10.Puzzle10{}
	case "11":
		puzzle = puzzle11.Puzzle11{}
	case "12":
		puzzle = puzzle12.Puzzle12{}
	case "13":
		puzzle = puzzle13.Puzzle13{}
	case "14":
		puzzle = puzzle14.Puzzle14{}
	case "15":
		puzzle = puzzle15.Puzzle15{}
	case "16":
		puzzle = puzzle16.Puzzle16{}
	case "17":
		puzzle = puzzle17.Puzzle17{}
	case "18":
		puzzle = puzzle18.Puzzle18{}
	case "19":
		puzzle = puzzle19.Puzzle19{}
	case "20":
		puzzle = puzzle20.Puzzle20{}
	case "21":
		puzzle = puzzle21.Puzzle21{}
	case "22":
		puzzle = puzzle22.Puzzle22{}
	case "23":
		puzzle = puzzle23.Puzzle23{}
	case "24":
		puzzle = puzzle24.Puzzle24{}
	case "25":
		puzzle = puzzle25.Puzzle25{}
	default:
		return nil, fmt.Errorf("Unknown puzzle %v", puzzleid)
	}
	switch partid {
	case "1":
		return puzzle.Part1(input)
	case "2":
		return puzzle.Part2(input)
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