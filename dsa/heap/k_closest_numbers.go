package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/pair"
	"github.com/liyue201/gostl/ds/priorityqueue"
	"github.com/liyue201/gostl/utils/comparator"
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

func k_closest_number(nums []int, n, k, t int) []int {
	res := []int{}
	maxHeap := priorityqueue.New[pair.Pair](comparator.Reverse[pair.Pair](func(a, b pair.Pair) int {
		af := a.Front.(int)
		bf := b.Front.(int)
		if af == bf {
			return 0
		} else {
			if af < bf {
				return -1
			}
			return 1
		}
	}))
	numsPairs := make([]pair.Pair, n)
	for i := 0; i < n; i++ {
		if nums[i] < t {
			numsPairs[i] = *pair.MakePair(t-nums[i], 1)
		} else {
			numsPairs[i] = *pair.MakePair(nums[i]-t, -1)
		}
	}
	for i := 0; i < n; i++ {
		maxHeap.Push(numsPairs[i])
		if maxHeap.Size() > k {
			maxHeap.Pop()
		}
	}
	for !maxHeap.Empty() {
		p := maxHeap.Pop()
		if p.Back == -1 {
			res = append(res, p.Front.(int)+t)
		} else {
			res = append(res, t-p.Front.(int))
		}
	}

	return res
}

func main() {
	n, k, t := readInt(), readInt(), readInt()

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = readInt()
	}

	fmt.Println(k_closest_number(nums, n, k, t))
}
