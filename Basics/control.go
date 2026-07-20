package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a command line argument")
		return
	}

	argument := os.Args[1]

	// Switch statement with expression
	switch argument {
	case "0":
		fmt.Println("Zero")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("Two, Three, or Four!")
		fallthrough
	default:
		fmt.Println("Value:", argument)
	}

	value, err := strconv.Atoi(argument)
	if err != nil {
		fmt.Println("Couldn't convert to integer:", argument)
		return 
	}

	//Switch with no expression
	switch {
		case value == 0:
			fmt.Println("Zero")
		case value > 0:
			fmt.Println("Positive")
		case value < 0:
			fmt.Println("Negative")
		default:
			fmt.Println("Unknown value:", value)
	}
}
