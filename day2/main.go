package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type (
	Game struct {
		ID    int
		Hands []Hand
	}

	Hand struct {
		Red   int
		Green int
		Blue  int
	}
)

func main() {
	games, err := readGames("input.txt")
	if err != nil {
		log.Fatalln("failed to read games:", err)
	}

	// part 1: calculate sum of IDs from possible games, given a number of cubes
	possibleRedCubes, possibleGreenCubes, possibleBlueCubes := 12, 13, 14
	sumOfIds := 0
	for _, game := range games {
		maxRed, maxGreen, maxBlue := maxColours(game)
		if maxRed <= possibleRedCubes && maxGreen <= possibleGreenCubes && maxBlue <= possibleBlueCubes {
			sumOfIds += game.ID
		}
	}
	log.Println("Sum of IDs:", sumOfIds)

	// part 2: calculate sum of powers of games' minimum needed colours
	sumOfPowers := 0
	for _, game := range games {
		maxRed, maxGreen, maxBlue := maxColours(game)
		sumOfPowers += maxRed * maxGreen * maxBlue
	}
	log.Println("Sum of powers:", sumOfPowers)
}

func readGames(file string) ([]Game, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open input file: %v", err)
	}

	var games []Game
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		id, err := strconv.Atoi(parts[0][5:])
		if err != nil {
			return nil, fmt.Errorf("failed to parse game id: %v", err)
		}

		var hands []Hand
		for _, handInput := range strings.Split(parts[1], "; ") {
			var hand Hand
			handParts := strings.Split(handInput, ", ")
			for _, part := range handParts {
				if numRed, isRed := strings.CutSuffix(part, " red"); isRed {
					hand.Red, err = strconv.Atoi(numRed)
					if err != nil {
						return nil, fmt.Errorf("failed to parse red count: %v", err)
					}
				} else if numGreen, isGreen := strings.CutSuffix(part, " green"); isGreen {
					hand.Green, err = strconv.Atoi(numGreen)
					if err != nil {
						return nil, fmt.Errorf("failed to parse green count: %v", err)
					}
				} else if numBlue, isBlue := strings.CutSuffix(part, " blue"); isBlue {
					hand.Blue, err = strconv.Atoi(numBlue)
					if err != nil {
						return nil, fmt.Errorf("failed to parse blue count: %v", err)
					}
				}
			}

			hands = append(hands, hand)
		}

		games = append(games, Game{
			ID:    id,
			Hands: hands,
		})
	}

	return games, nil
}

// maxColours returns the highest number of each cube colour seen in a game.
func maxColours(game Game) (red, green, blue int) {
	for _, hand := range game.Hands {
		if red < hand.Red {
			red = hand.Red
		}
		if green < hand.Green {
			green = hand.Green
		}
		if blue < hand.Blue {
			blue = hand.Blue
		}
	}
	return
}
