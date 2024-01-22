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

// 2,3,4,5,18,17,6 ==> expected output : 17

func container_water(heights []int, n int) int {
	currentLeft := 0
	currentRight := n - 1
	currentWater := 0
	maxWater := 0

	for currentLeft < currentRight {
		currentWater = (currentRight - currentLeft) * min(heights[currentLeft], heights[currentRight])
		maxWater = max(maxWater, currentWater)
		fmt.Println(heights[currentLeft], heights[currentRight], currentWater, maxWater)
		if heights[currentLeft] < heights[currentRight] {
			currentLeft++	
		} else {
			currentRight--
		}
	}

	return maxWater
}

func main() {
	n := readInt()
	heights := []int{}
	for i := 0; i < n; i++ {
		heights = append(heights, readInt())
	}

	waterQty := container_water(heights, n)
	fmt.Println(waterQty)
}
