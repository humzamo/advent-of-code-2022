package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	grid := loadInput("Input.txt")

	fmt.Println("The answer to part one is:", partOne(grid))
	fmt.Println("The answer to part two is:", partTwo(grid))
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
			if checkAboveP1(grid, i, j) || checkBelowP1(grid, i, j) || checkRightP1(grid, i, j) || checkLeftP1(grid, i, j) {
				sum++
			}

		}
	}
	// add the trees at the boundary
	sum += len(grid)*2 + len(grid[0])*2 - 4
	return sum
}

func partTwo(grid [][]int) int {
	heights := []int{}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			heights = append(heights, checkAboveP2(grid, i, j)*checkBelowP2(grid, i, j)*checkLeftP2(grid, i, j)*checkRightP2(grid, i, j))
		}
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	return heights[0]
}

func checkAboveP1(grid [][]int, i, j int) bool {
	for i2 := 0; i2 <= i-1; i2++ {
		if grid[i][j] <= grid[i-1-i2][j] {
			return false
		}
	}
	return true
}

func checkBelowP1(grid [][]int, i, j int) bool {
	limit := len(grid) - 1 - i
	for i2 := 0; i2 < limit; i2++ {
		if grid[i][j] <= grid[i+1+i2][j] {
			return false
		}
	}
	return true
}

func checkLeftP1(grid [][]int, i, j int) bool {
	for j2 := 0; j2 <= j-1; j2++ {
		if grid[i][j] <= grid[i][j-1-j2] {
			return false
		}
	}
	return true
}

func checkRightP1(grid [][]int, i, j int) bool {
	limit := len(grid[0]) - 1 - j
	for j2 := 0; j2 < limit; j2++ {
		if grid[i][j] <= grid[i][j+1+j2] {
			return false
		}
	}
	return true
}

func checkAboveP2(grid [][]int, i, j int) int {
	i2 := 0
	for i2 <= i-1 {
		if grid[i][j] <= grid[i-1-i2][j] {
			break
		}
		i2++
	}

	if i == i2 {
		return i2
	}
	return i2 + 1
}

func checkBelowP2(grid [][]int, i, j int) int {
	limit := len(grid) - 1 - i
	i2 := 0
	for i2 < limit {
		if grid[i][j] <= grid[i+1+i2][j] {
			break
		}
		i2++
	}

	if i2 == limit {
		return i2
	}
	return i2 + 1
}

func checkLeftP2(grid [][]int, i, j int) int {
	j2 := 0
	for j2 <= j-1 {
		if grid[i][j] <= grid[i][j-1-j2] {
			break
		}
		j2++
	}

	if j2 == j {
		return j2
	}
	return j2 + 1
}

func checkRightP2(grid [][]int, i, j int) int {
	j2 := 0
	limit := len(grid[0]) - 1 - j
	for j2 < limit {
		if grid[i][j] <= grid[i][j+1+j2] {
			break
		}
		j2++
	}

	if j2 == limit {
		return j2
	}
	return j2 + 1
}
