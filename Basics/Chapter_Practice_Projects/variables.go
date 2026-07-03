package main

import (
	"fmt"
	"math"
)

// Global variables
var global int = 1234
var anotherGlobal = -5678

func main() {

	var j int                          // locally declared variable
	i := global + anotherGlobal        // auto assigning and combining variables
	fmt.Println("Initial j value:", j) // printing variables after a println
	j = global                         // reassigning j variable to global variable

	k := math.Abs(float64(anotherGlobal))
	fmt.Printf("Global=%d, i=%d, j=%d k=%.2f.\n", global, i, j, k)
}
