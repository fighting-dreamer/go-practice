package main

func canVisitRooms(rooms [][]int) bool {
	// BFS approach
	q := []int{}
	v := len(rooms)
	vis := make([]bool, v)

	q = append(q, 0)
	vis[0] = true
	for len(q) != 0 {
		top := q[0]
		q = q[1:] // q.pop()

		// if vis[top] == true {
		// 	continue
		// }

		for i := 0; i < len(rooms[top]); i++ {
			if vis[rooms[top][i]] == true {
				continue
			}
			q = append(q, rooms[top][i])
			vis[rooms[top][i]] = true
		}
	}

	cnt := 0
	for i := 0; i < v; i++ {
		if vis[i] {
			cnt++
		}
	}

	return cnt == v
}
