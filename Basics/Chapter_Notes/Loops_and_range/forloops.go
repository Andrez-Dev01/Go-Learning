package main

import "fmt"

func main() {
	
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// Proper way to write in Go
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// For loop used as a while loop
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()

	// aSlice loop
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("Index:", i, "Value:", v)
	}
}