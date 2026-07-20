// Command Line Arguments

/* When we want user input we usually read from command line arguments
 * By default the command line arguments in are stored in the os.Args slice
 */

 // In this program  the use of os.Args is demonstrated to find the maximum and minimum of a list of numbers while ignoring inlavid input like characters or strings

package main

import (
	"fmt" // for printing to the console
	"os" // for accessing command line arguments
	"strconv" // for converting strings to integers
)

func main() {
	// We make sure that we have at least one command line argument
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return // Exit the program if no arguments are provided
	}

	var min, max float64
	var initialized = 0
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if initialized == 0 {
			min = n
			max = n 
			initialized = 1
			continue
		}
		if n < min {
			min = n 
		}
		if n > max {
			max = n
		}
	}
		fmt.Println("Min:", min, "Max:", max)
}