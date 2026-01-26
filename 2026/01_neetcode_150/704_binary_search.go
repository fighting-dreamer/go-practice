package main

import "slices"

func binary_search(arr []int, n int, k int) int {
	start := 0
	end := n - 1

	for start < end {
		mid := (start + end) / 2

		if arr[mid] == k {
			return mid
		} else {
			if arr[mid] < k {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func lower_bound(arr []int, n int, k int) int {
	start := 0
	end := n

	for start < end {
		mid := (start + end) / 2

		if arr[mid] < k {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func upper_bound(arr []int, n int, k int) int {
	start := 0
	end := n
	for start < end {
		mid := start + (end-start)/2

		if arr[mid] <= k {
			start = mid + 1
		} else {
			end = mid
		}
	}

	return start
}

func main() {
	n := readInt()
	target := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	slices.Sort(arr)
	binary_search(arr, n, target)
	lower_bound(arr, n, target)
	upper_bound(arr, n, target)
}
