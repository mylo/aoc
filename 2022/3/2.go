package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var priority = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

type Group struct {
	rucksacks []Rucksack
}

type Rucksack struct {
	items []string
}

func Intersect(arr1 []string, arr2 []string) []string {
	// Create map of values from first
	var firstMap = map[string]bool{}
	for _, v := range arr1 {
		firstMap[v] = true
	}

	var intersectionMap = map[string]bool{}
	for _, v := range arr2 {
		if firstMap[v] {
			intersectionMap[v] = true
		}
	}
	keys := make([]string, 0, len(intersectionMap))
	for k := range intersectionMap {
		keys = append(keys, k)
	}
	return keys
}

func (g Group) Intersect() []string {
	res := g.rucksacks[0].items
	for i := 1; i < len(g.rucksacks); i++ {
		res = Intersect(res, g.rucksacks[i].items)
	}
	return res
}

func (r Rucksack) ToString() string {
	return fmt.Sprintf("%s\n", r.items)
}

func Print(groups []Group) {
	for i, g := range groups {
		fmt.Printf("Group: %d\n", i)
		for _, r := range g.rucksacks {
			fmt.Printf("%s", r.ToString())
		}
		fmt.Printf("\t %s", g.Intersect())
	}
}

func (g Group) Priority() int {
	intersection := g.Intersect()
	total := 0
	for _, v := range intersection {
		total += priority[v]
	}
	return total
}

func TotalPriority(gs []Group) int {
	total := 0
	for _, g := range gs {
		total += g.Priority()
	}
	return total
}

func LineToRucksack(line string) Rucksack {
	return Rucksack{
		items: strings.Split(line, ""),
	}
}

func readFile() ([]Group, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var groups []Group
	var rucksacks []Rucksack
	for scanner.Scan() {
		rucksacks = append(rucksacks, LineToRucksack(scanner.Text()))
		rsLength := len(rucksacks)
		if rsLength != 0 && rsLength%3 == 0 {
			// reset and add to group
			groups = append(groups, Group{rucksacks})
			rucksacks = []Rucksack{}
		}
	}
	return groups, scanner.Err()
}

func main() {
	gs, _ := readFile()
	fmt.Printf("Total: %d\n", TotalPriority(gs))
}
