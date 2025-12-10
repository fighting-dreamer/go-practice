package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"
)

type FKPair struct {
	Weight int
	Val    int
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

func density(val, wt int) float64 {
	return float64(val) / float64(wt)
}

func fractional_knapsack(arr []FKPair, n int, cap int) float64 {
	res := 0.0 // max Profit

	slices.SortFunc(arr, func(a FKPair, b FKPair) int {
		return -cmp.Compare(density(a.Val, a.Weight), density(b.Val, b.Weight)) // descending order
	})
	
	fmt.Println(arr)

	remainingCap := cap
	for i := 0; i < n; i++ {
		if remainingCap < arr[i].Weight {
			res += float64(remainingCap) * float64(arr[i].Val) / float64(arr[i].Weight)
			remainingCap = 0
			break
		} else {
			res += float64(arr[i].Val)
			remainingCap -= arr[i].Weight
		}
	}

	return res
}

func main() {
	n := readInt()
	cap := readInt()
	arr := make([]FKPair, n)
	for i := 0; i < n; i++ {
		arr[i] = FKPair{
			Val:    readInt(),
			Weight: readInt(),
		}
	}
	res := fractional_knapsack(arr, n, cap)
	fmt.Println(res)
}
