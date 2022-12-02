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

type Outcome int64

const (
	Lose Outcome = iota
	Draw
	Win
)

type OutcomeMove struct {
	Move    Move
	Outcome Outcome
}

var GameOutcomes = map[Move]map[Outcome]Move{
	Rock: {
		Win:  Paper,
		Draw: Rock,
		Lose: Scissor,
	},
	Paper: {
		Win:  Scissor,
		Draw: Paper,
		Lose: Rock,
	},
	Scissor: {
		Win:  Rock,
		Draw: Scissor,
		Lose: Paper,
	},
}

type Game struct {
	Move           Move
	Counter        Move
	DesiredOutcome Outcome
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

func (m Move) ToString() string {
	switch m {
	case Rock:
		return "R"
	case Paper:
		return "P"
	case Scissor:
		return "S"
	default:
		return ""
	}
}

func (o Outcome) ToString() string {
	switch o {
	case Win:
		return "W"
	case Lose:
		return "L"
	case Draw:
		return "D"
	default:
		return ""
	}
}

func (g Game) ToString() string {
	return fmt.Sprintf("%s %s = %s (%d)\n", g.Move.ToString(), g.Counter.ToString(), g.DesiredOutcome.ToString(), g.Outcome())
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
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissor
	default:
		return -1
	}
}

func ToDesiredOutcome(m string) Outcome {
	switch m {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return -1
	}
}

func ToCounter(m Move, o Outcome) Move {
	return GameOutcomes[m][o]
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
		opponentMove := MoveToMove(moves[0])
		outcome := ToDesiredOutcome(moves[1])
		games = append(games, Game{
			Move:           opponentMove,
			DesiredOutcome: outcome,
			Counter:        ToCounter(opponentMove, outcome),
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
