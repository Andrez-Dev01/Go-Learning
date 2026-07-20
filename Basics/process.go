// This program shows different uses of error variables between different kinds of user input.
// We go from specific errors to general errors and making sure the same string is a float64 or an integer

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	// Examinging the valadity of the input
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			nInts++
			total++
			continue
		}
		// Is it a float?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}
		// If it's not an integer or a float, add it to the invalid slice
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}

}
