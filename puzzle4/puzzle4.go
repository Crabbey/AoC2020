package puzzle4

import (
	"strings"
	"strconv"
	"fmt"
	"regexp"
	"github.com/crabbey/aoc2020/common"
	"github.com/davecgh/go-spew/spew"
)

var _ = spew.Dump
var _ = fmt.Println

type Puzzle4 struct {

}

var HairColourRegexp = regexp.MustCompile(`^#[a-fA-F0-9]{6}$`)
var EyeColourRegexp = regexp.MustCompile(`^(?:amb|blu|brn|gry|grn|hzl|oth)$`)
var PassportIDRegexp = regexp.MustCompile(`^[0-9]{9}$`)

type Passport struct {
	BirthYear      string // byr(Birth Year)
	IssueYear      string // iyr (Issue Year)
	ExpirationYear string // eyr (Expiration Year)
	Height         string // hgt (Height)
	HairColour     string // hcl (Hair Color)
	EyeColour      string // ecl (Eye Color)
	PassportID     string // pid (Passport ID)
	CountryID      string // cid (Country ID)
}

func (p *Passport) IsValid() bool {
	if p.BirthYear == "" || p.IssueYear == "" || p.ExpirationYear == "" {
		return false
	}
	if p.Height == "" || p.HairColour == "" || p.EyeColour == "" {
		return false
	}
	if p.PassportID == "" {
		return false
	}

	return true
}

func Between(in string, lower, upper int) bool {
	input, _ := strconv.Atoi(in)
	if input >= lower && input <= upper {
		return true
	}
	return false
}

func (p *Passport) IsValidHairColour() bool {
	return HairColourRegexp.MatchString(p.HairColour)
}

func (p *Passport) IsValidEyeColour() bool {
	return EyeColourRegexp.MatchString(p.EyeColour)
}

func (p *Passport) IsValidPassportID() bool {
	return PassportIDRegexp.MatchString(p.PassportID)
}

func (p *Passport) IsValidHeight() bool {
	if p.Height == "" {
		return false
	}
	if p.Height[len(p.Height)-2:] == "cm" {
		return Between(p.Height[0:len(p.Height)-2], 150, 193)
	} else if p.Height[len(p.Height)-2:] == "in" {
		return Between(p.Height[0:len(p.Height)-2], 59, 76)
	}
	return false
}

func (p *Passport) IsStrictValid() bool {
	if !Between(p.BirthYear, 1920, 2002) || !Between(p.IssueYear, 2010, 2020) || !Between(p.ExpirationYear, 2020, 2030) {
		return false
	}
	if !p.IsValidHeight() || !p.IsValidHairColour() ||!p.IsValidEyeColour() {
		return false
	}
	if !p.IsValidPassportID() {
		return false
	}

	return true
}

func (p Puzzle4) ParsePassports(input []string) []Passport {
	var ret []Passport
	passport := Passport{}
	for _, line := range input {
		if line == "" {
			ret = append(ret, passport)
			passport = Passport{}
			continue
		}
		parts := strings.Split(line, " ")
		for _, part := range parts {
			kv := strings.Split(part, ":")
			switch kv[0] {
			case "byr":
				passport.BirthYear = kv[1]
			case "iyr":
				passport.IssueYear = kv[1]
			case "eyr":
				passport.ExpirationYear = kv[1]
			case "hgt":
				passport.Height = kv[1]
			case "hcl":
				passport.HairColour = kv[1]
			case "ecl":
				passport.EyeColour = kv[1]
			case "pid":
				passport.PassportID = kv[1]
			case "cid":
				passport.CountryID = kv[1]
			default:
				fmt.Printf("Unknown KV pair: %v", parts)
			}
		}
	}
	ret = append(ret, passport)
	return ret
}

func (p Puzzle4) Part1(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}
	passports := p.ParsePassports(i)
	count := 0
	for _, passport := range passports {
		if passport.IsValid() {
			count++
		}
	}
	output.Text = fmt.Sprintf("Valid passports: %v", count)
	return output, nil
}

func (p Puzzle4) Part2(input common.AoCInput, output *common.AoCSolution) (*common.AoCSolution, error) {
	i, err := input.Read()
	if err != nil {
		spew.Dump(i)
		return output, err
	}

	passports := p.ParsePassports(i)
	count := 0
	for _, passport := range passports {
		if passport.IsStrictValid() {
			count++
		}
	}
	output.Text = fmt.Sprintf("Valid passports: %v", count)
	
	return output, nil
}