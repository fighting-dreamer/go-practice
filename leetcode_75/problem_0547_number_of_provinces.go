package main

func do_bfs(src int, vis []bool, isConnected [][]int) {
	q := []int{}

	q = append(q, src)
	vis[src] = true
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		for i := 0; i < len(isConnected[top]); i++ {
			if isConnected[top][i] == 1 && vis[i] == false {
				q = append(q, i)
				vis[i] = true
			}
		}
	}

	return 
}

func findNumberOfProvinces(isConnected [][]int) int {
	v := len(isConnected)
	vis := make([]bool, v)
	res := 0
	for i := 0; i < v; i++ {
		if vis[i] {
			continue
		}
		do_bfs(i, vis, isConnected)
		res++
	}

	return res
}
