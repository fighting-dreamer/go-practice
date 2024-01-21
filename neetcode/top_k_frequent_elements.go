package main

import (
	"cmp"
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/pair"
	"github.com/liyue201/gostl/ds/priorityqueue"
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
	topK := readInt()
	fmt.Println(topK)

	mp := map[int]int{}
	for i := 0; i < n; i++ {
		mp[readInt()]++
	}

	pq := priorityqueue.New[pair.Pair](func(a, b pair.Pair) int {
		return cmp.Compare[int](a.Back.(int), b.Back.(int))
	}, priorityqueue.WithGoroutineSafe())

	for k, v := range mp {
		pq.Push(*pair.MakePair(k, v))
		if pq.Size() > topK {
			pq.Pop()
		}
	}
	res := []int{}
	for {
		if pq.Empty() {
			break
		}
		p := pq.Pop()
		res = append(res, p.Front.(int))
	}
	fmt.Println("result :\n", res)
	return
}
