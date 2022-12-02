package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Move int64

const (
	Rock Move = iota
	Paper
	Scissor
)

type Game struct {
	Move    Move
	Counter Move
}

func (g Game) ShapeValue() int {
	switch g.Counter {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissor:
		return 3
	default:
		return 0
	}
}

func (g Game) GameValue() int {
	if g.Move == g.Counter {
		// Draw
		return 3
	}
	if g.Move == Rock {
		if g.Counter == Paper {
			return 6
		}
		if g.Counter == Scissor {
			return 0
		}
	}
	if g.Move == Paper {
		if g.Counter == Scissor {
			return 6
		}
		if g.Counter == Paper {
			return 0
		}
	}
	if g.Move == Scissor {
		if g.Counter == Rock {
			return 6
		}
		if g.Counter == Paper {
			return 0
		}
	}
	return 0
}

func (g Game) Outcome() int {
	return g.GameValue() + g.ShapeValue()
}

func MoveToMove(m string) Move {
	switch m {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	default:
		return -1
	}
}

func CounterToMove(m string) Move {
	switch m {
	case "X":
		return 0
	case "Y":
		return 1
	case "Z":
		return 2
	default:
		return -1
	}
}

func readFile() ([]Game, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		moves := strings.Split(scanner.Text(), " ")
		games = append(games, Game{
			Move:    MoveToMove(moves[0]),
			Counter: CounterToMove(moves[1]),
		})
	}
	return games, scanner.Err()
}

func main() {
	games, _ := readFile()
	total := 0
	for _, g := range games {
		total += g.Outcome()
	}
	fmt.Printf("Total score: %d\n", total)
}
