package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elf struct {
	Idx           int
	TotalCalories int
	Calories      []int
}

func readFile() ([]Elf, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var elves []Elf
	idx := 0
	currentElf := Elf{TotalCalories: 0, Idx: idx, Calories: []int{}}
	for scanner.Scan() {
		nextLine := scanner.Text()
		if nextLine != "" {
			var cals, err = strconv.Atoi(scanner.Text())
			if err != nil {
				return nil, err
			}
			currentElf.TotalCalories = currentElf.TotalCalories + cals
			currentElf.Calories = append(currentElf.Calories, cals)
		} else {
			idx = idx + 1
			elves = append(elves, currentElf)
			currentElf = Elf{TotalCalories: 0, Idx: idx + 1, Calories: []int{}}
		}
	}
	return elves, scanner.Err()
}

func main() {
	elves, _ := readFile()
	sort.Slice(elves, func(i, j int) bool {
		return elves[i].TotalCalories > elves[j].TotalCalories
	})
	fmt.Printf("1st: Elf no %d, calories: %d\n", elves[0].Idx, elves[0].TotalCalories)
	fmt.Printf("2nd: Elf no %d, calories: %d\n", elves[1].Idx, elves[1].TotalCalories)
	fmt.Printf("3rd: Elf no %d, calories: %d\n", elves[2].Idx, elves[2].TotalCalories)
	fmt.Printf("----------------------------\n")
	fmt.Printf("Total: %d\n", elves[0].TotalCalories+elves[1].TotalCalories+elves[2].TotalCalories)
}
