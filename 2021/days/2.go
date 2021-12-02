package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Direction string
	Unit      int
}

func readFile() ([]Command, error) {
	file, err := os.Open("../inputs/2.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []Command
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		var unit, err = strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		lines = append(lines, Command{
			Direction: line[0],
			Unit:      unit,
		})
	}
	return lines, scanner.Err()
}

func part1(lines []Command) {
	var solution = 0
	var horizontalPos = 0
	var depth = 0
	for _, x := range lines {
		if x.Direction == "up" {
			depth -= x.Unit
		} else if x.Direction == "down" {
			depth += x.Unit
		} else if x.Direction == "forward" {
			horizontalPos += x.Unit
		}
	}
	solution = horizontalPos * depth
	fmt.Printf("Horizontal: (%d) * Depth: (%d) = %d\n", horizontalPos, depth, solution)
	fmt.Printf("part1: %d\n", solution)
}

func part2(lines []Command) {
	var solution = 0
	var horizontalPos = 0
	var depth = 0
	var aim = 0
	for _, x := range lines {
		if x.Direction == "up" {
			aim -= x.Unit
		} else if x.Direction == "down" {
			aim += x.Unit
		} else if x.Direction == "forward" {
			horizontalPos += x.Unit
			depth += (aim * x.Unit)
		}
	}
	solution = horizontalPos * depth
	fmt.Printf("Horizontal: (%d) * Depth: (%d) = %d\n", horizontalPos, depth, solution)
	fmt.Printf("part2: %d\n", solution)
}

func main() {
	fmt.Println("Starting 2.go")
	lines, err := readFile()
	if err != nil {
		fmt.Printf("Error is %s", err)
		panic(err)
	}
	part1(lines)
	part2(lines)
}
