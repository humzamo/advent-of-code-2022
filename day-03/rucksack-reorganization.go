package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputList := loadInputList("Input.txt")

	fmt.Println("The answer to part one is:", partOne(inputList))
	fmt.Println("The answer to part two is:", partTwo(inputList))
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

func partOne(input []string) int {
	sum := 0
	for _, s := range input {
		letter := findCommonLetterPartOne(s[0:len(s)/2], s[len(s)/2:])
		sum += letterValue(letter)
	}
	return sum
}

func partTwo(input []string) int {
	sum := 0
	for i := 0; i < len(input); {
		letter := findCommonLetterPartTwo(input[i], input[i+1], input[i+2])
		sum += letterValue(letter)
		i += 3
	}
	return sum
}

func findCommonLetterPartOne(s1, s2 string) rune {
	hash := make(map[rune]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			return e
		}
	}
	return 0
}

func findCommonLetterPartTwo(s1, s2, s3 string) rune {
	hash := make(map[rune]bool)
	hash2 := make(map[rune]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			hash2[e] = true
		}
	}
	for _, e := range s3 {
		if hash2[e] {
			return e
		}
	}
	return 0
}

func letterValue(l rune) int {
	// lowercase letter value, priorities 1 through 26
	if l > 96 {
		return int(l) - 96
	}

	// uppercase letter value, priorities 27 through 52
	return int(l) - 38
}
