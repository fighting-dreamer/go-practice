package main

import (
	"fmt"
	"log"
	"math"
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

// ------------------------------- 01 knapsack based : START ----------

func _01_knapsack_code(wtArr []int, priceArr []int, n int, cap int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+cap)
	}

	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0 // there are i number of items but knapsack capacity is zero
	}
	for i := 0; i < 1+cap; i++ {
		dp[0][i] = 0 // there are no items, indifferent of knapsack capacity
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+cap; j++ {
			wt_of_i_th_item := wtArr[i-1]
			if wt_of_i_th_item <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-wt_of_i_th_item]+priceArr[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][cap]
}

func _01_knapsack() {
	n := readInt()
	cap := readInt()

	wtArr := make([]int, n)
	priceArr := make([]int, n)
	for i := 0; i < n; i++ {
		wtArr[i] = readInt()
	}
	for i := 0; i < n; i++ {
		priceArr[i] = readInt()
	}

	fmt.Printf("_01_knapsack_code(wtArr, priceArr, n, cap): %v\n", _01_knapsack_code(wtArr, priceArr, n, cap))
}

func subset_sum_code(arr []int, n int, s int) bool {
	dp := make([][]bool, 1+n) // dp[i][j] tells is target sum j is possible with first i items
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]bool, 1+s)
	}
	for i := 0; i < 1+n; i++ {
		dp[i][0] = true // i items and sum required is zero is possible
	}
	for i := 1; i < 1+s; i++ {
		dp[0][i] = false // no items and non-zero postive sum required => not possible
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][s]
}

func subset_problem() {
	n := readInt()
	s := readInt()
	arr := make([]int, n)

	subset_sum_code(arr, n, s)
}

func equal_sum_partition_code(arr []int, n int, s int) bool {
	if s&1 == 1 {
		return false // the sum is odd, cant be divided equally
	}
	s = s / 2

	return subset_sum_code(arr, n, s)
}

func count_subsets_with_given_sum_code(arr []int, n int, s int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}
	for i := 0; i < 1+n; i++ {
		dp[i][0] = 1 // i items, and sum required is zero, ways to do that is 1
	}
	for i := 1; i < 1+s; i++ {
		dp[0][i] = 0 // no items, sum required is non-zero, ways = 0
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if arr[i-1] <= j {
				// we can atain sum j with ith element probably
				dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][s]
}

func count_subsets_with_given_difference(arr []int, n int, diff int) int {
	// s1 - s2 = diff
	// s1 + s2 = total
	// s1 = (total + diff) / 2
	total := 0
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	targetSum := (total + diff) / 2
	return count_subsets_with_given_sum_code(arr, n, targetSum)
}

func count_subsets_with_given_sum() {
	n := readInt()
	s := readInt()
	arr := make([]int, n)

	count_subsets_with_given_sum_code(arr, n, s)
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func count_subsets_with_minimum_difference_code(arr []int, n int) int {
	// s1 - s2 = minimum
	// s1 + s2 = total
	// targetSum (or s1) = (total + minimum) / 2

	totalS := 0
	for i := 0; i < n; i++ {
		totalS += arr[i]
	}

	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+totalS)
	}

	for i := 0; i < 1+n; i++ {
		for j := 0; j < 1+totalS; j++ {
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			if i == 0 {
				dp[i][j] = 0
			}
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	minDiff := math.MaxInt
	for i := 0; i <= totalS/2; i++ {
		if dp[n][i] != 0 {
			minDiff = min(minDiff, Abs(Abs(totalS-i)-i))
		}
	}

	fmt.Println("Min Diff : ", minDiff)

	return minDiff
}

func count_subsets_with_minimum_difference() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	fmt.Printf("count_subsets_with_minimum_difference_code(arr, n): %v\n", count_subsets_with_minimum_difference_code(arr, n))
}

func equal_sum_partition() {
	n := readInt()
	arr := make([]int, n)
	s := 0
	for i := 0; i < n; i++ {
		arr[i] = readInt()
		s += arr[i]
	}

	equal_sum_partition_code(arr, n, s)
}

// ------------------------------- 01 knapsack based : END ----------

// ------------------------------- unbounded knapsack based : START ----------

