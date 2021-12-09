package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bingo struct {
	Nums  []int
	Cards [][][]int
}

func readFile() (Bingo, error) {
	file, _ := os.Open("../inputs/4.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	game := Bingo{}
	nums := []int{}
	for _, char := range strings.Split(scanner.Text(), ",") {
		num, _ := strconv.Atoi(char)
		nums = append(nums, num)
	}
	game.Nums = nums

	// Read the rest
	scanner.Scan()
	index := 0
	currentCard := make([][]int, 5)
	var cards [][][]int
	for scanner.Scan() {
		if scanner.Text() != "" {
			// currentRow := []int{}
			for _, char := range strings.Fields(scanner.Text()) {
				num, _ := strconv.Atoi(char)
				currentCard[index] = append(currentCard[index], num)
			}
			index++
		} else {
			index = 0
			cards = append(cards, currentCard)
			currentCard = make([][]int, 5)
		}
	}

	game.Cards = cards

	return game, nil
}

func printFormatted(card [][]int) {
	for _, row := range card {
		fmt.Printf("%d\t%d\t%d\t%d\t%d\n", row[0], row[1], row[2], row[3], row[4])
	}
	fmt.Print("\n")
}

func didWinRow(card [][]int) bool {
	for _, row := range card {
		var didWin = true
		for _, char := range row {
			if char != -1 {
				didWin = false
			}
		}
		if didWin {
			return true
		}
	}
	return false
}

func didWinDiagonal(card [][]int) bool {
	// First diagonal
	printFormatted(card)
	didWinBotUp := card[0][4] == -1 && card[1][0] == -1 && card[2][0] == -1 && card[3][0] == -1 && card[4][0] == -1
	didWinTopDown := card[0][0] == -1 && card[1][1] == -1 && card[2][2] == -1 && card[3][3] == -1 && card[4][4] == -1
	return didWinBotUp || didWinTopDown
}

func didWinCol(card [][]int) bool {
	for i := 0; i < 5; i++ {
		if card[0][i] == -1 &&
			card[1][i] == -1 &&
			card[2][i] == -1 &&
			card[3][i] == -1 &&
			card[4][i] == -1 {
			return true
		}
	}
	return false
}

func findAndMark(card [][]int, input []int) [][]int {
	for _, mark := range input {
		for i, row := range card {
			for j, char := range row {
				if char == mark {
					card[i][j] = -1
				}
			}
		}
	}
	return card
}

func didWin(card [][]int, input []int) bool {
	m := findAndMark(card, input)
	if didWinRow(m) {
		return true
	}
	if didWinCol(m) {
		return true
	}
	return false
}

func sumUnmarked(card [][]int, input []int) int {
	m := findAndMark(card, input)
	printFormatted(m)
	count := 0
	for _, row := range m {
		for _, char := range row {
			if char != -1 {
				count += char
			}
		}
	}
	return count
}

type Win struct {
	inputIndex int
	cardIndex  int
}

func getBestWin() Win {
	game, _ := readFile()
	bestCard := -1
	winAt := len(game.Nums) + 1
	for cIndex, card := range game.Cards {
		// Check when this card wins.
		for i := 0; i < len(game.Nums); i++ {
			didWin := didWin(card, game.Nums[0:i+1])
			if i <= winAt && didWin {
				winAt = i
				bestCard = cIndex
			}
		}
	}
	return Win{
		inputIndex: winAt,
		cardIndex:  bestCard,
	}
}

func getWorstWin() Win {
	game, _ := readFile()
	worstCard := -1
	personalWorst := -1
	for cIndex, card := range game.Cards {
		// Check when this card wins.
		hasWon := false
		for i := 0; i < len(game.Nums); i++ {
			if didWin(card, game.Nums[0:i+1]) && !hasWon {
				hasWon = true
				if personalWorst <= i {
					personalWorst = i
					worstCard = cIndex
					break
				}
			}
		}
	}
	return Win{
		inputIndex: personalWorst,
		cardIndex:  worstCard,
	}
}

func part1() {
	game, _ := readFile()
	bestWin := getBestWin()
	card := game.Cards[bestWin.cardIndex]
	inputThatWon := game.Nums[bestWin.inputIndex]

	sum := sumUnmarked(card, game.Nums[0:bestWin.inputIndex+1])
	fmt.Printf("Best card: %d - wins at %d with number %d\n", bestWin.cardIndex, bestWin.inputIndex, inputThatWon)
	fmt.Printf("Sum of unmarked: %d * won at: %d = %d\n", sum, inputThatWon, sum*inputThatWon)
}

func part2() {
	game, _ := readFile()
	worstWin := getWorstWin()
	card := game.Cards[worstWin.cardIndex]
	inputThatWon := game.Nums[worstWin.inputIndex]

	sum := sumUnmarked(card, game.Nums[0:worstWin.inputIndex+1])
	fmt.Printf("Worst card: %d - wins at %d with number %d\n", worstWin.cardIndex, worstWin.inputIndex, inputThatWon)
	fmt.Printf("Sum of unmarked: %d * won at: %d = %d\n", sum, inputThatWon, sum*inputThatWon)
}

func main() {
	fmt.Println("Starting 4.go")
	// part1()
	part2()
}
