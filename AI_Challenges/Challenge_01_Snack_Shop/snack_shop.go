/*
================================================================================
CHALLENGE 01: Snack Shop Checkout (Updated)
================================================================================

LESSON GOAL:
Practice problem-solving with everything you've learned so far:
  - variables, slices, arithmetic
  - for loops and range
  - if/else, switch
  - os.Args, strconv, and basic error handling
  - fmt.Printf formatted output

SCENARIO:
You run a snack shop. Customers order snacks and pay with cash.
Your program reads the order from command-line arguments, prints a receipt,
and reports change or amount still owed.

MENU (use slices — do NOT hard-code each item separately):
  Index 0: Chips  = $2.50
  Index 1: Soda   = $1.75
  Index 2: Candy  = $1.00

HOW TO RUN:
  go run snack_shop.go <chips> <sodas> <candies> <cash>

  Example:
  go run snack_shop.go 2 3 1 15.00

YOUR TASK:

1. SET UP SLICES
   - names  := []string{"Chips", "Soda", "Candy"}
   - prices := []float64{2.50, 1.75, 1.00}
   - Parse chips, sodas, candies from os.Args[1], [2], [3] using strconv.Atoi
   - Parse cash from os.Args[4] using strconv.ParseFloat
   - If len(os.Args) != 5, print usage and return
   - If any conversion fails, print an error and return

2. LOOP WITH RANGE
   - Build quantities slice from the parsed values
   - Use a for loop with range to:
     - calculate each line total (qty * price)
     - add to subtotal
     - count total items

3. DISCOUNT WITH SWITCH (on totalItems)
   - totalItems < 3  → 0% off
   - totalItems 3-4  → 5% off
   - totalItems >= 5 → 10% off

4. PAYMENT WITH SWITCH OR IF/ELSE
   - cash >= finalTotal → print change
   - cash < finalTotal  → print amount still owed
   - cash == finalTotal → print "Thank you! Exact payment!"

5. PRINT A RECEIPT with fmt.Printf (see example below)

EXAMPLE:
  $ go run snack_shop.go 2 3 1 15.00

  --- Snack Shop Receipt ---
  Chips: 2 x $2.50 = $5.00
  Soda:  3 x $1.75 = $5.25
  Candy: 1 x $1.00 = $1.00
  Subtotal:        $11.25
  Items: 6 (10% discount)
  Final total:     $10.13
  Cash given:      $15.00
  Change:          $4.87

HINTS:
  - Match the pattern from control.go for args + strconv + error checks
  - range over a slice: for i, qty := range quantities { ... }
  - Use prices[i] and names[i] inside the loop
  - switch with no expression (like control.go lines 38-47) works well for ranges

BONUS:
  - Reject negative quantities with an if inside the loop
  - Add a second loop that prints a mini sales bar chart using fmt.Print
    e.g. Chips: **  (one * per item sold)

================================================================================
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	names := []string{"Chips", "Soda", "Candy"}
	prices := []float64{2.50, 1.75, 1.00}

	fmt.Println("------Welcome to the snack shop!------")
	fmt.Println("Here's our menu")
	fmt.Println()
	for i, name := range names {
		fmt.Printf("%d: %s = $%.2f\n", i+1, name, prices[i])
	}
	fmt.Println()

	if len(os.Args) != 5 {
		fmt.Println("Usage: go run snack_shop.go <chips> <sodas> <candies> <cash>")
		return
	}

	chips, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Chips Error:", err)
		return
	}

	sodas, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Sodas Error:", err)
		return
	}

	candies, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Candies Error:", err)
		return
	}

	cash, err := strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		fmt.Println("Cash Error:", err)
		return
	}

	quantities := []int{chips, sodas, candies}

	subtotal := 0.0
	totalItems := 0
	for i, qty := range quantities {
		if qty < 0 {
			fmt.Printf("Error: %s quantity cannot be negative\n", names[i])
			return
		}
		totalItems += qty
		subtotal += float64(qty) * prices[i]
	}

	discountRate := 0.0
	switch {
	case totalItems < 3:
		discountRate = 0.0
	case totalItems <= 4:
		discountRate = 0.05
	default:
		discountRate = 0.10
	}

	finalTotal := subtotal * (1 - discountRate)

	fmt.Println("--- Snack Shop Receipt ---")
	for i, qty := range quantities {
		fmt.Printf("%s: %d x $%.2f = $%.2f\n", names[i], qty, prices[i], float64(qty)*prices[i])
	}
	fmt.Printf("Subtotal:        $%.2f\n", subtotal)
	fmt.Printf("Items: %d (%d%% discount)\n", totalItems, int(discountRate*100))
	fmt.Printf("Final total:     $%.2f\n", finalTotal)
	fmt.Printf("Cash given:      $%.2f\n", cash)

	switch {
	case cash == finalTotal:
		fmt.Println("Thank you! Exact payment!")
	case cash > finalTotal:
		fmt.Printf("Change:          $%.2f\n", cash-finalTotal)
	case cash < finalTotal:
		fmt.Printf("Still owe:       $%.2f\n", finalTotal-cash)
	}

	fmt.Println("-----Thank you for your purchase!-----")
}
