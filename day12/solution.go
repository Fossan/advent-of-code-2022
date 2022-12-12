package main

import (
	_ "embed"
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

func loadMaps() (map[image.Point]rune, map[image.Point]int, []image.Point, image.Point) {
	var start, end image.Point
	height := map[image.Point]rune{}
	for x, row := range strings.Fields(input) {
		for y, char := range row {
			height[image.Point{X: x, Y: y}] = char
			if char == 'S' {
				start = image.Point{X: x, Y: y}
			} else if char == 'E' {
				end = image.Point{X: x, Y: y}
			}
		}
	}
	height[start], height[end] = 'a', 'z'

	distance := map[image.Point]int{end: 0}
	queue := []image.Point{end}

	return height, distance, queue, start
}

func traverse(current image.Point, height map[image.Point]rune, distance map[image.Point]int, queue *[]image.Point) {
	for _, direction := range []image.Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		next := current.Add(direction)
		_, valid := height[next]
		_, visited := distance[next]

		if !visited && valid && height[current] <= height[next]+1 {
			distance[next] = distance[current] + 1
			*queue = append(*queue, next)
		}
	}
}

func part1() int {
	height, distance, queue, start := loadMaps()

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		traverse(current, height, distance, &queue)
	}

	return distance[start]
}

func part2() int {
	height, distance, queue, _ := loadMaps()
	var shortestPath *image.Point
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if height[current] == 'a' && shortestPath == nil {
			shortestPath = &current
		}

		traverse(current, height, distance, &queue)
	}
	return distance[*shortestPath]
}

func main() {
	println(part1())
	println(part2())
}
