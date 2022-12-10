package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	grid := loadInput("Input.txt")

	// print out grid
	for _, l := range grid {
		fmt.Println(l)
	}
	fmt.Println("The answer to part one is:", partOne(grid))
}

func loadInput(inputFileName string) [][]int {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	var output [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := []int{}
		for _, s := range scanner.Text() {
			line = append(line, int(s-'0'))
		}
		output = append(output, line)
	}

	defer file.Close()
	return output
}

func partOne(grid [][]int) int {
	sum := 0

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if checkAbove(grid, i, j) || checkBelow(grid, i, j) || checkRight(grid, i, j) || checkLeft(grid, i, j) {
				sum++
			}

		}
	}
	sum += len(grid)*2 + len(grid[0])*2 - 4
	return sum
}

func checkAbove(grid [][]int, i, j int) bool {
	// fmt.Println(fmt.Sprintf("checking for i=%v, j=%v, val=%v", i, j, grid[i][j]))
	for i2 := 0; i2 <= i-1; i2++ {
		if grid[i][j] <= grid[i-1-i2][j] {
			return false
		}
	}
	return true
}

func checkBelow(grid [][]int, i, j int) bool {
	limit := len(grid) - 1 - i
	for i2 := 0; i2 < limit; i2++ {
		if grid[i][j] <= grid[i+1+i2][j] {
			return false
		}
	}
	return true
}

func checkLeft(grid [][]int, i, j int) bool {
	for j2 := 0; j2 <= j-1; j2++ {
		if grid[i][j] <= grid[i][j-1-j2] {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, i, j int) bool {
	limit := len(grid[0]) - 1 - j
	for j2 := 0; j2 < limit; j2++ {
		if grid[i][j] <= grid[i][j+1+j2] {
			return false
		}
	}
	return true
}
