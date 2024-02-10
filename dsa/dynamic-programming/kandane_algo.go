package main

import "fmt"

func find_max_sum_contigous_subarray_kandane_algo(nums []int, n int) int {
	// we will return the max sum
	// for the sub-array, we must understand what leads us to not select the section coming before the start of subarray and after the end of that sub-array.
	// example : [-5, 4 -9 .... 6 -3 4 .... 4 5 1] say the answer is [6 -3 4]
	// we could have included the segment before the 6, we did not => it must be having some negtive element just before 6 coz that negtive will lead to decrease of value of sum of the subarray.
	// the segment must be having a total sum of segment as negtive
	// same with segment coming after 4, it must be negtive and the sum of that segment is negtive in total.
	// what we learn is that if we ever get to a sum of segment as negtive, it can be discarded or left to start a new segment.
	// at each point, we want to keep track of hte max sum segment or the max sum
	// say we keep on summing the elements, when the sum till ith element is negtive, we can discard to start the new segment
	// at each element, we can keep track of the max sum, as with first rule or property, we know we are discsrding any segment once it becomes negtive.
	// here we are saying a empty sub-array is also  valid answer than a negtive one.
	// if we find the end result to be empty, we can find the smallest element and use that to create a subarray of size one with that element.
	curr_segment_sum := 0
	max_sum_till_now := 0
	for i := 0; i < n; i++ {
		curr_segment_sum = curr_segment_sum + nums[i]
		if curr_segment_sum > max_sum_till_now {
			max_sum_till_now = curr_segment_sum
		}
		if curr_segment_sum < 0 {
			curr_segment_sum = 0 // start fresh
		}
	}

	return max_sum_till_now
}

func main() {
	// we want to find the sub-array with max sum
	n := readInt()
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(find_max_sum_contigous_subarray_kandane_algo(nums, n))
}
