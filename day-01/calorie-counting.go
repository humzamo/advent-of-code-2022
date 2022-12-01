package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	numberList := loadNumbersList("Input.txt")
	caloriesList := caloriesList(numberList)

	partOneAnswer := partOne(caloriesList)
	fmt.Println("The answer to part one is:", partOneAnswer)

	partTwoAnswer := partTwo(caloriesList)
	fmt.Println("The answer to part one is:", partTwoAnswer)
}

func loadNumbersList(inputFileName string) []int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var numberList []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			numberList = append(numberList, 0)
			continue
		}
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("The input list contains a value which is not an integer!")
			log.Fatal(err)
		}
		numberList = append(numberList, number)
	}
	defer file.Close()
	return numberList
}

func caloriesList(list []int) []int {
	calories := []int{}
	sum := 0
	for _, num := range list {
		if num == 0 {
			calories = append(calories, sum)
			sum = 0
			continue
		}
		sum += num
	}
	return calories
}

// calculate max of calories list
func partOne(list []int) int {
	m := 0
	for i, e := range list {
		if i == 0 || e > m {
			m = e
		}
	}
	return m
}

// calculate sum of highest three calories
func partTwo(list []int) int {
	sort.Ints(list)
	topThreeItems := list[len(list)-3:]
	return topThreeItems[0] + topThreeItems[1] + topThreeItems[2]
}
