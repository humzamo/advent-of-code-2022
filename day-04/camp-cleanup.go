package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputList := loadInputList("Input.txt")
	assignmentList := parseAssignments(inputList)

	fmt.Println("The answer to part one is:", partOne(assignmentList))
	fmt.Println("The answer to part two is:", partTwo(assignmentList))
}

func loadInputList(inputFileName string) []string {
	// file, err := os.Open(inputFileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	reader := strings.NewReader(`2-4,6-8
	2-3,4-5
	5-7,7-9
	2-8,3-7
	6-6,4-6
	2-6,4-8`)

	var inputList []string
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	// defer file.Close()
	return inputList
}

type assignment struct {
	start int
	end   int
}

func parseAssignments(input []string) [][]assignment {
	assignmentsList := [][]assignment{}
	for _, s := range input {
		s = strings.TrimSpace(s)
		splitString := strings.Split(s, ",")
		assignment1Strings := strings.Split(splitString[0], "-")
		assignment2Strings := strings.Split(splitString[1], "-")

		assignment1start, _ := strconv.Atoi(assignment1Strings[0])
		assignment1end, _ := strconv.Atoi(assignment1Strings[1])
		assignment2start, _ := strconv.Atoi(assignment2Strings[0])
		assignment2end, _ := strconv.Atoi(assignment2Strings[1])

		assignment1 := assignment{
			start: assignment1start,
			end:   assignment1end,
		}

		assignment2 := assignment{
			start: assignment2start,
			end:   assignment2end,
		}

		assignmentsList = append(assignmentsList, []assignment{assignment1, assignment2})
	}
	return assignmentsList
}

func partOne(assignmentList [][]assignment) int {
	sum := 0
	for _, assignments := range assignmentList {
		// if first assignment is contained in second assignment
		if assignments[0].start <= assignments[1].start && assignments[1].end <= assignments[0].end {
			sum++
			continue
		}

		// if second assignment is contained in first assignment
		if assignments[1].start <= assignments[0].start && assignments[0].end <= assignments[1].end {
			sum++
			continue
		}
	}
	return sum
}

func partTwo(assignmentList [][]assignment) int {
	sum := 0
	for _, assignments := range assignmentList {
		// if there is any overlap between assignments
		if assignments[0].end >= assignments[1].start && assignments[1].end >= assignments[0].start {
			sum++
		}
	}
	return sum
}
