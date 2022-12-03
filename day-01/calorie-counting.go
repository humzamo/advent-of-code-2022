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

	fmt.Println("The answer to part one is:", topCalories(caloriesList, 1))
	fmt.Println("The answer to part two is:", topCalories(caloriesList, 3))
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

	sort.Slice(calories, func(i, j int) bool {
		return calories[i] > calories[j]
	})
	return calories
}

func topCalories(list []int, limit int) int {
	sum := 0
	for i := 0; i < limit; i++ {
		sum += list[i]
	}
	return sum
}
