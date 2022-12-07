package main

import (
	_ "embed"
	"sort"
	"strconv"
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

type directory struct {
	parent      *directory
	files       map[string]uint64
	directories map[string]*directory
}

func (d directory) size() uint64 {
	totalSize := uint64(0)
	for _, v := range d.files {
		totalSize = totalSize + v
	}
	for _, v := range d.directories {
		totalSize = totalSize + v.size()
	}
	return totalSize
}

func (d directory) sumLowerThan(size uint64) uint64 {
	sizeTotal := uint64(0)
	if d.size() <= size {
		sizeTotal += d.size()
	}
	for _, v := range d.directories {
		sizeTotal += v.sumLowerThan(size)
	}
	return sizeTotal
}

func (d directory) dirsBiggerThan(size uint64) []directory {
	dirs := make([]directory, 0)

	if d.size() >= size {
		dirs = append(dirs, d)
	}

	for _, dir := range d.directories {
		dirs = append(dirs, dir.dirsBiggerThan(size)...)
	}
	return dirs
}

func readRootDir() directory {
	rootDir := directory{directories: make(map[string]*directory), files: make(map[string]uint64), parent: nil}
	currentDir := &rootDir
	for _, command := range inputArr {
		switch {
		case strings.HasPrefix(command, "$ cd"):
			selectedDir := strings.Split(command, " ")[2]
			switch selectedDir {
			case "..":
				currentDir = currentDir.parent
			case "/":
				currentDir = &rootDir
			default:
				currentDir = currentDir.directories[selectedDir]
			}
		case strings.HasPrefix(command, "dir"):
			dirName := strings.Split(command, " ")[1]
			newDir := directory{directories: make(map[string]*directory), files: make(map[string]uint64), parent: currentDir}
			currentDir.directories[dirName] = &newDir
		default:
			// must be a file then
			info := strings.Split(command, " ")
			filename := info[1]
			filesize, _ := strconv.Atoi(info[0])
			currentDir.files[filename] = uint64(filesize)
		}
	}
	return rootDir
}

func part1() uint64 {
	return readRootDir().sumLowerThan(100_000)
}

func part2() uint64 {
	rootDir := readRootDir()
	neededSpace := 30_000_000 - (70_000_000 - rootDir.size())
	dirs := rootDir.dirsBiggerThan(neededSpace)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size() < dirs[j].size()
	})
	return dirs[0].size()
}

func main() {
	println(part1())
	println(part2())
}
