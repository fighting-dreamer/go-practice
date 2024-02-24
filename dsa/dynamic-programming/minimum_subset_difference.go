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

func minimum_subset_difference(set []int, n int) int {
	if n == 0 || n == 1 {
		return -1
	}

	sum := func(arr []int) int {
		res := 0
		for i := 0; i < len(arr); i++ {
			res += arr[i]
		}

		return res
	}(set)
	// say sub-set difference is between s1 and s2
	//  for us, s1 ++ s2 is the original set
	// say s1's sum is s_1, for s2's sum is s_2
	// s_1 - s_2 = min, s_1 + s_2 = sum
	// min that we want is (sum - 2*s_1) => we try to find the sub set possible of size half of total and see if for n items of set, any of it is possible to create and and soon as we find it closer to sum/2 we can use that as answer asn it minimizes the (sum - 2*s_1)

	halfSum := sum/2 + (sum % 2)
	dp := make([][]bool, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = maeke([]bool, 1+halfSum)
	}
	dp[0][0] = true
	for i := 1; i < 1+n; i++ {
		dp[i][0] = true
	}
	for i := 1; i < 1+halfSum; i++ {
		dp[0][i] = false
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+halfSum; j++ {
			if set[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-set[i-1]]
			}
		}
	}

	for i := halfSum; i >= 0; i-- {
		if dp[n][i] == true {
			// min sum := (sum - i) - (i)
			// fmt.Printf("dp[%d][%d] : %d\n",n, i, i)
			return int(math.Abs(float64(sum - 2*i)))
		}
	}
	return -1 // not possible.
}

func main() {
	n := readInt()
	set := make([]int, n)

	for i := 0; i < n; i++ {
		set[i] = readInt()
	}

	fmt.Println(minimum_subset_difference(set, n))
}
