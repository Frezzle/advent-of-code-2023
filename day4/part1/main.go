package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	totalPoints := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		parts = strings.Split(strings.TrimSpace(parts[1]), "|")
		winningNums := getNumbersFromSpacedString(parts[0])
		myNums := getNumbersFromSpacedString(parts[1])

		myNumsMap := make(map[int]bool, len(myNums))
		for _, num := range myNums {
			myNumsMap[num] = true
		}

		cardPoints := 0
		for _, winningNum := range winningNums {
			if myNumsMap[winningNum] {
				if cardPoints == 0 {
					cardPoints = 1
				} else {
					cardPoints *= 2
				}
			}
		}
		totalPoints += cardPoints
	}

	log.Println("Total points:", totalPoints)
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
