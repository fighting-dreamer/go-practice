package main

import (
	"fmt"
	"log"
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

func _01_knapsack(wtArr []int, priceArr []int, n int, k int) int {
	dp := make([][]int, 1+n) // dp store the res max value if there were n items and knackack capacity was k
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+k)
	}
	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0 // i items but knacksack capacity is zero
	}
	for i := 0; i < 1+k; i++ {
		dp[0][i] = 0 // no item but i`th capicity of knacksack
	}

	for i := 1; i <= n; i++ {
		for j := 0; j < 1+k; j++ {
			// we choose wheather w ecan take ith item or not
			if wtArr[i-1] <= j {
				notPickingIthItem := dp[i-1][j]
				pickingIthItem := dp[i-1][j-wtArr[i-1]] + priceArr[i-1]
				dp[i][j] = max(notPickingIthItem, pickingIthItem)
			} else {
				// weight of item is already greater than the knapsack capacity.
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return 0
}

func main() {
	n := readInt()
	k := readInt()

	wtArr := make([]int, n)
	priceArr := make([]int, n)
	for i := 0; i < n; i++ {
		wtArr[i] = readInt()
	}
	for i := 0; i < n; i++ {
		priceArr[i] = readInt()
	}

	fmt.Printf("_01_knapsack(wtArr, priceArr, n, k): %v\n", _01_knapsack(wtArr, priceArr, n, k))

}
