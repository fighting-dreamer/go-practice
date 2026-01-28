package main

import (
	"fmt"
	"log"
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

func task_scheduler(arr []byte, n int, interval int) []byte {
	taskMp := map[byte]int{}
	intervalMap := map[byte]int{}
	for i := 0; i < n; i++ {
		taskMp[arr[i]]++
	}
	for k, _ := range taskMp {
		intervalMap[k] = 0
	}
	res := []byte{}
	time := 0
	for len(taskMp) > 0 {
		bestTask := byte('-')
		maxCount := -1
		for k, v := range taskMp {
			if intervalMap[k] <= time {
				if v > maxCount {
					maxCount = v
					bestTask = k
				}
			}
		}
		if bestTask != '-' {
			res = append(res, bestTask)
			taskMp[bestTask]--
			if taskMp[bestTask] == 0 {
				delete(taskMp, bestTask)
			}
			intervalMap[bestTask] = time + interval + 1
		} else {
			res = append(res, '0')
		}
		time++
	}

	return res
}

func main() {
	n := readInt()
	s := readString()

	interval := readInt()

	res := task_scheduler([]byte(s), n, interval)
	for i := 0; i < len(res); i++ {
		fmt.Printf("%c ", res[i])
	}
}
