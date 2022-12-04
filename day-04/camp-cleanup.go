package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputList := loadInputList("Input.txt")

	// fmt.Println("The answer to part one is:", partOne(inputList))
	// fmt.Println("The answer to part two is:", partTwo(inputList))
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

type assignment struct{
	start int
	end int
}

func parseAssignment(input []string) [][]assignment{
	assignmentsList := [][]assignments{}
	for _,s := range input {
		splitString := strings.Split(s,",")
		assignment1Strings := strings.Split(splitString[0],"-")
		assignment2Strings := strings.Split(splitString[1],"-")

		assignment1 := assignment{
			start: assignment1Strings[0],
			end: assignment1Strings[1]
		}

		assignment2 := assignment{
			start: assignment2Strings[0],
			end: assignment2Strings[1]
		}

		assignments = append(assignments, []assignment{assignment1,assignment2})
	}
	return assignmentsList
}