package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string

func (s *Stack) Pop() string {
	if !s.IsEmpty() {
		lastIndex := len(*s) - 1
		element := (*s)[lastIndex]
		*s = (*s)[:lastIndex]
		return element
	}
	return ""
}

func (s *Stack) Push(item string) {
	*s = append(*s, item)
}

func (s *Stack) Prepend(item string) {
	*s = append([]string{item}, *s...)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) ToString() string {
	if s.IsEmpty() {
		return "(empty)\n"
	}
	return fmt.Sprintf("%s\n", strings.Join(*s, " | "))
}

type Instruction struct {
	Amount int
	From   int
	To     int
}

type Data struct {
	Stacks       []Stack
	Instructions []Instruction
}

func CountStacks(line string) int {
	return ((len(strings.Split(line, "")) - 1) / 4) + 1
}

func GetStackIdx(idx int) int {
	return (idx - 1) / 4
}

func (d Data) PrependToStack(line string) {
	chars := strings.Split(line, "")
	for i, v := range chars {
		if v != "" && v != " " && v != "[" && v != "]" {
			d.Stacks[GetStackIdx(i)].Prepend(v)
		}
	}
}

func ParseInstruction(line string) Instruction {
	// move 1 from 2 to 1
	pattern := regexp.MustCompile("[0-9]+")
	matches := pattern.FindAllString(line, -1)
	amount, _ := strconv.Atoi(matches[0])
	from, _ := strconv.Atoi(matches[1])
	to, _ := strconv.Atoi(matches[2])
	return Instruction{
		Amount: amount,
		From:   from,
		To:     to,
	}
}

func readFile() (Data, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return Data{}, err
	}
	defer file.Close()
	data := Data{}
	scanner := bufio.NewScanner(file)
	instructions := false
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			firstLine = false
			// Init stacks
			data.Stacks = make([]Stack, CountStacks(line))
		}
		if line == "" {
			instructions = true
			continue
		}
		if !instructions {
			data.PrependToStack(line)
		} else {
			is := ParseInstruction(line)
			data.Instructions = append(data.Instructions, is)
		}
	}
	return data, nil
}

func (data Data) PrintStacks() {
	for i, s := range data.Stacks {
		fmt.Printf("[%d]: %s", i, s.ToString())
	}
}

func (data Data) Transform() {
	for _, is := range data.Instructions {
		for j := 0; j < is.Amount; j++ {
			moved := data.Stacks[is.From-1].Pop()
			if moved != "" {
				data.Stacks[is.To-1].Push(moved)
			}
		}
	}
}

func main() {
	data, _ := readFile()
	data.PrintStacks()
	data.Transform()
	last := ""
	for _, s := range data.Stacks {
		last += s.Pop()
	}
	fmt.Printf("------\n%s\n", last)
}
