package main

import (
	_ "embed"
	"fmt"
	"image"
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

type set map[image.Point]struct{}

func (s set) add(point image.Point) {
	s[point] = struct{}{}
}

func (s set) len() int {
	return len(s)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sgn(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func part1() int {
	directions := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	visited := set{}
	rope := make([]image.Point, 10)

	for _, s := range inputArr {
		var direction rune
		var step int
		fmt.Sscanf(s, "%c %d", &direction, &step)

		for i := 0; i < step; i++ {
			rope[0] = rope[0].Add(directions[direction])
			for i := 1; i < len(rope); i++ {
				d := rope[i-1].Sub(rope[i])
				if abs(d.X) > 1 || abs(d.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{X: sgn(d.X), Y: sgn(d.Y)})
				}
			}

			visited.add(rope[1])
		}
	}
	return visited.len()
}

func part2() int {
	directions := map[rune]image.Point{'U': {0, -1}, 'R': {1, 0}, 'D': {0, 1}, 'L': {-1, 0}}
	visited := set{}
	rope := make([]image.Point, 10)

	for _, s := range inputArr {
		var direction rune
		var step int
		fmt.Sscanf(s, "%c %d", &direction, &step)

		for i := 0; i < step; i++ {
			rope[0] = rope[0].Add(directions[direction])

			for i := 1; i < len(rope); i++ {
				d := rope[i-1].Sub(rope[i])
				if abs(d.X) > 1 || abs(d.Y) > 1 {
					rope[i] = rope[i].Add(image.Point{X: sgn(d.X), Y: sgn(d.Y)})
				}
			}

			visited.add(rope[len(rope)-1])
		}
	}
	return visited.len()
}

func main() {
	println(part1())
	println(part2())
}