func rod_cutting_code(larr []int, parr []int, n int, rodLen int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+rodLen)
	}
	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0 // value of rod len = 0
	}
	for i := 1; i < 1+rodLen; i++ {
		dp[0][i] = 0 // no item,
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+rodLen; j++ {
			if larr[i-1] <= j {
				// length of rod (ie. j) is greater than i-th section : larr[i - 1], the it here is not same as ith that;s why i - 1
				// we an also re-choose the same rod segment so we are using dp[i][j - larr[i - 1]], notice the i in the first seelction
				dp[i][j] = max(dp[i-1][j], parr[i-1]+dp[i][j-larr[i-1]])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][rodLen]
}

func rod_cutting() {
	n := readInt()
	lArr := make([]int, n)
	priceArr := make([]int, n)
	rodLen := readInt()

	fmt.Printf("rod_cutting_code(lArr, priceArr, n, rodLen): %v\n", rod_cutting_code(lArr, priceArr, n, rodLen))
}

func coin_change_1_code(arr []int, n int, s int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}

	// dp[0][0] = 0 : imp
	for i := 0; i < 1+n; i++ {
		dp[i][0] = 0 // sum reqd is 0, and items avaliable is i, minimum coins => 0
	}

	for i := 1; i < 1+s; i++ {
		// you can use INF = 1e9 in place of "math.MaxInt - 1" to simplify.
		dp[0][i] = math.MaxInt - 1 // non zero sum reqd fom 0 coins
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if arr[i-1] <= j {
				dp[i][j] = min(dp[i-1][j], dp[i][j-arr[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][s]
}

func coin_change_1_minimum_coin_change() {
	// give minimum coins for given sum
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	s := readInt()
	fmt.Printf("coin_change_1_code(arr, n, s): %v\n", coin_change_1_code(arr, n, s))
}

func coin_change_2_maximum_ways(arr []int, n int, s int) int {
	dp := make([][]int, 1+n)
	for i := 0; i < 1+n; i++ {
		dp[i] = make([]int, 1+s)
	}

	for i := 0; i < 1+n; i++ {
		dp[i][0] = 1 // i items, sum requied is 0 => 1 way
	}

	for i := 1; i < 1+s; i++ {
		dp[0][i] = 0 // no item for non-zero sum => no way => 0
	}

	for i := 1; i < 1+n; i++ {
		for j := 1; j < 1+s; j++ {
			if arr[i-1] <= j {
				// we can find a coin under sum j, using ith coin
				dp[i][j] = dp[i-1][j] + dp[i][j-arr[i-1]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[n][s]
}

func coin_change_2_maximum_ways() {
	// given n coin denominations, maximum ways to get the target sum s
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	s := readInt()
	coin_change_2_maximum_ways_code(arr, n, s)
}

// ------------------------------- unbounded knapsack based : END ----------

// ------------------------------- lcs based : START ----------

func lcs_code(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, 1+l1) // longest length of common sub-sequence
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}
	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l1][l2]
}

func lcs() {
	// longest common sub-sequence
	s1 := readString()
	s2 := readString()

	fmt.Printf("lcs_code(s1, s2): %v\n", lcs_code(s1, s2))
}

func longest_common_substring_code(s1 string, s2 string) int {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, 1+l1) // suffix ending with i,j does they match, if so, whats the length of substring thats common to both.
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}
	res := 0
	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				res = max(res, dp[i][j])
			} else {
				dp[i][j] = 0 // we are cehcking if substring till i is same as substring j
			}
		}
	}

	return res
}

func longest_common_substring() {
	s1 := readString()
	s2 := readString()

	fmt.Printf("longest_common_substring(s1, s2): %v\n", longest_common_substring_code(s1, s2))
}

func printing_longest_common_subsequence(s1 string, s2 string) string {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, 1+l1)
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}

	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	// building the lcs string
	res := make([]byte, dp[l1][l2])
	index := dp[l1][l2] - 1
	i, j := l1, l2

	for i >= 0 && j >= 0 {
		if dp[i][j] == dp[i-1][j-1]+1 {
			res[index] = s1[i-1]
			i--
			j--
			index--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}

	return string(res)
}

func printing_lcs() {
	s1 := readString()
	s2 := readString()

	fmt.Printf("printing_longest_common_subsequence(s1, s2): %v\n", printing_longest_common_subsequence(s1, s2))
}

func shortest_common_subsequence() int {
	s1 := readString()
	s2 := readString()
	return len(s1) + len(s2) - lcs_code(s1, s2)
}

