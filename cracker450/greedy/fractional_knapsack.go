package main

import (
	"cmp"
	"fmt"
	"log"
	"slices"

	"github.com/liyue201/gostl/ds/pair"
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

func fractional_knapsack(value []float64, weights []float64, n int, w float64) float64 {
	valuePairs := make([]pair.Pair, n)
	for i := 0; i < n; i++ {
		valuePairs[i] = *pair.MakePair(value[i], weights[i])
	}
	slices.SortFunc(valuePairs, func(a, b pair.Pair) int {
		if a.Front.(float64)/a.Back.(float64)-b.Front.(float64)/b.Back.(float64) < 1e-6 {
			return 0
		}
		return cmp.Compare(b.Front.(float64)/b.Back.(float64), a.Front.(float64)/b.Back.(float64))
	})

	profit := 0.0
	for i := 0; i < n; i++ {
		if w < valuePairs[i].Back.(float64) {
			profit += (valuePairs[i].Front.(float64) * w) / valuePairs[i].Back.(float64)
			break
		}
		profit += valuePairs[i].Front.(float64)
		w -= valuePairs[i].Back.(float64)
	}

	return profit
}

func main() {
	n, w := readInt(), readDouble()
	value := make([]float64, n)
	weights := make([]float64, n)
	for i := 0; i < n; i++ {
		value[i], weights[i] = readDouble(), readDouble()
	}
	fmt.Println(fractional_knapsack(value, weights, n, w))
}
