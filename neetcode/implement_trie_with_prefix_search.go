package main

import (
	"fmt"
	"log"
	"strings"
)

type TrieNode struct {
	Val   byte
	child [26]*TrieNode
}

func AddString(root *TrieNode, s string) {
	for i := 0; i < len(s); i++ {
		child := root.child[s[i]-'A']
		if child == nil {
			root.child[s[i]-'A'] = &TrieNode{
				Val: s[i],
			}
		}
		root = root.child[s[i]-'A']
	}
}

func PrefixSearchString(root *TrieNode, s string) bool {
	for i := 0; i < len(s); i++ {
		if root.child[s[i]-'A'] == nil {
			return false
		}
		root = root.child[s[i]-'A']
	}
	return true
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

func main() {
	s := strings.ToUpper(readString())
	root := &TrieNode{
		Val: byte('-'),
	}
	AddString(root, s)
	fmt.Println(PrefixSearchString(root, s))
	// should be true too :
	fmt.Println(PrefixSearchString(root, s[:len(s)-2]))
	// should be false
	fmt.Println(PrefixSearchString(root, strings.ToUpper("fnwnwlgwlg")))
}
