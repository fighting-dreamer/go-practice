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

/*
	// Pseudo Algo :
	we have [...nums], target ===> array of arrays

	for each num_i if it is less than target, we can probably make a solution array.

	combination_sum_helper(nums, target, solution_till_now) :
	 if target is zero => solution till now is a valid candidate solution.

	 result of type array-of-array
	 for each num_i
		if num_i <= target :
			check if combination_sum(nums, target - num_i, combine(solution-till_now ,num_i)) is resulting in a solution
			if check is true :
				add all candidate solutions found to the result


	if length of result > 0 {
		we found some valid candidate solutions
		return the result and the status that we found valid solutions.
	}

	return empty result.
*/

func combination_sum_helper(nums []int, n, t int, curr []int) ([][]int, bool) {
	if t == 0 {
		return [][]int{curr}, true
	}
	res := [][]int{}
	for i := 0; i < n; i++ {
		if t < nums[i] {
			continue
		}
		currRes, flag := combination_sum_helper(nums, n, t-nums[i], append(curr, nums[i]))
		if flag == true {
			res = append(res, currRes...)
		}
	}
	// fmt.Println(curr, " : ", res)
	if len(res) > 0 {
		return res, true
	}

	return [][]int{}, false // empty result
}

func combination_sum(nums []int, n, t int) [][]int {
	res := [][]int{}
	if t == 0 {
		return res
	}
	curr := []int{}
	for i := 0; i < n; i++ {
		if t >= nums[i] {
			currRes, flag := combination_sum_helper(nums, n, t-nums[i], append(curr, nums[i]))
			if flag {
				res = append(res, currRes...)
			}
		}
	}

	return res
}

// 4 7 2 3 6 7 ==> [[2,2,3], [7]]
// 3 8 2 3 5 => [[2,2,2,2], [2,3,3], [3,5]]
func main() {
	n, target := readInt(), readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(combination_sum(nums, n, target))
}
