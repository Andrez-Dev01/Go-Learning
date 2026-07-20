// In this program we implement the function which(1) by dividing into three logical parts

// Part one is reading the argument input which is the name of the executable file the utility is searching for

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Part 2 logical part

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}

	file := arguments[1]                  // Read the argument of the program and save it into the file variable
	path := os.Getenv("PATH")             // get the contents of the PATH enviroment
	pathSplit := filepath.SplitList(path) // SPlit the file into seperate list
	for _, directory := range pathSplit { // Use a for loop to return into a slice
		fullPath := filepath.Join(directory, file)
		// Does it exist?
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()
		// A regular file?
		if !mode.IsRegular() {
			continue
		}

		// Is it executable
		if mode&0111 != 0 {
			fmt.Println(fullPath)
			return
		}
	}
}
