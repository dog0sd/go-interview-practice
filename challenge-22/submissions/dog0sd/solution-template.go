package main

import (
	"fmt"
	"math"
)

func main() {
	// Standard U.S. coin denominations in cents
	denominations := []int{1, 5, 10, 25, 50}

	// Test amounts
	amounts := []int{87, 42, 99, 33, 7}

	for _, amount := range amounts {
		// Find minimum number of coins
		
		minCoins := MinCoins(amount, denominations)

		// Find coin combination
		coinCombo := CoinCombination(amount, denominations)

		// Print results
		fmt.Printf("Amount: %d cents\n", amount)
		fmt.Printf("Minimum coins needed: %d\n", minCoins)
		fmt.Printf("Coin combination: %v\n", coinCombo)
		fmt.Println("---------------------------")
	}

}

// MinCoins returns the minimum number of coins needed to make the given amount.
// If the amount cannot be made with the given denominations, return -1.
func MinCoins(amount int, denominations []int) int {
	counter := 0
	if amount == 0 {
		return 0
	}
	if len(denominations) == 0 || amount < denominations[0]{
		return -1
	}
	amount_left := amount
	for i := range denominations {
	    deno := denominations[len(denominations) - 1 - i]
	    count := int(math.Floor(float64(amount_left) / float64(deno)))
	    counter += count
	    amount_left -= count * deno
	} 
	return counter
}

// CoinCombination returns a map with the specific combination of coins that gives
// the minimum number. The keys are coin denominations and values are the number of
// coins used for each denomination.
// If the amount cannot be made with the given denominations, return an empty map.
func CoinCombination(amount int, denominations []int) map[int]int {
	var result = map[int]int{}
	if len(denominations) == 0 || amount < denominations[0] {
		return result
	}
	amount_left := amount
	for i := range denominations {
	    deno := denominations[len(denominations) - 1 - i]
	    count := int(math.Floor(float64(amount_left) / float64(deno)))
	    result[deno] = count
	    amount_left -= count * deno
	}
	for k, v := range result {
		if v == 0 {
			delete(result, k)
		}
	}
	
	return result
}
