package main

import (
	"fmt"
	"log"
	"slices"
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

func count_subsets_with_given_sum(set []int, n, s int) int {
	fmt.Println("Before : ", set)
	// sorting in decending order
	slices.SortFunc(set, func(i, j int) int {
		// i and j are i/jth element itself, not indexes
		if i < j {
			return 1
		}
		return -1
	})

	fmt.Println("After : ", set)

	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}
	dp[0][0] = 1
	for i := 1; i < 1+n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < 1+s; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if set[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-set[i-1]]
			}
		}
	}
	fmt.Println("----------------------")
	for i := 0; i < 1+n; i++ {
		for j := 0; j < 1+s; j++ {
			fmt.Print(dp[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("----------------------")
	return dp[n][s]
}

// input : 4 6 3 3 3 3 ==> 6
// input : 4 3 0 1 1 1 ==> 2
// form above input, you see that the element of value 0 is making the output to be 1, whereas it should be 2, so to get right answer, sort the set in descending order

func main() {
	n, s := readInt(), readInt()
	set := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = readInt()
	}

	fmt.Println(count_subsets_with_given_sum(set, n, s))
}
