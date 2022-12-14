package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Directory struct {
	Parent *Directory
	Dirs   []*Directory
	Name   string
	Size   int
}

const TAB = "  "
const TOTAL_SIZE = 70_000_000
const UPDATE_SIZE = 30_000_000

func (d *Directory) GetDirectories() (dirs []*Directory) {
	if len(d.Dirs) != 0 {
		dirs = append(dirs, d)
	}
	for _, d2 := range d.Dirs {
		dirs = append(dirs, d2.GetDirectories()...)
	}
	return dirs
}

func (d *Directory) ToString(indentation int) (ret string) {
	for i := 0; i < indentation; i++ {
		ret += fmt.Sprintf("  ")
	}
	indentation++
	if len(d.Dirs) != 0 {
		ret += fmt.Sprintf("- %s (dir, size=%d)\n", d.Name, d.Size)
	} else {
		ret += fmt.Sprintf("- %s: [file, size=%d]\n", d.Name, d.Size)
	}
	for _, f := range d.Dirs {
		ret += fmt.Sprintf("%s", f.ToString(indentation))
	}
	return ret
}

func (d *Directory) IncrementSize(size int) {
	if d.Parent != nil {
		d.Size += size
		d.Parent.IncrementSize(d.Size)
	}
}

func (d *Directory) Touch(dir Directory) {
	dir.Dirs = []*Directory{}
	dir.Parent = d
	d.Dirs = append(d.Dirs, &dir)
}

func LineToFile(line string) Directory {
	// file
	split := strings.Split(line, " ")
	pattern := regexp.MustCompile("[0-9]+")
	size, _ := strconv.Atoi(pattern.FindAllString(split[0], -1)[0])
	name := split[1]
	return Directory{
		Name: name,
		Size: size,
	}
}

func LineToDir(line string) Directory {
	dirName := strings.Split(line, "dir ")[1]
	return Directory{
		Name: dirName,
	}
}

func (d Directory) Find(dirName string) *Directory {
	for _, dir := range d.Dirs {
		if dir.Name == dirName {
			return dir
		}
	}
	return nil
}

func readFile() (Directory, error) {
	file, err := os.Open("./input.txt")
	if err != nil {
		return Directory{}, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var currentDirectory *Directory
	rootDirectory := Directory{
		Dirs: []*Directory{},
		Name: "/",
		Size: 0,
	}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "$") {
			command := strings.Split(line, "$ ")[1]
			if command == "ls" {
				continue
			} else {
				dirToCdTo := strings.Split(command, " ")[1]
				// comand is cd
				// set pointer
				if dirToCdTo == "/" {
					currentDirectory = &rootDirectory
				} else if dirToCdTo == ".." {
					currentDirectory = currentDirectory.Parent
				} else {
					currentDirectory = currentDirectory.Find(dirToCdTo)
				}
			}
		} else {
			if strings.Contains(line, "dir") {
				currentDirectory.Touch(LineToDir(line))
			} else {
				file := LineToFile(line)
				currentDirectory.Touch(file)
				currentDirectory.IncrementSize(file.Size)
				rootDirectory.Size += file.Size
			}
		}
	}
	return rootDirectory, nil
}

func main() {
	data, _ := readFile()
	currentSize := data.Size
	freeSpace := TOTAL_SIZE - currentSize
	spaceToDelete := UPDATE_SIZE - freeSpace
	directories := data.GetDirectories()
	current := TOTAL_SIZE
	// fmt.Printf("%s", data.ToString(0))
	fmt.Printf("freeSpace: %d\n", freeSpace)
	fmt.Printf("spaceToDelete: %d\n", spaceToDelete)
	for _, d := range directories {
		if spaceToDelete <= d.Size && d.Size <= current {
			current = d.Size
		}
	}
	fmt.Printf("Smallest to delete: %d", current)
}
