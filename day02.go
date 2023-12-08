package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type RoundData struct {
	red   int64
	green int64
	blue  int64
}

func newRoundData() RoundData {
	return RoundData{
		red:   0,
		green: 0,
		blue:  0,
	}
}

func extractRoundData(rawRound string) RoundData {
	re, _ := regexp.Compile(`(\d+) (\w+)`)
	groups := re.FindAllStringSubmatch(rawRound, -1)

	roundData := newRoundData()

	for _, found := range groups {
		value, _ := strconv.ParseInt(found[1], 10, 32)
		ballColor := found[2]

		switch color := ballColor; color {
		case "red":
			roundData.red = value
		case "green":
			roundData.green = value
		case "blue":
			roundData.blue = value
		default:
			fmt.Printf("Incorrect ball color %s.\n", ballColor) // better logger error
		}
	}

	return roundData
}

func validateRound(data RoundData, limit RoundData) bool {
	if data.red > limit.red {
		return false
	}
	if data.green > limit.green {
		return false
	}
	if data.blue > limit.blue {
		return false
	}

	return true
}

func processingGame(gameContent string, gameLimit RoundData) bool {
	// gameContent="1 red, 5 blue; 5 green; 6 green, 8 blue, 2 red; 1 red, 6 blue, 6 green"
	rawRounds := strings.Split(gameContent, ";")

	for _, rawRound := range rawRounds {
		roundData := extractRoundData(rawRound)

		isRoundValid := validateRound(roundData, gameLimit)

		if !isRoundValid {
			return false
		}
	}
	return true
}

func day02() {
	readFile, err := os.Open("inputs/2.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	result := int64(0)
	gameLimit := RoundData{
		red:   12,
		green: 13,
		blue:  14,
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		re, _ := regexp.Compile(`Game (\d+):(.*)`)
		groups := re.FindAllStringSubmatch(line, -1)

		gameNumber, _ := strconv.ParseInt(groups[0][1], 10, 32)
		gameContent := groups[0][2]

		if processingGame(gameContent, gameLimit) {
			result += gameNumber
		}
	}

	fmt.Println("result:", result)

	readFile.Close()
}
