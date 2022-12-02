package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string
var inputArr []string

func init() {
	input = strings.TrimRight(input, "\n")
	inputArr = strings.Split(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

const (
	Rock string = "X"
	Win
	Paper = "Y"
	Draw
	Scissors = "Z"
	Lose
)

var playPoints = map[string]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var equivalents = map[string]string{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var winRules = map[string]string{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var loseRules = map[string]string{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

func part1() int {
	totalPoints := 0
	for _, value := range inputArr {
		split := strings.Split(value, " ")
		op := equivalents[split[0]]
		me := split[1]
		if op == me {
			totalPoints = totalPoints + 3
		} else if winRules[op] != me {
			totalPoints = totalPoints + 6
		}
		totalPoints = totalPoints + playPoints[me]
	}
	return totalPoints
}

func part2() int {
	totalPoints := 0
	for _, value := range inputArr {
		split := strings.Split(value, " ")
		op := equivalents[split[0]]
		me := split[1]
		if me == Win {
			totalPoints = totalPoints + playPoints[winRules[op]]
		} else if me == Draw {
			totalPoints = totalPoints + 3
			totalPoints = totalPoints + playPoints[op]
		} else if me == Lose {
			totalPoints = totalPoints + 6
			totalPoints = totalPoints + playPoints[loseRules[op]]
		}
	}
	return totalPoints
}

func main() {
	println(part1())
	println(part2())
}
