package main

func dfs(src int, neighbour [][]int, g []map[int]int, vis []bool, changes *int) {
	for i := 0; i < len(neighbour[src]); i++ {
		v := neighbour[src][i]

		if vis[v] == true {
			continue
		}
		if g[v][src] == 0 {
			(*changes)++
		}
		vis[v] = true
		dfs(v, neighbour, g, vis, changes)
	}
}

func minReorder(n int, connections [][]int) int {

	// I need to know the neighbouring nodes that I can reach or they can reach for a given node.
	// for each neighbour, If we can reach that neighbour => we need a edge reversal/re-routing => change++
	// we keep on looking for every neighbour.
	neighbour := make([][]int, n)
	for i := 0; i < n; i++ {
		neighbour[i] = []int{}
	}

	g := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		g[i] = map[int]int{}
	}
	for i := 0; i < len(connections); i++ {
		u := connections[i][0]
		v := connections[i][1]

		g[u][v] = 1
		neighbour[u] = append(neighbour[u], v)
		neighbour[v] = append(neighbour[v], u)
	}
	vis := make([]bool, n)

	vis[0] = true
	changes := 0
	dfs(0, neighbour, g, vis, &changes)
	return changes
}
