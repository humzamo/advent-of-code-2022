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
	transformations := parseTransformations(inputList2)

	fmt.Println("The answer to part one is:", calculateAnswer(parseInitialStack(inputList1), transformations, partOne))
	fmt.Println("The answer to part two is:", calculateAnswer(parseInitialStack(inputList1), transformations, partTwo))
}

type stack map[string][]string

type transformation struct {
	count int
	from  string
	to    string
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

func calculateAnswer(stack stack, transformations []transformation, moveBlocks func(stack *stack, t transformation, blocks *[]string)) string {
	for _, t := range transformations {
		blocks := stack[t.from][len(stack[t.from])-t.count:]
		currentFrom := stack[t.from]
		stack[t.from] = currentFrom[:len(currentFrom)-t.count]

		moveBlocks(&stack, t, &blocks)
	}

	return stackTops(stack)
}

// move each block one at a time
func partOne(stack *stack, t transformation, blocks *[]string) {
	for i := t.count - 1; 0 <= i; i-- {
		currentTo := (*stack)[t.to]
		(*stack)[t.to] = append(currentTo, (*blocks)[i])
	}
}

// move each chunk of blocks in one go
func partTwo(stack *stack, t transformation, blocks *[]string) {
	currentTo := (*stack)[t.to]
	(*stack)[t.to] = append(currentTo, *blocks...)
}

// returns the top block of each stack in order
func stackTops(stack stack) string {
	result := ""
	for i := 1; i < len(stack)+1; i++ {
		key := strconv.Itoa(i)
		result += stack[key][len(stack[key])-1]
	}
	return result
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
		for i := len(s) - 2; i >= 0; i-- {
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
