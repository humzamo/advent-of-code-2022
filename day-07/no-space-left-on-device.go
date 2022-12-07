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
	// fmt.Println(inputList)
	// assignmentList := parseAssignments(inputList)

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
	parentName  string
	files       []file
	directories []*directory
}

func parseDirectory(input []string) directory {
	d := directory{
		name: "/",
	}
	currentDir := &d
	tempDir := &directory{}

	for i, s := range input {
		if i == 0 {
			continue
		}
		// fmt.Println(d.files)
		// if the line is an instruction
		if s[0:4] == "$ cd" {
			currentDir = currentDir.insert(tempDir)
			tempDir = &directory{}

			if s[4:6] == ".." {
				currentDir = currentDir.findNode(currentDir.parentName)
			} else {
				currentDir = currentDir.findNode(s[5:])
			}
			continue
		}

		// after an ls, add everything to the current directory
		if s[0:4] == "$ ls" {
			continue
		}

		splitString := strings.Split(s, " ")
		// fmt.Println(splitString)
		if splitString[0] == "dir" {
			dirs := tempDir.directories
			tempDir.directories = append(dirs, &directory{
				name:       splitString[1],
				parentName: currentDir.name,
			})
		} else {
			size, _ := strconv.Atoi(splitString[0])
			files := tempDir.files
			tempDir.files = append(files, file{
				name: splitString[1],
				size: size,
			})
		}
		fmt.Println(tempDir)
	}
	return *currentDir
}

func (root *directory) insert(d *directory) *directory {
	initialDir := root.directories
	root.directories = append(initialDir, d)

	initialFiles := root.files
	root.files = append(initialFiles, d.files...)
	return root
}

func (node *directory) findNode(dirName string) *directory {
	if node == nil {
		return nil
	}
	if node.name == dirName {
		return node
	}
	for _, child := range node.directories {
		child.findNode(dirName)
	}

	return node
}
