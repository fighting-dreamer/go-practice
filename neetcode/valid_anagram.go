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
	str1 := []rune(readString())
	str2 := []rune(readString())
	sort.Slice(str1, func(i, j int) bool {
		return str1[i] < str1[j]
	})
	sort.Slice(str2, func(i, j int) bool {
		return str2[i] < str2[j]
	})
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
	return
}
