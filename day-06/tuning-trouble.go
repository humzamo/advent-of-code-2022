package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	signalString := loadInput("Input.txt")

	fmt.Println("The answer to part one is:", calculateMarker(signalString, 4))
	fmt.Println("The answer to part two is:", calculateMarker(signalString, 14))
}

func loadInput(inputFileName string) string {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func calculateMarker(signal string, limit int) int {
	for i := 0; i < len(signal); i++ {
		if hasDuplicates(signal[i : i+limit]) {
			return i + limit
		}
	}
	return 0
}

func hasDuplicates(str string) bool {
	var bits = make(map[rune]bool)
	for _, char := range str {
		if bits[char] {
			return false
		}
		bits[char] = true
	}
	return true
}
