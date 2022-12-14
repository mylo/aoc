package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Grid [][]int

func (g Grid) HighestScenicScore() int {
	bestYet := 0
	for y, yArr := range g {
		for x, _ := range yArr {
			currTreeScenicScore := g.ScenicScore(y, x)
			if bestYet <= currTreeScenicScore {
				bestYet = currTreeScenicScore
			}
		}
	}
	return bestYet
}

type Scene struct {
	top    []int
	bottom []int
	left   []int
	right  []int
}

func (g Grid) GetLines(x int, y int) ([]int, []int) {
	lineX := []int{}
	for _, v := range g {
		lineX = append(lineX, v[x])
	}
	return lineX, g[y]
}

func GetPartialScene(pos int, line []int) ([]int, []int) {
	first := line[0:pos]
	second := line[pos+1:]
	return first, second
}

func (g Grid) GetFullScene(x int, y int) Scene {
	lineX, lineY := g.GetLines(x, y)
	top, bottom := GetPartialScene(x, lineX)
	left, right := GetPartialScene(y, lineY)
	return Scene{
		top,
		left,
		right,
		bottom,
	}
}

func GetScore(arr []int, height int) int {
	score := 0
	for _, v := range arr {
		if v < height {
			score++
		}
	}
	return score
}

func (s Scene) GetScenicScore(height int) int {
	top := GetScore(s.top, height)
	bottom := GetScore(s.bottom, height)
	left := GetScore(s.left, height)
	right := GetScore(s.right, height)
	return top * bottom * left * right
}

func (g Grid) ScenicScore(x int, y int) int {
	scene := g.GetFullScene(x, y)
	return scene.GetScenicScore(g[x][y])
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
	fmt.Printf("Total visible: %d", grid.HighestScenicScore()())
}
