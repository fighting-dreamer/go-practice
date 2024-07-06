package main

import "slices"

func max_number_k_sum(nums []int, k int) int {
	slices.Sort(nums)
	n := len(nums)
	l, r := 0, n-1
	res := 0
	for l < r {
		if nums[l]+nums[r] == k {
			l++
			r--
			res++
		} else {
			if nums[l]+nums[r] > k {
				r--
			} else {
				l++
			}
		}
	}

	return res
}
