package main

import (
	"fmt"
	"log"
	"slices"
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

func group_anagrams(arr []string, n int) [][]string {
	mp := map[string][]string{}
	for i := 0; i < n; i++ {
		sortedStr := []byte(arr[i])
		slices.Sort(sortedStr)
		mp[string(sortedStr)] = append(mp[string(sortedStr)], arr[i])
	}

	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}

func main() {
	n := readInt()
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		arr[i] = readString()
	}

	fmt.Printf("group_anagrams(arr, n): %v\n", group_anagrams(arr, n))
}
