package main

// type Node[T any] struct {
// 	Id int
// 	Wt T
// }

// type Graph[T any] struct {
// 	V   int
// 	E   int
// 	adj [][]*Node[T]
// }

// func (g *Graph[T]) Adj() [][]*Node[T] {
// 	return g.adj
// }
// func NewGraph[T any](v int, e int) *Graph[T] {
// 	return &Graph[T]{
// 		V:   v,
// 		E:   e,
// 		adj: make([][]*Node[T], v),
// 	}
// }

// func (g *Graph[T]) AddEdge(u, v int, wt T) {
// 	if u < 0 || u >= g.V || v < 0 || v >= g.V {
// 		return
// 	}
// 	g.adj[u] = append(g.adj[u], &Node[T]{
// 		Id: v,
// 		Wt: wt,
// 	})
// 	g.adj[v] = append(g.adj[v], &Node[T]{
// 		Id: u,
// 		Wt: wt,
// 	})
// }
// func (g *Graph[T]) AddDirectedEdge(u, v int, wt T) {
// 	if u < 0 || u >= g.V || v < 0 || v >= g.V {
// 		return
// 	}
// 	g.adj[u] = append(g.adj[u], &Node[T]{Id: v, Wt: wt})
// }

// func (g *Graph[T]) printGraphAdjList() {
// 	for i := 0; i < g.V; i++ {
// 		for j := 0; j < len(g.adj[i]); j++ {
// 			n := g.adj[i][j]
// 			fmt.Println(i, n.Id, n.Wt)
// 		}
// 	}
// }

// func dfs(src int, g *Graph[float64], vis []bool, val []float64) {
// 	vis[src] = true

// 	for i := 0; i < len(g.Adj()[src]); i++ {
// 		v := g.Adj()[src][i]

// 		if vis[v.Id] == false {
// 			val[v.Id] = v.Wt * val[src]
// 			dfs(v.Id, g, vis, val)
// 		}
// 	}
// }

// func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
// 	symbols := map[string]int{}
// 	for i := 0; i < len(equations); i++ {
// 		symbols[equations[i][0]] = 1
// 		symbols[equations[i][1]] = 1
// 	}
// 	v := len(symbols)
// 	e := len(equations)
// 	i := 0
// 	for k, _ := range symbols {
// 		symbols[k] = i
// 		i++
// 	}

// 	g := NewGraph[float64](v, e)
// 	for i := 0; i < e; i++ {
// 		// u / v = wt
// 		// v -> (u, wt)
// 		u := symbols[equations[i][0]]
// 		v := symbols[equations[i][1]]
// 		wt := values[i]
// 		g.AddDirectedEdge(u, v, wt)
// 	}

// 	indegree := make([]int, g.V)
// 	for i := 0; i < g.E; i++ {
// 		indegree[symbols[equations[i][1]]]++
// 	}

// 	vis := make([]bool, g.V)
// 	val := make([]float64, g.V)
// 	for i := 0; i < g.V; i++ {
// 		if indegree[i] == 0 {
// 			val[i] = 1.0
// 			dfs(i, g, vis, val)
// 		}
// 	}

// 	// for i := 0; i < g.V; i++ {
// 	// 	fmt.Println(i, val[i])
// 	// }

// 	res := make([]float64, len(queries))
// 	for i, q := range queries {
// 		id1, ok1 := symbols[q[0]]
// 		id2, ok2 := symbols[q[1]]
// 		if ok1 == false || ok2 == false {
// 			res = append(res, -1.0)
// 		} else {
// 			res[i] = val[id1] / val[id2]
// 		}
// 	}
// 	return res
// }

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// floydwarshall
	symbols := map[string]int{}
	for i := 0; i < len(equations); i++ {
		u := equations[i][0]
		v := equations[i][1]
		symbols[u] = 1
		symbols[v] = 1
	}
	i := 0
	for k, _ := range symbols {
		symbols[k] = i
		i++
	}
	n := len(symbols)
	distance := make([][]float64, n)
	for i := 0; i < n; i++ {
		distance[i] = make([]float64, n)
		distance[i][i] = 1.0
	}
	for i := 0; i < len(equations); i++ {
		u := equations[i][0]
		v := equations[i][1]
		distance[symbols[v]][symbols[u]] = values[i]
		distance[symbols[u]][symbols[v]] = 1.0 / values[i]
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if distance[i][j] != 0 {
				distance[i][j] = -1
			}
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distance[i][k] == -1 || distance[k][j] == -1 {
					continue
				}
				if distance[i][j] == -1 || distance[i][k]*distance[k][j] < distance[i][j] {
					distance[i][j] = distance[i][k] * distance[k][j]
				}
			}
		}
	}

	res := make([]float64, len(queries))
	for i, q := range queries {
		u := symbols[q[0]]
		v := symbols[q[1]]

		if distance[v][u] != -1 {
			res[i] = distance[v][u]
		}else {
			res[i] = -1.0
		}
	}

	return res
}
