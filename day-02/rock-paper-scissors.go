package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputList := loadInputList("Input.txt")

	fmt.Println("The answer to part one is:", calculateScore(inputList, partOnestrategyToScore))
	fmt.Println("The answer to part one is:", calculateScore(inputList, partTwoStrategyToScore))
}

func loadInputList(inputFileName string) []string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var inputList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}
	defer file.Close()
	return inputList
}

const (
	rock     = 1
	paper    = 2
	scissors = 3
	loss     = 0
	draw     = 3
	win      = 6
)

var partOnestrategyToScore = map[string]int{
	"A X": rock + draw,
	"A Y": paper + win,
	"A Z": scissors + loss,
	"B X": rock + loss,
	"B Y": paper + draw,
	"B Z": scissors + win,
	"C X": rock + win,
	"C Y": paper + loss,
	"C Z": scissors + draw,
}

var partTwoStrategyToScore = map[string]int{
	"A X": scissors + loss,
	"A Y": rock + draw,
	"A Z": paper + win,
	"B X": rock + loss,
	"B Y": paper + draw,
	"B Z": scissors + win,
	"C X": paper + loss,
	"C Y": scissors + draw,
	"C Z": rock + win,
}

func calculateScore(strategies []string, strategyToScore map[string]int) int {
	score := 0
	for _, strategy := range strategies {
		score += strategyToScore[strategy]
	}
	return score
}
