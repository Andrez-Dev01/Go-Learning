package main

import (
	"fmt"
)

var AmountOfMoney, Amount int

func main() {

	AmountOfMoney = 100
	Amount = 25	
	i := AmountOfMoney - Amount
	fmt.Printf("AmountOfMoney=%d, Amount=%d, i=%d\n", AmountOfMoney, Amount, i)
}
