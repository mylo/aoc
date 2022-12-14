package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Operation struct {
	Cycles int
	Value  int
}

func readFile() ([]int, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return []int{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	ops := []int{}
	for scanner.Scan() {
		lineArr := strings.Split(scanner.Text(), " ")
		op := lineArr[0]
		// Always add 0, noop or addx
		ops = append(ops, 0)
		if op == "addx" {
			value, _ := strconv.Atoi(lineArr[1])
			ops = append(ops, value)
		}
	}
	return ops, nil
}

func SignalStrength(ops []int) int {
	sum := 0
	var registerX = 1
	for i := 1; i < len(ops)+1; i++ {
		op := ops[i-1]
		if ((i - 20) % 40) == 0 {
			sum += (registerX * i)
			fmt.Printf("Multiplying register x %d with %d =>: %d\n", registerX, i, sum)
		}
		registerX += op
	}
	return sum
}

func main() {
	ops, _ := readFile()
	fmt.Printf("%v", SignalStrength(ops))
}
