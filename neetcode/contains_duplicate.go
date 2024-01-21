package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/algorithm/sort"
	"github.com/liyue201/gostl/ds/slice"
	"github.com/liyue201/gostl/utils/comparator"
)

func test_stl_lib() {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 3)
	s = append(s, 2)
	fmt.Println(s)
	sw := slice.NewSliceWrapper(s)
	fmt.Println(sw)
	sort.Sort[int](sw.Begin(), sw.End(), comparator.IntComparator)
	fmt.Println(sw)
}

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

func containsDuplicate(nums []int, n int) bool {
	numw := slice.NewSliceWrapper[int](nums)
	sort.Sort[int](numw.Begin(), numw.End(), comparator.IntComparator)
	for i := 0; i < n-1; i++ {
		if numw.At(i) == numw.At(i+1) {
			return true
		}
	}

	return false
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	fmt.Println(containsDuplicate(nums, n))
}
