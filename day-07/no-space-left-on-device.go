package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputList := loadInputList("Input.txt")
	tree := parseDirectory(inputList)

	fmt.Println("The answer to part one is:", tree.partOne())
	fmt.Println("The answer to part two is:", tree.partTwo())

}

const (
	directoryLimit    = 100000
	totalDiskSpace    = 70000000
	requiredDiskSpace = 30000000
)

var (
	validDirectories = []int{}
)

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	parent      *directory
	files       []file
	directories []*directory
	size        int
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

func parseDirectory(input []string) directory {
	// start off at the root
	currentDir := &directory{
		name: "/",
	}

	for i, s := range input {
		// skip the first iteration as we've already added the root directory
		if i == 0 {
			continue
		}

		// either go back one level to the parent or go in one level
		if s[0:4] == "$ cd" {
			directoryName := s[5:]
			if directoryName == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.enterDirectory(directoryName)
			}
			continue
		}

		// after an ls, add everything to the current directory
		if s[0:4] == "$ ls" {
			continue
		}

		splitString := strings.Split(s, " ")
		if splitString[0] == "dir" {
			currentDir.directories = append(currentDir.directories, &directory{
				name:   splitString[1],
				parent: currentDir,
			})
		} else {
			size, _ := strconv.Atoi(splitString[0])
			currentDir.files = append(currentDir.files, file{
				name: splitString[1],
				size: size,
			})
		}
	}
	currentDir = currentDir.goToRoot()
	currentDir.sumDirectory()
	return *currentDir
}

func (d *directory) partOne() int {
	sum := 0
	for _, directory := range d.directories {
		if directory.size <= directoryLimit {
			sum += directory.size
		}
		if directory.directories != nil {
			sum += directory.partOne()
		}
	}
	return sum
}

func (d *directory) partTwo() int {
	totalUnusedSpace := totalDiskSpace - d.size
	d.calculateValidDirectories(totalUnusedSpace)

	sort.Slice(validDirectories, func(i, j int) bool {
		return validDirectories[i] < validDirectories[j]
	})
	return validDirectories[0]
}

func (d *directory) enterDirectory(name string) *directory {
	for _, directory := range d.directories {
		if directory.name == name {
			return directory
		}
	}
	return nil
}

func (d *directory) goToRoot() *directory {
	for {
		if d.parent == nil {
			return d
		}
		d = d.parent
		d.goToRoot()
	}
}

// calculates the size of each directory in the tree
func (d *directory) sumDirectory() int {
	for _, f := range d.files {
		d.size += f.size
	}

	for _, folder := range d.directories {
		d.size += folder.sumDirectory()
	}

	return d.size
}

func (d *directory) calculateValidDirectories(totalUnusedSpace int) {
	for _, directory := range d.directories {
		if totalUnusedSpace+directory.size >= requiredDiskSpace {
			validDirectories = append(validDirectories, directory.size)
		}
		if directory.directories != nil {
			directory.calculateValidDirectories(totalUnusedSpace)
		}
	}
}
