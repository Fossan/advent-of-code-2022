package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string
var inputArr []string

const cycleModulo = 40

func init() {
	input = strings.TrimRight(input, "\n")
	inputArr = strings.Split(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func checkCycle(cycle int, total *int, x int) {
	if cycle%cycleModulo == 20 {
		*total += cycle * x
	}
}

func drawCRT(crt [][]string, crtPos int, crtIdx int, x int) {
	if crtPos == x-1 || crtPos == x || crtPos == x+1 {
		crt[crtIdx][crtPos] = "#"
	}
}

func performCycleWithCRT(cycle *int, crtPos *int, crtIdx *int) {
	*cycle += 1
	*crtPos += 1
	if *cycle%cycleModulo == 1 {
		*crtIdx += 1
		*crtPos = 0
	}
}

func part1() int {
	total := 0
	cycle := 1
	x := 1
	for _, v := range inputArr {
		xx := 0
		if v == "noop" {
			cycle += 1
			continue
		}
		fmt.Sscanf(v, "addx %d", &xx)
		cycle += 1
		checkCycle(cycle, &total, x)
		cycle += 1
		x += xx
		checkCycle(cycle, &total, x)
	}
	return total
}

func part2() [][]string {
	cycle := 1
	x := 1
	crt := make([][]string, 6)
	for i, _ := range crt {
		crt[i] = make([]string, cycleModulo)
		for j, _ := range crt[i] {
			crt[i][j] = "."
		}
	}
	crtIdx := 0
	crtPos := 0
	for _, v := range inputArr {
		xx := 0
		if v == "noop" {
			drawCRT(crt, crtPos, crtIdx, x)
			performCycleWithCRT(&cycle, &crtPos, &crtIdx)
			continue
		}
		fmt.Sscanf(v, "addx %d", &xx)

		drawCRT(crt, crtPos, crtIdx, x)
		performCycleWithCRT(&cycle, &crtPos, &crtIdx)

		drawCRT(crt, crtPos, crtIdx, x)
		performCycleWithCRT(&cycle, &crtPos, &crtIdx)
		x += xx
	}
	return crt
}

func main() {
	println(part1())
	part2 := part2()
	for _, p := range part2 {
		fmt.Println(p)
	}
}
