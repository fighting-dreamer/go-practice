package main

import (
	"fmt"
	"log"
	"math"

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

func distance(p pair.Pair) float64 {
	return math.Hypot(p.Front.(float64), p.Back.(float64))
}

func k_closest_point_from_origin(nums []pair.Pair, n, k int) []pair.Pair {
	distPair := make([]pair.Pair, n)
	for i := 0; i < n; i++ {
		distPair[i] = *pair.MakePair(distance(nums[i]), nums[i])
	}
	maxHeap := priorityqueue.New[pair.Pair](comparator.Reverse[pair.Pair](func(a, b pair.Pair) int {
		af := a.Front.(float64)
		bf := b.Front.(float64)
		if af == bf {
			return 0
		} else {
			if af < bf {
				return -1
			}
			return 1
		}
	}))
	res := []pair.Pair{}
	for i := 0; i < n; i++ {
		maxHeap.Push(distPair[i])
		if maxHeap.Size() > k {
			maxHeap.Pop()
		}
	}
	for !maxHeap.Empty() {
		res = append(res, maxHeap.Pop().Back.(pair.Pair))
	}

	return res
}

func main() {
	n, k := readInt(), readInt()
	nums := make([]pair.Pair, n)
	for i := 0; i < n; i++ {
		nums[i] = *pair.MakePair(float64(readInt()), float64(readInt()))
	}

	res := k_closest_point_from_origin(nums, n, k)
	fmt.Println(res)
}