func print_shortest_common_supersequence_code(s1 string, s2 string) string {
	l1, l2 := len(s1), len(s2)
	dp := make([][]int, 1+l1)
	for i := 0; i < 1+l1; i++ {
		dp[i] = make([]int, 1+l2)
	}
	for i := 0; i < 1+l1; i++ {
		for j := 0; j < 1+l2; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	lcs := dp[l1][l2]

	res := make([]byte, l1+l2-lcs)
	index := l1 + l2 - lcs - 1
	i, j := l1, l2
	for i > 0 && j > 0 {
		if dp[i][j] == dp[i-1][j-1]+1 {
			res[index] = s1[i-1]
			i--
			j--
		} else {
			if dp[i-1][j] > dp[i][j-1] {
				res[index] = s1[i-1]
				i--
			} else {
				res[index] = s2[j-1]
				j--
			}
		}
		index--
	}

	for i >= 0 {
		res[index] = s1[i-1]
		i--
		index--
	}

	for j >= 0 {
		res[index] = s2[j-1]
		j--
		index--
	}

	return string(res)
}

func print_shortes_common_supersequence() {
	s1 := readString()
	s2 := readString()

	print_shortest_common_supersequence_code(s1, s2)
}

func minimum_insertions_to_make_string_palindrome_code(s string) int {
	s1 := s
	reverseS := make([]rune, len(s))
	copy(reverseS, []rune(s))
	slices.Reverse(reverseS)

	res := len(s) - lcs_code(s1, string(reverseS))
	return res
}

func minimum_insertions_to_make_string_palindrome() {
	s := readString()
	fmt.Printf("minimum_insertions_to_make_string_palindrome_code(s): %v\n", minimum_insertions_to_make_string_palindrome_code(s))
}

// ------------------------------- lcs based : END ----------

// ------------------------------- matrix multiplication chain based : START ----------

func mcm_code(arr []int, n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	// for length of section to be 2 upto n
	for l := 2; l <= n; l++ {
		// starting section from index i = 0 upto n - len
		for i := 0; i < n-l; i++ {
			dp[i][i+l] = math.MaxInt
			for k := i + 1; k < i+l; k++ {
				cost := dp[i][k] + dp[k][i+l] + arr[i]*arr[k]*arr[i+l]
				dp[i][i+l] = min(dp[i][i+l], cost)
			}
		}
	}

	return dp[0][n-1]
}

func mcm() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Printf("mcm_code(arr, n): %v\n", mcm_code(arr, n))
}

func egg_dropping_memoization_code(e int, f int, m map[[2]int]int) int {
	if e == 0 {
		m[[2]int{e, f}] = math.MaxInt
		return m[[2]int{e, f}]
	}
	if e == 1 || f == 1 {
		m[[2]int{e, f}] = f
		return f
	}
	if v, ok := m[[2]int{e, f}]; ok {
		return v
	}

	res := math.MaxInt
	for i := 1; i <= f; i++ {
		temp := 1 + max(
			egg_dropping_memoization_code(e-1, i-1, m),
			egg_dropping_memoization_code(e, f-i, m))
		res = min(res, temp)
	}
	m[[2]int{e, f}] = res
	return m[[2]int{e, f}]
}

func egg_dropping_code(no_of_eggs int, floors int) int {
	if no_of_eggs == 0 {
		return -1 // infinite
	}
	if no_of_eggs == 1 {
		return floors
	}
	if floors <= 1 {
		return floors
	}

	// we choose kth floor to be probably giving best answer
	// at kth floor, two scenarios : either egg break or does not break
	// if break : the problem is to be solved with e-1 eggs and for k - 1 floors
	// if not break : we have to solve problem for e eggs, f - k floors

	m := map[[2]int]int{}

	return egg_dropping_memoization_code(no_of_eggs, floors, m)
}

func egg_dropping() {
	// minimize the number of attempts
	n := readInt()
	f := readInt()

	fmt.Printf("egg_dropping_code(n, f): %v\n", egg_dropping_code(n, f))
}

// ------------------------------- matrix multiplication chain based : END ----------

func main() {
	// _01_knapsack()
	// subset_problem()
	// equal_sum_partition()
	// count_subsets_with_given_sum()
	// count_subsets_with_minimum_difference()

	// rod_cutting()
	// coin_change_1_minimum_coin_change()
	// coin_change_2_maximum_ways()

	// lcs()
	// longest_common_substring() // best
	// shortest_common_supersequence() = n + m - lcs_value
	// printing_shortest_common_supersequence() // best
	// minimum_insertions_to_make_string_palindrome()

	// mcm()
	// palindrome_partitioning() // not done, good
	// egg_dropping() // best

	// stock_buy_sell_

}
