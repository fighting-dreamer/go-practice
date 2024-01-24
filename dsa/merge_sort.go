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

func merge(nums []int, start, mid, end int) {
	head1 := start
	end1 := mid

	head2 := mid + 1
	end2 := end

	var index, head int
	index = 0
	tmp := make([]int, end-start+1)
	for head1 <= end1 && head2 <= end2 { // important to have '<=' and not simple '!='
		if nums[head1] <= nums[head2] {
			head = head1
			head1++
		} else {
			head = head2
			head2++
		}
		tmp[index] = nums[head]
		index++
	}

	for head1 <= end1 {
		tmp[index] = nums[head1]
		index++
		head1++
	}
	for head2 <= end2 {
		tmp[index] = nums[head2]
		index++
		head2++
	}

	for i := 0; i < len(tmp); i++ {
		nums[start+i] = tmp[i]
	}
}

func merge_sort(nums []int, start, end int) {
	// fmt.Println(start, end)
	if start >= end {
		return
	}

	mid := (start + end) / 2
	merge_sort(nums, start, mid)
	merge_sort(nums, mid+1, end)
	merge(nums, start, mid, end)
}

func mergeSort(nums []int, n int) {
	merge_sort(nums, 0, n-1)
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}
	fmt.Println(nums)
	mergeSort(nums, n)
	fmt.Println(nums)
}
