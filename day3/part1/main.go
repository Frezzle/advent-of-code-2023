package main

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Number struct {
	Value        string // e.g. "42"
	IsEnginePart bool   // whether to count it towards sum
}

// The position of something in a text file e.g. a number or symbol.
type Position struct {
	Line   int // e.g. line 1, 2, 3, ...
	Column int // e.g. column 1, 2, 3, ...
}

func main() {
	// First pass:
	// - get numbers and their positions (numbers can be repeated)
	// - get symbol positions
	numsByPosition := make(map[Position]*Number)
	var symbolPositions []Position
	lines := strings.Split(input, "\n")
	for li, line := range lines {
		var currentNum *Number
		for ci, char := range line {
			charType := getCharType(char)
			switch charType {
			case "digit":
				if currentNum == nil {
					currentNum = &Number{Value: string(char)}
				} else {
					currentNum.Value += string(char)
				}
				numsByPosition[Position{Line: li + 1, Column: ci + 1}] = currentNum
			case "symbol":
				currentNum = nil
				symbolPositions = append(symbolPositions, Position{Line: li + 1, Column: ci + 1})
			default:
				currentNum = nil
			}
		}
	}

	// Second pass:
	// - for every symbol, go through every position around it
	// - if the position has a number occupying it then mark it as being an engine part
	for _, symbol := range symbolPositions {
		for l := symbol.Line - 1; l <= symbol.Line+1; l++ {
			for c := symbol.Column - 1; c <= symbol.Column+1; c++ {
				if num, ok := numsByPosition[Position{Line: l, Column: c}]; ok {
					num.IsEnginePart = true
				}
			}
		}
	}

	// We have many positions that point to the same number,
	// so we must first grab all unique numbers to be able to count them properly.
	nums := make(map[*Number]struct{})
	for _, num := range numsByPosition {
		nums[num] = struct{}{}
	}

	// Finally, sum all engine parts together.
	sum := 0
	for num := range nums {
		if num.IsEnginePart {
			val, err := strconv.Atoi(num.Value)
			if err != nil {
				log.Fatalf("failed to convert string %q to number: %v\n", num.Value, err)
			}
			sum += val
		}
	}

	log.Println("Sum:", sum)
}

// Get type of character; "digit", "symbol" or "empty".
func getCharType(c rune) string {
	if c == '.' {
		return "empty"
	}
	if c >= 48 && c <= 57 {
		return "digit"
	}
	return "symbol"
}
