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

func find_max_contigous_subarray_product(nums []int, n int) int {
	// nums can contain zero and negtives
	// no zero cases :
	// if nums contain only postive => just product of all nums is the answer
	// if the nums contains negtives , we know product of two negtive make postive, cases :
	// .... // if negtive number cone even number of times : product of whole array is the answer
	// .... // if negtive number come odd number of times : we have to leave on of the negtive number.
	// .... // .. essentially, either a prefix or a suffix of the negtive number that is being left is the answer.

	// cases with zero :
	// if w are having zeros, the subarray around zero or between two zero need to be considered only for solution and for those the above cases can be applied.
	// example : [-5 9 -5 -1 0 4 -7 1 5 9 -4 0 7 -8 9 4]
	// for these, when we encounter the zero ,we update the product to 1 again.
	res := 0
	prefixProduct := 1
	suffixProduct := 1

	for i := 0; i < n; i++ {
		if prefixProduct == 0 {
			prefixProduct = 1
		}
		if suffixProduct == 0 {
			suffixProduct = 1
		}

		prefixProduct = prefixProduct * nums[i]
		suffixProduct = suffixProduct * nums[n-i-1]

		res = max(res, prefixProduct, suffixProduct)
	}

	return res
}

func main() {
	n := readInt()
	nums := make([]int, n)

	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(find_max_contigous_subarray_product(nums, n))
}
