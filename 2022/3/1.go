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

type Rucksack struct {
	first  []string
	second []string
}

func (r Rucksack) Intersect() []string {
	// Create map of values from first
	var firstMap = map[string]bool{}
	for _, v := range r.first {
		firstMap[v] = true
	}

	var intersectionMap = map[string]bool{}
	for _, v := range r.second {
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

func (r Rucksack) ToString() string {
	return fmt.Sprintf("%s - %s: %s\n", r.first, r.second, r.Intersect())
}

func (r Rucksack) Priority() int {
	intersection := r.Intersect()
	total := 0
	for _, v := range intersection {
		total += priority[v]
	}
	return total
}

func TotalPriority(rs []Rucksack) int {
	total := 0
	for _, r := range rs {
		total += r.Priority()
	}
	return total
}

func readFile() ([]Rucksack, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var rs []Rucksack
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")
		len := len(chars)
		first := chars[0 : len/2]
		second := chars[len/2 : len]
		rs = append(rs, Rucksack{first, second})
	}
	return rs, scanner.Err()
}

func main() {
	rucksacks, _ := readFile()
	fmt.Printf("Total: %d\n", TotalPriority(rucksacks))
}
