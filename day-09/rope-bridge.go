package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	instructions := loadInputList("Input.txt")
	fmt.Println(instructions)

	// fmt.Println("The answer to part one is:", partOne(assignmentList))
	// fmt.Println("The answer to part two is:", partTwo(assignmentList))
}

func loadInputList(inputFileName string) []instruction {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var instructions []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), " ")
		steps, _ := strconv.Atoi(splitText[1])
		instructions = append(instructions, instruction{
			motion: splitText[0],
			steps:  steps,
		})
	}

	defer file.Close()
	return instructions
}

type instruction struct {
	motion string // left, right, up, down
	steps  int
}
