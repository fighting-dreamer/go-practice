package main

// true : [2 3 1 1 4]
// false : [3 2 1 0 4]
func jump_game(arr []int, n int) bool {
	farthest := 0
	for i := 0; i < n; i++ {
		if i > farthest {
			return false
		}

		if arr[i]+i > farthest {
			farthest = arr[i] + i
		}

		if farthest >= n-1 {
			return true
		}
	}

	return false
}

func jump_game_2(arr []int, n int) int {
	farthest := 0
	end := 0
	jumps := 0

	for i := 0; i < n-1; i++ {
		if arr[i]+i > farthest {
			farthest = arr[i] + i
		}

		// if i reach end, i will jump
		if i == end {
			// if we are not able to go and reach end
			if farthest <= i {
				return -1 // we cant go beyond i, we are stuck
			}

			jumps++
			end = farthest // we dont need to make more jumps till we reach farthes possible till ith position

			if end >= n-1 {
				break
			}
		}
	}

	return jumps
}

func main() {
	n := readInt()
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = readInt()
	}

	jump_game(arr, n)
}
