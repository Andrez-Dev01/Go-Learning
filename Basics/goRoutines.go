// This program teaches synchromnizeing goroutines using time.Sleep() but this isn't the correct way just an example

package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

func main() {
	for i := 0; i < 4; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}
