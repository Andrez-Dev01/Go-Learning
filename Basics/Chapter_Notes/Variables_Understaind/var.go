package main

import (
	"fmt"
)

func main() {
	//Go has different ways to declare variables
	//If no value give Go complier will give it  0 value
	//This method is more for declaring global and local variables without initial value

	//var type
	var string = "Hello World!"
	fmt.Println(string)

	// := short assignment statement
	//Most used in returning values from functions and for loops with the range keyword
	//This method cannot be used outside of a function because it does not permit it there
	i := 12
	fmt.Println(i)
}
