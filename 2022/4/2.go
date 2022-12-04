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
	Sections []int
	From     int
	To       int
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
	sections := []int{}
	for i := from; i <= to; i++ {
		sections = append(sections, i)
	}
	return Assignment{From: from, To: to, Sections: sections}
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

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Overlap(A Assignment, B Assignment) bool {
	for _, v := range A.Sections {
		if contains(B.Sections, v) {
			return true
		}
	}
	return false
}

func (shift Shift) Overlap() bool {
	assignmentA := shift.Assignments.Assignments[0]
	assignmentB := shift.Assignments.Assignments[1]
	return Overlap(assignmentA, assignmentB)
}

func (sL ShiftList) Overlap() int {
	total := 0
	for _, s := range sL.Shifts {
		if s.Overlap() {
			total += 1
		}
	}
	return total
}

func main() {
	shiftList, _ := readFile()
	fmt.Printf("Overlap: %d", shiftList.Overlap())
}
