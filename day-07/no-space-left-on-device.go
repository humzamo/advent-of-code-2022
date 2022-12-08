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
	tree.sumDirectory()
	// fmt.Println("used space:", tree.size)

	// fmt.Println(tree)

	fmt.Println("The answer to part one is:", tree.partOne())
	// fmt.Println("unused space:", totalDiskSpace-tree.size)

	// tree.partTwo()
	// fmt.Println(validDirectories)

	// sort.Slice(validDirectories, func(i, j int) bool {
	// 	return validDirectories[i] < validDirectories[j]
	// })
	// fmt.Println(validDirectories)

	fmt.Println("The answer to part two is:", tree.partTwo())

}

const (
	totalDiskSpace    = 70000000
	requiredDiskSpace = 30000000
	// totalUsedSpace    = 42677139
	// totalUsedSpace = 48381165 // example
	// totalUnusedSpace = 21618835 //example
	// totalUnusedSpace = 27322861
)

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
	size        int
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

		// either go back one level to the parent or go in one level
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

func (d *directory) directorySize() int {
	for _, directory := range d.directories {
		if directory.directories == nil {
			size := 0
			for _, f := range directory.files {
				size += f.size
			}
			directory.size = size
		}
		directory.size += directory.directorySize()
	}
	return d.size
}

func (d *directory) sumDirectory() int {
	sum := 0
	for _, f := range d.files {
		sum += f.size
	}
	d.size += sum
	for _, folder := range d.directories {
		d.size += folder.sumDirectory()
	}

	return d.size
}

func (d *directory) partOne() int {
	sum := 0
	for _, directory := range d.directories {
		if directory.size <= 100000 {
			sum += directory.size
		}
		if directory.directories != nil {
			sum += directory.partOne()
		}
	}
	return sum
}

var validDirectories = []int{}

func (d *directory) calculateValidDirectories(totalUnusedSpace int) []int {
	for _, directory := range d.directories {
		if totalUnusedSpace+directory.size >= requiredDiskSpace {
			validDirectories = append(validDirectories, directory.size)
		}
		if directory.directories != nil {
			directory.calculateValidDirectories(totalUnusedSpace)
		}
	}
	return validDirectories
}

func (d *directory) partTwo() int {
	totalUnusedSpace := totalDiskSpace - d.size
	directories := d.calculateValidDirectories(totalUnusedSpace)

	sort.Slice(directories, func(i, j int) bool {
		return directories[i] < directories[j]
	})
	return directories[0]
}
