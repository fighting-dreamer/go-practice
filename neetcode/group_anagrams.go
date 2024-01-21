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

func main() {
	n := readInt()
	mp := map[string][]string{}
	for i := 0; i < n; i++ {
		s := readString()
		ss := []byte(s)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] < ss[j]
		})

		mp[string(ss)] = append(mp[string(ss)], s)
	}

	for _, v := range mp {
		fmt.Println(v)
	}
}
