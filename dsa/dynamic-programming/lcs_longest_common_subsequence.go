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

func lcs_recursive_helper(s1, s2 string, index1, index2 int) int {
	if index1 >= len(s1) || index2 >= len(s2) {
		return 0
	}

	if s1[index1] == s2[index2] {
		return 1 + lcs_recursive_helper(s1, s2, index1+1, index2+1)
	}
	return max(lcs_recursive_helper(s1, s2, index1+1, index2), lcs_recursive_helper(s1, s2, index1, index2+1))
}

func lcs_recursive(s1, s2 string) int {
	// either the first shar will be same for two strings Or it will not be
	// if it is not, the next possible longest string can be found by eliminating that char from string1 OR string 2
	index1 := 0
	index2 := 0

	return lcs_recursive_helper(s1, s2, index1, index2)
}

func lcs_memoization_helper(s1, s2 string, index1, index2 int, dp [][]int) int {
	if index1 >= len(s1) || index2 >= len(s2) {
		return 0
	}

	if dp[index1][index2] != 0 {
		return dp[index1][index2]
	}

	if s1[index1] == s2[index2] {
		dp[index1][index2] = 1 + lcs_memoization_helper(s1, s2, index1+1, index2+1, dp)
		return dp[index1][index2]
	}
	r1 := lcs_memoization_helper(s1, s2, index1+1, index2, dp)
	r2 := lcs_memoization_helper(s1, s2, index1, index2+1, dp)
	dp[index1][index2] = max(r1, r2)

	return dp[index1][index2]
}

func lcs_memoization(s1, s2 string) int {
	//from recursivefunction we understand that the lcs_recursive is computing for lcs_value for pair of (index1, index2)
	// we create a matrix to keep track of such value if already computed.
	// example of scenario : index pair say (5,7) --> 5,6 and 6,5 --> 4,5 AND 5,5 | 6,4 --> the 4,5 could be reused but the 5,5 --> 4,5 and it is computed again.

	dp := make([][]int, len(s1))
	for i := 0; i < len(s1); i++ {
		dp[i] = make([]int, len(s2))
	}

	lcs_memoization_helper(s1, s2, 0, 0, dp)
	// fmt.Println("------------------------------------------")
	// for i := 0; i < len(s1); i++ {
	// 	for j := 0; j < len(s2); j++ {
	// 		fmt.Print(dp[i][j], " ")
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("------------------------------------------")
	return dp[0][0] // here the value of substring 0-xi and 0-yi is dp[len(s1) - xi][len(s2) - yi] and for whole s1 and s2 lengths it is dp[len(s1) - len(s1)][len(s2) - len(s2)]
}

func lcs_bottom_up(s1, s2 string) int {
	l1 := len(s1)
	l2 := len(s2)
	dp := make([][]int, 1+l1)
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}

	for i := 0; i < 1+l1; i++ {
		dp[i][0] = 0
	}
	for i := 0; i < 1+l2; i++ {
		dp[0][i] = 0
	}

	for i := 1; i < 1+l1; i++ {
		for j := 1; j < 1+l2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}

	return dp[l1][l2]
}

func lcs(s1, s2 string) int {
	recursive := lcs_recursive(s1, s2)
	memoization := lcs_memoization(s1, s2)
	bottomUp := lcs_bottom_up(s1, s2)

	fmt.Println("recusrive : ", recursive)
	fmt.Println("memoization : ", memoization)
	fmt.Println("bottom up : ", bottomUp)

	return recursive
}

func main() {
	s1, s2 := readString(), readString()

	fmt.Println(lcs(s1, s2))

}
