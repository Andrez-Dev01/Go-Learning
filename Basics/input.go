// fmt.Scanln() this function helps read input from the user and store it in a variable and is passed as a pointer to the variable
// But this function is rarely used for user input  because we usually read from command line arguments or from a file


// THe fmt package also conataints adition functions for reading input from the console with os.Stdin

// In this program we read from standard input 

package main 

import (
	"fmt"
)

func main() {
	// Getting User Input
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello, ", name, "! Nice to meet you.")

}