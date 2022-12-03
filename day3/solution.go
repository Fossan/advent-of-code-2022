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

type set map[int]struct{}

func (s set) add(value int) {
	s[value] = struct{}{}
}

func (s set) has(value int) bool {
	_, ok := s[value]
	return ok
}

func shift(c int32) int {
	if c >= 97 && c <= 122 {
		return int(c - 96)
	} else {
		return int(c - 38)
	}
}

func (s set) intersect(s2 set) set {
	sIntersection := set{}
	for k := range s {
		_, ok := s2[k]
		if ok {
			sIntersection[k] = struct{}{}
		}
	}
	return sIntersection
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func part1() int {
	sum := 0
	for _, v := range inputArr {
		rucksackPivot := len(v) / 2
		comp1 := v[:rucksackPivot]
		comp2 := v[rucksackPivot:]
		s1 := set{}
		s2 := set{}
		for _, vv := range comp1 {
			s1.add(shift(vv))
		}
		for _, vv := range comp2 {
			if s1.has(shift(vv)) {
				s2[shift(vv)] = struct{}{}
			}
		}
		for s := range s2 {
			sum = sum + s
		}
	}
	return sum
}

func part2() int {
	sum := 0
	for _, v := range chunkBy[string](inputArr, 3) {
		s1 := set{}
		s2 := set{}
		s3 := set{}
		for _, vv := range v[0] {
			s1.add(shift(vv))
		}
		for _, vv := range v[1] {
			s2.add(shift(vv))
		}
		for _, vv := range v[2] {
			s3.add(shift(vv))
		}
		for s := range s1.intersect(s2.intersect(s3)) {
			sum = sum + s
		}
	}
	return sum
}

func main() {
	println(part1())
	println(part2())
}
