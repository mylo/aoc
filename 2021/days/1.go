package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile() ([]int, error) {
	file, err := os.Open("../inputs/1.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []int
	for scanner.Scan() {
		var depth, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, depth)
	}
	return lines, scanner.Err()
}

func part1(lines []int) {
	var currentDepth = lines[0]
	var increased = 0
	// We have our input.
	for _, x := range lines {
		if currentDepth < x {
			increased += 1
		}
		currentDepth = x
	}
	fmt.Printf("part1: %d\n", increased)
}

func sum(numbs []int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}

func part2(lines []int) {
	var increased = 0
	for i := 0; i < len(lines)-3; i++ {
		curr := sum(lines[i : i+3])
		next := sum(lines[i+1 : i+4])
		if curr < next {
			increased += 1
		}
	}
	fmt.Printf("part2: %d\n", increased)
}

func main() {
	fmt.Println("Starting 1.go")
	lines, err := readFile()
	if err != nil {
		fmt.Printf("Error is %s", err)
		panic(err)
	}
	part1(lines)
	part2(lines)
}
