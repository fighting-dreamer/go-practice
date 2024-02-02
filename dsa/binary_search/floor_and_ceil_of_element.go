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

func floor_of_element(nums []int, n, t int) int {
	// floor is largest element available that is element smaller than t
	// find the element smaller than t(say res) and if next element is also smaller but greater than res => keep looking to right
	fmt.Println("searching for floor of ", t)
	start := 0
	end := n - 1
	var res int
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] < t {
			res = nums[mid]
			// if there is a valid next element and it is less than t => keep searching in right => start = mid + 1 and continue
			//else => we have found the floor.
			if mid+1 != n && nums[mid+1] < t {
				start = mid + 1
				continue
			}
			return res
		} else {
			// if equal or greater, we want to search in left of the array.
			end = mid - 1
		}
	}

	return -1
}

func ceil_of_element(nums []int, n, t int) int {
	// ceil => smallest of all the larger elements thatn t
	// find a smaller element that t say res, if there exist an element to just left that is larger than t => search in left continue
	start := 0
	end := n - 1
	res := nums[0]

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] > t {
			res = nums[mid]
			if mid != 0 && nums[mid-1] > t {
				end = mid - 1
				continue
			}
			return res
		} else {
			start = mid + 1
		}
	}
	return -1
}

func main() {
	n, t := readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(floor_of_element(nums, n, t))
	fmt.Println(ceil_of_element(nums, n, t))
}
