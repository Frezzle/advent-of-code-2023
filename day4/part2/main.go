package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card struct {
	Number         int // 1, 2, 3, ...
	Instances      int // All cards start with 1 instance, a card with a copy has 2 instances, etc.
	NumbersMatched int // How many winning numbers appear in my numbers.
}

func main() {
	lines := strings.Split(input, "\n")

	// get all cards
	var cards []*Card
	cardsByNumber := make(map[int]*Card, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, ":")
		numParts := strings.Split(strings.TrimSpace(parts[1]), "|")
		winningNums := getNumbersFromSpacedString(numParts[0])
		myNums := getNumbersFromSpacedString(numParts[1])

		myNumsMap := make(map[int]bool, len(myNums))
		for _, num := range myNums {
			myNumsMap[num] = true
		}

		numsMatched := 0
		for _, winningNum := range winningNums {
			if myNumsMap[winningNum] {
				numsMatched++
			}
		}

		card := &Card{
			Number:         i + 1,
			Instances:      1,
			NumbersMatched: numsMatched,
		}
		cards = append(cards, card)
		cardsByNumber[card.Number] = card
	}

	// Calculate how many instances each card has.
	// Assumes that cards are ordered ascending by card number.
	for _, card := range cards {
		for i := 1; i <= card.NumbersMatched; i++ {
			if cardBelow, ok := cardsByNumber[card.Number+i]; ok {
				// Instead of actually copying cards, simulate many copies copying other cards,
				// and let it compound just by increasing these numbers.
				cardBelow.Instances += card.Instances
			}
		}
	}

	// count total number of cards
	totalCards := 0
	for _, card := range cards {
		totalCards += card.Instances
	}
	log.Println("Total cards:", totalCards)
}

func getNumbersFromSpacedString(s string) []int {
	s = strings.TrimSpace(s)
	parts := strings.Split(s, " ")
	var nums []int
	for _, part := range parts {
		if part == "" {
			continue
		}
		part = strings.TrimSpace(part)
		num, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("failed to convert %q to int\n", part)
		}
		nums = append(nums, num)
	}

	return nums
}
