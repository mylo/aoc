package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ShiftList struct {
	Shifts []Shift
}

type Shift struct {
	Assignments AssignmentList
}

type AssignmentList struct {
	Assignments []Assignment
}

type Assignment struct {
	From int
	To   int
}

func (aL AssignmentList) ToString() string {
	ret := ""
	for _, s := range aL.Assignments {
		ret += fmt.Sprintf("[%d, %d]", s.From, s.To)
	}
	return ret
}

func (sL ShiftList) Print() {
	for i, s := range sL.Shifts {
		fmt.Printf("Shift: %d > %s\n", i, s.Assignments.ToString())
	}
}

func ShiftToAssignment(shift string) Assignment {
	// 1-3 => [1, 2, 3]
	chars := strings.Split(shift, "-")
	from, _ := strconv.Atoi(chars[0])
	to, _ := strconv.Atoi(chars[1])
	return Assignment{From: from, To: to}
}

func LineToShift(line string) Shift {
	shifts := strings.Split(line, ",")
	assignments := []Assignment{}
	for _, v := range shifts {
		assignments = append(assignments, ShiftToAssignment(v))
	}
	return Shift{
		Assignments: AssignmentList{
			Assignments: assignments,
		},
	}
}

func readFile() (ShiftList, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return ShiftList{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	shifts := []Shift{}
	for scanner.Scan() {
		shifts = append(shifts, LineToShift(scanner.Text()))
	}
	return ShiftList{shifts}, scanner.Err()
}

func FullyContains(A Assignment, B Assignment) bool {
	// A shift A fully includes shift B, iff:
	// A.from >= B.from
	// B.to <= A.to
	return A.From <= B.From && B.To <= A.To
}

func (shift Shift) FullyContains() bool {
	assignmentA := shift.Assignments.Assignments[0]
	assignmentB := shift.Assignments.Assignments[1]
	fmt.Printf("%s\n", shift.Assignments.ToString())
	return FullyContains(assignmentA, assignmentB) || FullyContains(assignmentB, assignmentA)
}

func (sL ShiftList) FullyContains() int {
	total := 0
	for _, s := range sL.Shifts {
		if s.FullyContains() {
			total += 1
		}
	}
	return total
}

func main() {
	shiftList, _ := readFile()
	fmt.Printf("Fully contains: %d", shiftList.FullyContains())
}
