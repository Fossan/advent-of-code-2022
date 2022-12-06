package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

type set map[int32]struct{}

func (s set) add(value string) {
	for _, v := range value {
		s[v] = struct{}{}
	}
}

func (s set) len() int {
	return len(s)
}

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func findMarker(chars int) int {
	marker := 0
	for i, _ := range input {
		endPos := i + chars
		if endPos <= len(input) {
			signal := input[i:endPos]
			s := set{}
			s.add(signal)
			if s.len() == chars {
				marker = endPos
				break
			}
		}
	}
	return marker
}

func part1() int {
	return findMarker(4)
}

func part2() int {
	return findMarker(14)
}

func main() {
	println(part1())
	println(part2())
}
