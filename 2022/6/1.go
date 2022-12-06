package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const SIGNAL_LENGTH = 14

func readFile() ([]string, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return []string{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	line := ""
	for scanner.Scan() {
		line += scanner.Text()
	}
	return strings.Split(line, ""), nil
}

func IsUnique(chars []string) bool {
	for i, v := range chars {
		for j := i + 1; j < len(chars); j++ {
			if chars[j] == v {
				return false
			}
		}
	}
	return true
}

func FindMarkerPos(chars []string) ([]string, int) {
	for i := 0; i <= len(chars)-SIGNAL_LENGTH; i++ {
		copy := chars
		currSlice := copy[i : i+SIGNAL_LENGTH]
		if IsUnique(currSlice) {
			return currSlice, i + SIGNAL_LENGTH
		}
	}
	return []string{}, -1
}

func main() {
	data, _ := readFile()
	markerSlice, markerPos := FindMarkerPos(data)
	fmt.Printf("Found slice [%s] at %d", strings.Join(markerSlice, ","), markerPos)
}
