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

func find_pivot_in_rotated_sorted_array(nums []int, n int) int {
	start := 0
	end := n - 1
	// pivot is the smalles telement ie the element before it is either not exist or the biggest
	// example : 1 2 3 4 5 6 7 ===> 1 is pivot
	// example : 5 6 7 1 2 3 4 ===> 1 is hte pivot
	// example : 2 3 4 5 6 7 1 ===> 1 is the pivot

	// case : no duplicate elements :
	// pivot is always smaller or equal to last element in nums
	// if equal to pivot => element before it is either not exist or bigger
	// if mid is in first sorted section(mid is greater than last element of nums) => pivot lies in mid + 1 to end
	// if mid is in second sorted section(mid is smaller than last element of nums) => pivot lies in start to mid - 1

	// case : duplicate elements :
	// example : 4 5 6 7 1 2 3 4 : pivot is having prev element as bigger OR non-existent
	// example : 1 1 2 3 4 5 : pivot is having prev element as bigger OR non-existent
	// example : 1 2 3 4 5 1 : pivot is having prev element as bigger OR non-existent
	// example : 1 2 3 4 5 1 1 : pivot is having prev element as bigger OR non-existent
	// example : 4 4 4 5 6 7 8 1 2 3 4 4 4 : pivot is having prev element as bigger OR non-existent
	// pivot is always smaller or equal to last element in nums
	// if equal to pivot => element before it is either not exist or bigger
	// if mid is in first sorted section(mid is greater than or equal last element of nums) => pivot lies in mid + 1 to end
	// if mid is in second sorted section(mid is smaller than or equal last element of nums) => pivot lies in start to mid - 1
	// nums[start] < nums[end] ==> we are in sorted section, we should have nums[start] > nums[end] otherwise if nums[start] == nums[end] then nums[mid] is > nums[end] => start = mid and else end = mid

	for start <= end {
		mid := (start + end) / 2
		prev := (mid + n - 1) % n
		if mid == 0 || nums[prev] > nums[mid] {
			return mid
		} else {
			if nums[mid] > nums[n-1] {
				start = mid + 1
			} else if nums[mid] < nums[n-1] {
				end = mid - 1
			} else {
				// nums[start] must also be equal to nums[mid]
				// but the nums[end] could be either in first section or second.
				// nums[mid] is equal to nums[n - 1]
				// end'th element if > mid'th element => end'th element is in first half ==>
				// else end'th element is is in the second half.
				//
				if nums[mid] > nums[end] {
					start = mid
				} else {
					end = mid
				}
			}
		}
	}

	return -1
}

func main() {
	testcases := []struct {
		nums     []int
		expected int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			0,
		},
		{
			[]int{5, 6, 7, 1, 2, 3, 4},
			3,
		},
		{
			[]int{2, 3, 4, 5, 6, 7, 1},
			6,
		},
		{
			[]int{4, 5, 6, 7, 1, 2, 3, 4},
			4,
		},
		{
			[]int{1, 1, 2, 3, 4, 5},
			0,
		},
		{
			[]int{1, 2, 3, 4, 5, 1},
			5,
		},
		{
			[]int{1, 2, 3, 4, 5, 1, 1},
			5,
		},
		{
			[]int{4, 4, 4, 5, 6, 7, 8, 1, 2, 3, 4, 4, 4},
			7,
		},
		{
			[]int{4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 1, 2, 3, 4, 4, 4, 4},
			17,
		},
		{
			[]int{4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 7, 8, 1, 2, 3, 4, 4, 4, 4},
			16,
		},
		{
			[]int{4, 4, 4},
			0,
		},
	}
	n := len(testcases)
	for i := 0; i < n; i++ {
		nums := testcases[i].nums
		fmt.Println(i+1, " : ", nums)
		expected := testcases[i].expected
		got := find_pivot_in_rotated_sorted_array(nums, len(nums))
		fmt.Println("expected : ", expected, "    got : ", got)
	}

}
