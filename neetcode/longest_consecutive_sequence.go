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

func longest_consecutive_sequence(nums []int, n int) int {
	mp := map[int]int{}
	res := 0
	for i := 0; i < n; i++ {
		mp[nums[i]] = 1
	}
	// we want to get to the start of such sequence and start checking from there with increment of one.
	// for others, we just see if some consecutive lower element exist, if yes great and move to next element
	// if no lower consecutive element exist then probably this element is hte start of such sequence and then check till what length that sequence goes forward(at min it would be that element itself ie of length 1)
	// example input : 100, 4, 200, 1, 2, 3 ==> 4 (as in 1,2,3,4)
	for i := 0; i < n; {
		cnt := 0
		if mp[nums[i]-1] == 1 {
			i++
			continue
		} else {
			tmp := nums[i]
			for mp[tmp] == 1 {
				cnt++
				tmp++
			}
			res = max(res, cnt)
			i++
		}
	}

	return res
}

func main() {
	n := readInt()
	nums := []int{}
	for i := 0; i < n; i++ {
		nums = append(nums, readInt())
	}

	len := longest_consecutive_sequence(nums, n)
	fmt.Println("Length of longest consecutive sequence : ", len)
}
