package main

func longestSubarray(nums []int) int {
	n := len(nums)
	i, j := 0, 0

	zerosAvaliable := 1
	countOnes := 0
	res := 0

	for i < n && j < n {
		if nums[j] == 1 {
			countOnes++
			j++
		} else {
			if zerosAvaliable > 0 {
				zerosAvaliable--
				countOnes++
				j++
			} else {
				for i < n && nums[i] == 1 {
					i++
					countOnes--
				}
				zerosAvaliable++
				countOnes--
				i++
			}
		}
		res = max(res, countOnes - 1)
	}

	return res
}
