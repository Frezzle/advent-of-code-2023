package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println(rune('0'))
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var sum int
	for fileScanner.Scan() {
		line := fileScanner.Text()
		firstDigit := getFirstDigit(line)
		lastDigit := getFirstDigit(reverse(line))

		num, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}

	log.Println("Sum:", sum)
}

func getFirstDigit(s string) string {
	for _, c := range s {
		if c >= 49 && c <= 57 {
			return string(c)
		}
	}
	return "0"
}

func reverse(s string) string {
	var reversed string
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}
