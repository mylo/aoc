package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const COLOR_GREEN = "\033[32m"
const COLOR_RED = "\033[31m"

type Grid [][]int

func (g Grid) Print() {
	for y, yArr := range g {
		for x, height := range yArr {
			PrintValue(height, g.CheckVisibility(y, x, false))
		}
		fmt.Printf("\n")
	}
}

func (g Grid) CountVisible() int {
	visible := 0
	for y, yArr := range g {
		for x, _ := range yArr {
			if g.CheckVisibility(y, x, false) {
				visible++
			}
		}
	}
	return visible
}

func PrintValue(height int, visible bool) {
	color := COLOR_GREEN
	if visible {
		color = COLOR_RED
	}
	fmt.Printf(" %s%d ", color, height)
}

func IsVisible(arrayPos int, height int, line []int, debug bool) bool {
	leftAnyTreeHigher := false
	rightAnyTreeHigher := false
	left := line[0:arrayPos]
	right := line[arrayPos+1:]
	if debug {
		fmt.Printf("\n%v | %d | %v = %v at [%d]=%d)\n", left, height, right, line, arrayPos, line[arrayPos])
	}
	for _, v := range left {
		if height <= v {
			leftAnyTreeHigher = true
		}
	}
	for _, v := range right {
		if height <= v {
			rightAnyTreeHigher = true
		}
	}
	if debug {
		fmt.Printf("leftAnyTreeHigher: %v, rightAnyTreeHigher: %v", leftAnyTreeHigher, rightAnyTreeHigher)
	}
	return !(leftAnyTreeHigher && rightAnyTreeHigher)
}

func (g Grid) GetLines(x int, y int) ([]int, []int) {
	lineX := []int{}
	for _, v := range g {
		lineX = append(lineX, v[x])
	}
	return lineX, g[y]
}

func (g Grid) CheckVisibility(x int, y int, debug bool) bool {
	height := g[y][x]
	if x == 0 || x == len(g)-1 {
		return true
	}
	if y == 0 || y == len(g[0])-1 {
		return true
	}
	lineY, lineX := g.GetLines(x, y)
	if debug {
		fmt.Printf("\n[%d, %d]: %d\n", x, y, height)
		fmt.Printf("Line Y: %v at %d = %d\n", lineY, y, lineY[y])
		fmt.Printf("Line X: %v at %d = %d\n", lineX, x, lineX[x])
		fmt.Printf("\n%d: -> %v: %v\n", x, lineX, IsVisible(x, height, lineX, true))
		fmt.Printf("\n%d: -> %v: %v\n", y, lineY, IsVisible(y, height, lineY, true))
	}
	visibleTopToBottom := IsVisible(y, height, lineY, false)
	visibleLeftToRight := IsVisible(x, height, lineX, false)
	return visibleTopToBottom || visibleLeftToRight
}

func lineToIntArr(line string) []int {
	charArr := strings.Split(line, "")
	intArr := []int{}
	for _, v := range charArr {
		intVal, _ := strconv.Atoi(v)
		intArr = append(intArr, intVal)
	}
	return intArr
}

func readFile() (Grid, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return Grid{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, lineToIntArr(line))
	}
	return grid, nil
}

func main() {
	grid, _ := readFile()
	// grid.Print()
	fmt.Printf("Total visible: %d", grid.CountVisible())
}
