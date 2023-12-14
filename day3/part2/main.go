package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// The position of something in a text file e.g. a number or gear.
type Position struct {
	Line   int // e.g. line 1, 2, 3, ...
	Column int // e.g. column 1, 2, 3, ...
}

func main() {
	// First pass:
	// - get numbers and their positions (numbers can be repeated)
	// - get gear positions
	numsByPosition := make(map[Position]*string)
	var gearPositions []Position
	lines := strings.Split(input, "\n")
	for li, line := range lines {
		var currentNum *string
		for ci, char := range line {
			charType := getCharType(char)
			switch charType {
			case "digit":
				if currentNum == nil {
					currentNum = new(string)
				}
				*currentNum += string(char)
				numsByPosition[Position{Line: li + 1, Column: ci + 1}] = currentNum
			case "gear":
				currentNum = nil
				gearPositions = append(gearPositions, Position{Line: li + 1, Column: ci + 1})
			default:
				currentNum = nil
			}
		}
	}

	// Second pass:
	// - for every gear, count the unique numbers around it
	// - if the gear has exactly 2 numbers around it then add the numbers' product to the sum
	sum := 0
	for _, gear := range gearPositions {
		nums := make(map[*string]struct{})
		for l := gear.Line - 1; l <= gear.Line+1; l++ {
			for c := gear.Column - 1; c <= gear.Column+1; c++ {
				if num, ok := numsByPosition[Position{Line: l, Column: c}]; ok {
					nums[num] = struct{}{}
				}
			}
		}
		if len(nums) == 2 {
			product := 1
			for num := range nums {
				val, err := strconv.Atoi(*num)
				if err != nil {
					log.Fatalf("failed to convert string %q to number: %v\n", *num, err)
				}
				product *= val
			}
			sum += product
		}
	}

	log.Println("Sum:", sum)
}

// Get type of character; "digit", "gear" or "other".
func getCharType(c rune) string {
	if c == '*' {
		return "gear"
	}
	if c >= 48 && c <= 57 {
		return "digit"
	}
	return "other"
}
