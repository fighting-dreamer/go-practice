package main

func main() {
	n := readInt()
	W := readInt()
	wt := make([]int, n)
	// val is same as wt

	for i := 0; i < n; i++ {
		wt[i] = readInt()
	}

	subset_sum(W, wt, n)
}
