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

func simple_binary_search(nums []int, n int, target int) int {
	start, end := 0, n-1
	for start <= end {
		mid := start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else {
			if target > nums[mid] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}

	return -1
}

func rotated_array_binary_search_with_direct(nums []int, n, target int) int {
	start, end := 0, n-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == target {
			return mid
		} else {
			/*
				check mid lies in left segment or right => nums[mid] >= nums[start] => in left segment
					what condition target it can be in right segment =: in left segment but to right of mid OR target < nums[start] ==> target > nums[mid] || target < nums[start] => start = mid + 1
					else end = mid - 1
				in right segment
					what condition target can be in left segment := either left of mid but in right segment OR target > nums[end] ==> target < nums[mid] || target > nums[end] ==> end = mid - 1
					else start = mid + 1
			*/
			if nums[mid] >= nums[start] {
				// in left segment
				if target > nums[mid] || target < nums[start] {
					start = mid + 1
				} else {
					end = mid - 1
				}
			} else {
				if target < nums[mid] || target > nums[end] {
					end = mid - 1
				} else {
					start = mid + 1
				}
			}
		}
	}

	return -1
}

func find_pivot(nums []int, n int) int {
	if nums[0] < nums[n-1] { // case of no rotation.
		return n - 1
	}

	start, end := 0, n-1
	for start <= end {
		mid := start + (end-start)/2
		// fmt.Println(nums[start], nums[end], nums[mid])
		/*
			if mid == 0 || mid == n - 1 => it is a pivot.
			if the mid element is in left segment => nums[mid] > nums[n - 1]
				if nums[mid + 1] < nums[mid] => return mid is pivot
				else mid lies in right side => start = mid + 1
			else
				pivot lies in left side => end = mid - 1
		*/

		if mid == 0 || mid == n-1 {
			return mid
		}

		if nums[mid] > nums[n-1] {
			// in left segment
			if nums[mid] > nums[mid+1] {
				return mid
			}
			// else pivot is in right side.
			start = mid + 1
		} else {
			// mid in right segment, we want to reach pivot that is in left segment.
			end = mid - 1
		}
	}

	return -1
}

func rotated_array_binary_search_with_pivot(nums []int, n, target int) int {
	pivot := find_pivot(nums, n) // pivot is the maximum number : [4 5 6 7 0 1 2 3] have pivot as 7
	if pivot == -1 {
		// not rotated sorted array
		return -1
	}
	fmt.Println("pivot : ", pivot)
	if target > nums[n-1] ||
		pivot == n-1 { // case of no rotation.
		return simple_binary_search(nums[0:pivot+1], pivot+1, target)
	} else {
		return simple_binary_search(nums[pivot:], n-pivot, target)
	}
}

func main() {
	n := readInt()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}
	target := readInt()
	// fmt.Println(simple_binary_search(nums, n, target))
	fmt.Println(rotated_array_binary_search_with_pivot(nums, n, target))
	fmt.Println(rotated_array_binary_search_with_direct(nums, n, target))
}
