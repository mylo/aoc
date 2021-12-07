package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile() ([][]string, error) {
	file, err := os.Open("../inputs/3.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lines := [][]string{}
	index := 0
	for scanner.Scan() {
		if index == 0 {
			lines = make([][]string, len(scanner.Text()))
		}
		chars := strings.Split(scanner.Text(), "")
		for i, x := range chars {
			lines[i] = append(lines[i], x)
		}
		index++
	}
	return lines, scanner.Err()
}

func readFileSimple() ([]string, error) {
	file, _ := os.Open("../inputs/3.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

type LineValue struct {
	Gamma   string
	Epsilon string
}

func getMostCommonBit(line []string) string {
	var dict = make(map[string]int)
	for _, x := range line {
		dict[x] = dict[x] + 1
	}
	if dict["0"] <= dict["1"] {
		return "1"
	}
	return "0"
}

func getLineValue(line []string) LineValue {
	mostCommon := getMostCommonBit(line)
	val := LineValue{}
	if mostCommon == "1" {
		// 1
		val.Gamma = "1"
		val.Epsilon = "0"
	} else {
		// 0
		val.Gamma = "0"
		val.Epsilon = "1"
	}
	return val
}

func getBinaryValue(value string) int64 {
	val, err := strconv.ParseInt(value, 2, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func part1(lines [][]string) {
	gamma := ""
	epsilon := ""
	// Find the current array length. Used in the outer for loop
	for _, x := range lines {
		lineValue := getLineValue(x)
		gamma += lineValue.Gamma
		epsilon += lineValue.Epsilon
	}
	gammaValue := getBinaryValue(gamma)
	epsilonValue := getBinaryValue(epsilon)
	solution := gammaValue * epsilonValue
	fmt.Printf("part1: %d\n", solution)
}

//https://github.com/lynerist/Advent-of-code-2021-golang/blob/master/Day_03/day03_b.go
func filterValues(values []string, bitToConsider int, mostCommon bool) string {
	if len(values) == 1 {
		return values[0]
	}
	var valueWithZero, valueWithOne []string
	for _, value := range values {
		if rune(value[bitToConsider]) == '0' {
			valueWithZero = append(valueWithZero, value)
		} else {
			valueWithOne = append(valueWithOne, value)
		}
	}
	if len(valueWithOne) >= len(valueWithZero) == mostCommon {
		return filterValues(valueWithOne, bitToConsider+1, mostCommon)
	}
	return filterValues(valueWithZero, bitToConsider+1, mostCommon)
}

func part2(lines []string) {
	oxygen := getBinaryValue(filterValues(lines, 0, true))
	carbondioxide := getBinaryValue(filterValues(lines, 0, false))
	solution := oxygen * carbondioxide
	fmt.Printf("part1: %d\n", solution)
}

func main() {
	fmt.Println("Starting 3.go")
	// lines, _ := readFile()
	linesSimple, _ := readFileSimple()
	// part1(lines)
	part2(linesSimple)
}
