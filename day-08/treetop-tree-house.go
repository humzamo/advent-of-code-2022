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

	// print out grid
	// for _, l := range grid {
	// 	fmt.Println(l)
	// }
	// fmt.Println("The answer to part one is:", partOne(grid))
	fmt.Println("The answer to part one is:", partTwo(grid))
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

// func partOne(grid [][]int) int {
// 	sum := 0

// 	for i := 1; i < len(grid)-1; i++ {
// 		for j := 1; j < len(grid[0])-1; j++ {
// 			if checkAbove(grid, i, j) || checkBelow(grid, i, j) || checkRight(grid, i, j) || checkLeft(grid, i, j) {
// 				sum++
// 			}

// 		}
// 	}
// 	sum += len(grid)*2 + len(grid[0])*2 - 4
// 	return sum
// }

func partTwo(grid [][]int) int {
	heights := []int{}

	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			// if i == 3 && j == 2 {
			// fmt.Println(fmt.Sprintf("checking for i=%v, j=%v, val=%v", i, j, grid[i][j]))
			heights = append(heights, checkAbove(grid, i, j)*checkBelow(grid, i, j)*checkLeft(grid, i, j)*checkRight(grid, i, j))
			// }
		}
	}
	sort.Slice(heights, func(i, j int) bool {
		return heights[i] > heights[j]
	})
	return heights[0]
}

func checkAbove(grid [][]int, i, j int) int {
	i2 := 0
	for i2 <= i-1 {
		// fmt.Println(fmt.Sprintf("looking at for i=%v, j=%v, val=%v", i-1-i2, j, grid[i-1-i2][j]))

		if grid[i][j] <= grid[i-1-i2][j] {
			fmt.Println("blocking tree is", grid[i-1-i2][j])

			break
		}
		i2++
	}
	if i-1-i2 == -1 {
		fmt.Println("at the end so above is", i2)
		return i2
	}
	fmt.Println("above is", i2+1)
	fmt.Println("we go up to is", i-1-i2)

	return i2 + 1
}

func checkBelow(grid [][]int, i, j int) int {
	limit := len(grid) - 1 - i
	i2 := 0
	for i2 < limit {
		if grid[i][j] <= grid[i+1+i2][j] {
			break
		}
		i2++
	}

	if i+1+i2 == len(grid) {
		fmt.Println("below is", i2)
		return i2
	}
	fmt.Println("below is", i2+1)

	return i2 + 1
}

func checkLeft(grid [][]int, i, j int) int {
	j2 := 0
	for j2 <= j-1 {
		if grid[i][j] <= grid[i][j-1-j2] {
			break
		}
		j2++
	}

	if j-1-j2 == -1 {
		fmt.Println("left is", j2)
		return j2
	}
	fmt.Println("left is", j2+1)

	return j2 + 1
}

func checkRight(grid [][]int, i, j int) int {
	j2 := 0
	limit := len(grid[0]) - 1 - j
	for j2 < limit {
		if grid[i][j] <= grid[i][j+1+j2] {
			break
		}
		j2++
	}
	fmt.Println("we go up to is", j+1+j2)
	fmt.Println(len(grid))

	if j+1+j2 == len(grid) {
		fmt.Println("at end so right is", j2)
		return j2
	}
	fmt.Println("right is", j2+1)

	return j2 + 1
}
