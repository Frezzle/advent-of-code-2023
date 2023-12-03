package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	digitWords = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	reverseDigitWords = map[string]string{
		"eno":   "1",
		"owt":   "2",
		"eerht": "3",
		"ruof":  "4",
		"evif":  "5",
		"xis":   "6",
		"neves": "7",
		"thgie": "8",
		"enin":  "9",
	}
)

type Digit struct {
	Value    string // digit, 1 to 9
	Location int    // index where it starts in string, 0 to len(line)-1
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var sum int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		firstDigit := getFirstDigit(line, digitWords)
		lastDigit := getFirstDigit(reverse(line), reverseDigitWords)

		num, err := strconv.Atoi(firstDigit.Value + lastDigit.Value)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	log.Println("Sum:", sum)
}

func getFirstDigit(s string, digitWords map[string]string) Digit {
	firstDigitChar := getFirstDigitCharacter(s)
	firstDigitWord := getFirstDigitWord(s, digitWords)
	if firstDigitChar.Location < firstDigitWord.Location {
		return firstDigitChar
	}
	return firstDigitWord
}

func getFirstDigitCharacter(s string) Digit {
	for i, c := range s {
		if c >= 49 && c <= 57 {
			return Digit{
				Value:    string(c),
				Location: i,
			}
		}
	}

	return Digit{
		Value:    "0",
		Location: math.MaxInt,
	}
}

func getFirstDigitWord(s string, digitWords map[string]string) Digit {
	first := Digit{
		Value:    "0",
		Location: math.MaxInt,
	}
	for digitWord, digitValue := range digitWords {
		if i := strings.Index(s, digitWord); i >= 0 && i < first.Location {
			first = Digit{
				Value:    digitValue,
				Location: i,
			}
		}
	}

	return first
}

func reverse(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
