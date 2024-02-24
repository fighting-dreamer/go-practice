package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/stack"
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

func asteroid_collision(arr []int, n int) []int {
	st := stack.New[int]()
	for i := 0; i < n; i++ {
		if st.Empty() {
			st.Push(arr[i])
			continue
		}
		// + - ==> collision
		// + + ==> no collision
		// - + ==> no collision
		// - - ==> no collision
		if st.Top() > 0 && arr[i] < 0 {
			// collision scenario
			wt := -arr[i] // arr[i] is negtive and its weight is -1 * arr[i]
			for !st.Empty() && st.Top() < wt {
				st.Pop()
			}
			if st.Empty() {
				st.Push(arr[i])
				continue
			}
			if st.Top() == wt {
				st.Pop()
				continue
			}
		} else {
			st.Push(arr[i])
		}
	}
	res := make([]int, st.Size())
	i := st.Size() - 1
	for !st.Empty() {
		res[i] = st.Pop()
		i--
	}

	return res
}

func main() {
	n := readInt()
	arr := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	fmt.Println(asteroid_collision(arr, n))
}
