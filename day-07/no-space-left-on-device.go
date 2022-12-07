package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputList := loadInputList("Input.txt")
	tree := parseDirectory(inputList)
	fmt.Println(tree)

	// fmt.Println("The answer to part one is:", partOne(assignmentList))
	// fmt.Println("The answer to part two is:", partTwo(assignmentList))
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

type file struct {
	name string
	size int
}

type directory struct {
	name        string
	parent      *directory
	files       []file
	directories []*directory
}

func parseDirectory(input []string) directory {
	// start off at the root
	currentDir := &directory{
		name: "/",
	}

	for i, s := range input {
		if i == 0 {
			continue
		}

		// either cd back one directory or go in one lever
		if s[0:4] == "$ cd" {
			if (s[5:]) == ".." {
				currentDir = currentDir.parent
			} else {
				currentDir = currentDir.moveIntoDirectory(s[5:])
			}
			continue
		}

		// after an ls, add everything to the current directory
		if s[0:4] == "$ ls" {
			continue
		}

		splitString := strings.Split(s, " ")
		if splitString[0] == "dir" {
			dirs := currentDir.directories
			currentDir.directories = append(dirs, &directory{
				name:   splitString[1],
				parent: currentDir,
			})
		} else {
			size, _ := strconv.Atoi(splitString[0])
			files := currentDir.files
			currentDir.files = append(files, file{
				name: splitString[1],
				size: size,
			})
		}
	}
	currentDir = currentDir.goToRoot()
	return *currentDir
}

func (d *directory) moveIntoDirectory(name string) *directory {
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
