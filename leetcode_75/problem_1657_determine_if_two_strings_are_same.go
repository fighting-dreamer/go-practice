package main

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"
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

func check_if_they_are_similar(s1, s2 string) bool {
	// crux is that we check if they are similar by checking if all related conditions that are quired comply else it must be false only.
	// we check the characters that are present in s1 are present in s2
	// we check the frequencies numbers of characters are same even if the characters are not same.
	// example : aabbbbccc bbccccaaa here the frequency : [2,4,3] and [2,4,3] is same even if the characters of a given freuency are not same.

	if len(s1) != len(s2) {
		return false
	}

	s1_freqmap := map[byte]int{}
	s2_freqmap := map[byte]int{}
	for i := 0; i < len(s1); i++ {
		s1_freqmap[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		s2_freqmap[s2[i]]++
	}
	// fmt.Println(s1_freqmap)
	// fmt.Println(s2_freqmap)
	if len(s1_freqmap) != len(s2_freqmap) {
		return false
	}
	s11, s1_freqVec := strings.Builder{}, []int{}
	for k, count := range s1_freqmap {
		s11.WriteByte(k)
		s1_freqVec = append(s1_freqVec, count)
	}

	s22, s2_freqVec := strings.Builder{}, []int{}
	for k, count := range s1_freqmap {
		s22.WriteByte(k)
		s2_freqVec = append(s2_freqVec, count)
	}
	sa := []byte(s11.String())
	sb := []byte(s22.String())
	sort.Slice(sa, func(i, j int) bool {
		return sa[i] < sa[j]
	})
	sort.Slice(sb, func(i, j int) bool {
		return sb[i] < sb[j]
	})

	if string(sa) != string(sb) {
		return false
	}
	slices.Sort(s1_freqVec)
	slices.Sort(s2_freqVec)
	for i := 0; i < len(s1_freqVec); i++ {
		if s1_freqVec[i] != s2_freqVec[i] {
			return false
		}
	}
	return true
}

func main() {
	s1, s2 := readString(), readString()

	fmt.Println(check_if_they_are_similar(s1, s2))
}
