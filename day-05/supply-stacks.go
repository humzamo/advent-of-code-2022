package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputList1, inputList2 := loadInputList("Input.txt")

	initialStack := parseInitialStack(inputList1)
	transformations := parseTransformations(inputList2)

	fmt.Println(initialStack)
	fmt.Println(transformations)

	// fmt.Println("The answer to part one is:", partOne(assignmentList))
	// fmt.Println("The answer to part two is:", partTwo(assignmentList))
}

func loadInputList(inputFileName string) ([]string, []string) {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var inputList, transformationList []string
	finishedFirstParse := false
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			finishedFirstParse = true
			continue
		}
		if !finishedFirstParse {
			inputList = append(inputList, scanner.Text())
		} else {
			transformationList = append(transformationList, scanner.Text())
		}
	}

	defer file.Close()
	return inputList, transformationList
}

// Eg 1: A B C
type stack map[string][]string

type transformation struct {
	count int
	from  string
	to    string
}

func parseTransformations(input []string) []transformation {
	transformations := []transformation{}
	for _, t := range input {
		s := strings.Split(t, " ")
		count, _ := strconv.Atoi(s[1])
		transformations = append(transformations, transformation{
			count: count,
			from:  s[3],
			to:    s[5],
		})
	}
	return transformations
}

func parseInitialStack(input []string) stack {
	stack := stack{}
	original := [][]string{}
	for _, s := range input {
		split := chunkString(s, 4)
		original = append(original, split)
	}

	transposed := transpose(original)
	for _, s := range transposed {
		arr := []string{}
		for i := len(s) - 2; i > 0; i-- {
			if s[i] == "" {
				break
			}
			arr = append(arr, s[i])
		}
		stack[s[len(s)-1]] = arr
	}
	return stack
}

var re = regexp.MustCompile(`[^a-zA-Z0-9]`)

func chunkString(s string, chunkSize int) []string {
	var chunks []string
	runes := []rune(s)

	if len(runes) == 0 {
		return []string{s}
	}

	for i := 0; i < len(runes); i += chunkSize {
		nn := i + chunkSize
		if nn > len(runes) {
			nn = len(runes)
		}
		s := string(runes[i:nn])
		s = re.ReplaceAllString(s, "")
		chunks = append(chunks, s)
	}
	return chunks
}

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
