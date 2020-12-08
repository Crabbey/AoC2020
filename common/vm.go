package common

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println
var _ = strconv.Itoa

type VM struct {
	MemSpace []VMInstruction
	Accumulator int
	Position int
}

type VMInstruction struct {
	VM *VM
	Command string
	Arguments []string
}

func (i *VMInstruction) ParseArg(argpos int) int {
	x, _ := strconv.Atoi(i.Arguments[argpos])
	return x
}

func (i *VMInstruction) Execute() {
	switch i.Command {
	case "nop":
		i.VM.Position++
	case "jmp":
		i.VM.Position += i.ParseArg(0)
	case "acc":
		i.VM.Accumulator += i.ParseArg(0)
		i.VM.Position++
	}
}

func (v *VM) Step() {
	fmt.Printf("Executing position %v\n", v.Position)
	v.MemSpace[v.Position].Execute()
}

func ReadInstruction(input string, vm *VM) VMInstruction {
	ret := VMInstruction{
		VM: vm,
	}
	parts := strings.Split(input, " ")
	ret.Command = parts[0]
	ret.Arguments = parts[1:]
	return ret
}

func NewVM(input []string) *VM {
	ret := VM{}
	for _, line := range input {
		newInstruction := ReadInstruction(line, &ret)
		ret.MemSpace = append(ret.MemSpace, newInstruction)
	}
	return &ret
}