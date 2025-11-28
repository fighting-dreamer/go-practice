package main

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	res := -1
	for l < r {
		area := (r - l) * min(height[r], height[l])
		res = max(res, area)

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
