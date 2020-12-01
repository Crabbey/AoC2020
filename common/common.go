package common

import(
	"os"
	"fmt"
	"bufio"
)

type AoCInput struct {
	Path string
	InputFile string
}

type AoCSolution struct {
	Puzzle string
	Part string
	Text string
}

type AoCPuzzle interface {
	Part1(input AoCInput, output *AoCSolution) (*AoCSolution, error)
	Part2(input AoCInput, output *AoCSolution) (*AoCSolution, error)
}

func NewSolution(puzzleid string, partid string) (*AoCSolution) {
	return &AoCSolution{
		Puzzle: puzzleid,
		Part: partid,
		Text: "No solution found",
	}
}

func (a *AoCSolution) Print() {
	fmt.Printf("Puzzle %v Part %v Solution: %v\n", a.Puzzle, a.Part, a.Text)
}

func (a *AoCInput) Default(def string) {
	if a.InputFile == "" {
		a.InputFile = def
	}
}

func (a *AoCInput) Read() ([]string, error) {
	fmt.Println("Reading from " + a.Path + "/" +a.InputFile)
	file, err := os.Open(a.Path + "/" +a.InputFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}