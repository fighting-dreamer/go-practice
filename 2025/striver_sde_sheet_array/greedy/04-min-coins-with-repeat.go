package main

import (
	"fmt"
	"log"
	"math"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("could not read input number", err)
	}
	return n
}

func readString() string {
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatal("could not read input string", err)
	}
	return s
}

func readDouble() float64 {
	var f float64
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		log.Fatal("could not read input float", err)
	}
	return f
}

func readChar() rune {
	var r rune
	_, err := fmt.Scanf("%c", &r)
	if err != nil {
		log.Fatal("Could not read input char", err)
	}
	return r
}

// Canonical coin system : {1,2,5,10, 20, 50, 100, 500, 1000, 2000}
// non-canonical coin systm : {1,3,4} just an example
// consider : Capacity = 6
// with greedy approach : answer is {4, 1, 1} => 3 coins
// but real right answer is {3,3} => 2 coins (minimum coins)
// greedy approach work for canonical coin system
// we have to use DP for non-canonical system or general.

// mimimum coins, result coins set, remainder if any
// canonical coin system : {1,2,5,10, 20, 50, 100, 500, 1000, 2000}
func caconical_CoinSystem(arr []int, n int, V int) (int, []int, int) {
	res := 0
	resArr := []int{}
	for i := n - 1; i >= 0; i-- {
		for V > arr[i] {
			res += 1
			resArr = append(resArr, arr[i])
			V -= arr[i]
		}
	}

	return res, resArr, 0
}

// mimimum coins, result coins set, remainder if any
// non canonical coint system : {1,3,4}
func non_canonical_CointSystem(coins []int, n int, V int) (int, []int, int) {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+V)
	}

	// when no items => you can't have coins for given V
	for i := 0; i < 1+V; i++ {
		dp[0][i] = math.MaxInt32 - 1
	}

	for i := 1; i < 1+n; i++ {
		dp[i][0] = 0 // if no weight, no coins required => 0 coins.
	}

	// starting for coins from 1 ot more.
	for i := 1; i < 1+n; i++ {
		// starting from weight = 1 unit
		for j := 1; j < 1+V; j++ {
			if coins[i-1] <= j {
				coinsIfTaken := 1 + dp[i][j-coins[i-1]]
				coinsIfNotTaken := dp[i-1][j]
				dp[i][j] = min(coinsIfTaken, coinsIfNotTaken)
			} else {
				dp[i][j] = dp[i-1][j] // couldn't take ith coin
			}
		}
	}

	for i := 0; i < 1+n; i++ {
		fmt.Printf("%d | ", i)
		for j := 0; j < 1+V; j++ {
			fmt.Printf("%3d ", dp[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-----------------------------")
	return dp[n][V], []int{}, 0
}

func main() {
	n := readInt()
	V := readInt()
	coins := make([]int, n)
	for i := 0; i < n; i++ {
		coins[i] = readInt()
	}

	fmt.Println(non_canonical_CointSystem(coins, n, V))
}
