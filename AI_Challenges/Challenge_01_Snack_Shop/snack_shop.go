/*
================================================================================
CHALLENGE 01: Snack Shop Checkout
================================================================================

LESSON GOAL:
Practice problem-solving with variables, arithmetic, if/else, and formatted
output — using only what you've learned so far in this repo.

SCENARIO:
You run a small snack shop. A customer buys snacks and pays with cash.
Your program should calculate the total cost, check if they paid enough,
and print either the change or how much more they still owe.

RULES:
  - Chips  = $2.50 each
  - Soda   = $1.75 each
  - Candy  = $1.00 each

  - If the customer buys 5 or more items TOTAL (all snacks combined),
    they get 10% off the subtotal.
  - If they pay with exact cash (change == 0), print a thank-you message.

YOUR TASK:
1. Declare variables for how many of each snack the customer buys.
2. Declare a variable for how much cash the customer gives you.
3. Calculate:
   - subtotal (before discount)
   - total items
   - discount (if applicable)
   - final total (after discount)
4. Use if/else to:
   - Print the change if cash >= final total
   - Print how much more they owe if cash < final total
5. Use fmt.Printf to show a clear receipt (see example below).

EXAMPLE OUTPUT (when they have enough money):
  --- Snack Shop Receipt ---
  Chips: 2 x $2.50 = $5.00
  Soda:  3 x $1.75 = $5.25
  Candy: 1 x $1.00 = $1.00
  Subtotal:        $11.25
  Items: 6 (10% discount applied!)
  Final total:     $10.13
  Cash given:      $15.00
  Change:          $4.87
  Thank you! Come again!

EXAMPLE OUTPUT (when they don't have enough):
  --- Snack Shop Receipt ---
  ...
  Final total:     $10.13
  Cash given:      $8.00
  Still owe:       $2.13

HINTS:
  - Use float64 for prices and money (e.g. 2.50, not 2.5 as int).
  - totalItems := chips + sodas + candies
  - discount only applies when totalItems >= 5
  - Use %.2f in fmt.Printf for dollar amounts

BONUS (optional — uses os.Args like control.go):
  - If the program is run with 4 command-line arguments:
      go run snack_shop.go <chips> <sodas> <candies> <cash>
    use those values instead of hard-coded ones.
  - If wrong number of args, print a usage message and return.

================================================================================
*/

package main

import (
	"fmt"
)

func main() {
	// TODO: Declare your variables here
	// chips, sodas, candies (int)
	// cashGiven (float64)

	// TODO: Set starting values (pick your own, or use the example):
	// 2 chips, 3 sodas, 1 candy, $15.00 cash

	// TODO: Calculate subtotal, total items, discount, final total

	// TODO: Print receipt with fmt.Printf

	// TODO: if/else for change vs still owe

	// TODO: if change == 0, print thank-you message

	fmt.Println("Complete the challenge above! Read the comments at the top of this file.")
}
