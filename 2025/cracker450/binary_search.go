package main

import (
	"fmt"
	"log"
	"sort"
)

func readInt() int {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal("Couldn't read number")
	}

	return n
}

func readDouble() float64 {
	var n float64
	_, err := fmt.Scanf("%f", &n)
	if err != nil {
		log.Fatal("couldn't read number")
	}
	return n
}

func binary_serach(arr []int, n int, k int) int {
	start, end := 0, n-1
	var mid int
	for start <= end {
		mid = start + (end-start)/2

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

func isMinInRotated(arr []int, it int, n int) bool {
	prev := (it + n - 1) % n
	next := (it + 1) % n
	// if minimum have duplicates, it should come after curr position of its start
	if arr[next] >= arr[it] && arr[prev] >= arr[it] {
		return true
	}
	return false
}

func binary_search_rotated_array_minimum_elem(arr []int, n int) int {
	// we need a way to say the curr elem is a minimum elem in this rorated array
	// we need a way to judge that we should go for right section or left section
	start := 0
	end := n - 1
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if isMinInRotated(arr, mid, n) {
			return mid
		} else {
			// we have to chooes the next section that is un-sorted

			// this implies we are in rotated section and the mid lies in that section,
			// we are suppose to go to other section then.
			if arr[mid] > arr[end] {
				start = mid + 1
			} else {
				end = mid
			}
		}
	}
	return -1
}

func find_elem_in_rotated_array(arr []int, n int, k int) int {
	pivot := binary_search_rotated_array_minimum_elem(arr, n)
	if pivot == -1 {
		return -1
	}
	fmt.Println(pivot)
	print_array(arr[:pivot], pivot)
	print_array(arr[pivot:], n-pivot)
	left := binary_serach(arr[:pivot], pivot, k)
	right := binary_serach(arr[pivot:], n-pivot, k)

	if left != -1 {
		return left
	}
	return pivot + right
}

func find_lower_bound(arr []int, n int, k int) int {
	start := 0
	end := n
	for start < end {
		mid := (start + end) / 2

		if arr[mid] >= k { // mid element is of right side of k or equal
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

func find_upper_bound(arr []int, n int, k int) int {
	start := 0
	end := n

	for start < end {
		mid := (start + end) / 2

		if arr[mid] > k {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}

// -------------------------- helper---------------------

func sort_array(arr []int, n int) {
	sort.Ints(arr)
}

func print_array(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

// ------------------------------------------------- Testing ---------------------

func test_binary_search(arr []int, n int) {
	sort_array(arr, n)
	print_array(arr, n)

	k := readInt()
	fmt.Println("k : ", k, " index : ", binary_serach(arr, n, k))
}

func test_rotated_binary_serach(arr []int, n int) {
	fmt.Println(" index : ", binary_search_rotated_array_minimum_elem(arr, n))
}

func test_binary_search_in_rotated_array(arr []int, n int) {

	k := readInt()
	fmt.Println("k : ", k, " index : ", find_elem_in_rotated_array(arr, n, k))
}

func test_lower_bound(arr []int, n int) {
	k := readInt()
	fmt.Println("k : ", k, " lower index : ", find_lower_bound(arr, n, k))

}

func test_upper_bound(arr []int, n int) {
	k := readInt()
	fmt.Println("k : ", k, " upper index : ", find_upper_bound(arr, n, k))

}

func main() {
	n := readInt()

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}
	print_array(arr, n)

	// test_binary_search(arr, n)

	// test_rotated_binary_serach(arr, n)

	// test_binary_search_in_rotated_array(arr, n)

	test_lower_bound(arr, n)
	test_upper_bound(arr, n)
}
