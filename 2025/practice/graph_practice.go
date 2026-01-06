package main

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/liyue201/gostl/ds/deque"
	"github.com/liyue201/gostl/ds/priorityqueue"
)

type Edge[T any] struct {
	Id int
	Wt T
}

type Graph[T any] struct {
	V   int
	E   int
	adj [][]*Edge[T]
}

func (g *Graph[T]) Adj() [][]*Edge[T] {
	return g.adj
}
func NewGraph[T any](v int, e int) *Graph[T] {
	return &Graph[T]{
		V:   v,
		E:   e,
		adj: make([][]*Edge[T], v),
	}
}

func (g *Graph[T]) AddEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Edge[T]{
		Id: v,
		Wt: wt,
	})
	g.adj[v] = append(g.adj[v], &Edge[T]{
		Id: u,
		Wt: wt,
	})
}
func (g *Graph[T]) AddDirectedEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Edge[T]{Id: v, Wt: wt})
}
func (g *Graph[T]) printGraphAdjList() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			n := g.adj[i][j]
			fmt.Println(i, n.Id, n.Wt)
		}
	}
}

func undirected_graph() *Graph[int] {
	g := NewGraph[int](5, 4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	return g
}

func directed_graph() *Graph[int] {
	g := NewGraph[int](5, 4)

	g.AddDirectedEdge(1, 0, 1)
	g.AddDirectedEdge(0, 3, 1)
	g.AddDirectedEdge(2, 1, 1)
	// g.AddDirectedEdge(0, 2, 1)
	// g.AddDirectedEdge(4, 0, 1)
	g.AddDirectedEdge(3, 4, 1)

	return g
}

func undirected_graph_with_cycle() *Graph[int] {
	g := NewGraph[int](5, 6)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	return g
}

func directed_graph_with_cycle() *Graph[int] {
	g := NewGraph[int](5, 6)

	/*
		1 -> 0 -> 3
		\u2B06\uFE0F   \u2197\uFE0F  \u2196\uFE0F   \u2B06\uFE0F
		2       4

	*/

	g.AddDirectedEdge(1, 0, 1)
	g.AddDirectedEdge(0, 3, 1)
	g.AddDirectedEdge(2, 1, 1)
	g.AddDirectedEdge(0, 2, 1)
	g.AddDirectedEdge(4, 0, 1)
	g.AddDirectedEdge(3, 4, 1)

	return g
}

func directed_01_graph() *Graph[int] {
	g := NewGraph[int](9, 13)
	g.AddEdge(0, 1, 0)
	g.AddEdge(0, 7, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 7, 1)
	g.AddEdge(2, 3, 0)
	g.AddEdge(2, 5, 0)
	g.AddEdge(2, 8, 1)
	g.AddEdge(3, 4, 1)
	g.AddEdge(3, 5, 1)
	g.AddEdge(4, 5, 1)
	g.AddEdge(5, 6, 1)
	g.AddEdge(6, 7, 1)
	g.AddEdge(7, 8, 1)
	return g
}

func bfs_undirected_graph(g *Graph[int]) {
	g.printGraphAdjList()
	q := []int{}
	vis := make([]bool, g.V)

	// you put something in queue => mark it visited
	q = append(q, 0)
	vis[0] = true

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		fmt.Printf("%d ", u)
		for i := 0; i < len(g.adj[u]); i++ {
			v := g.adj[u][i].Id
			if vis[v] == false {
				q = append(q, v)
				vis[v] = true
			}
		}
	}
}
func bfs_directed_graph(g *Graph[int]) {
	g.printGraphAdjList()
	bfs_undirected_graph(g) // its same logic
}

func dfs_helper(g *Graph[int], curr int, vis []bool) {
	vis[curr] = true // mark visited
	fmt.Println(curr)

	for i := 0; i < len(g.adj[curr]); i++ {
		v := g.adj[curr][i].Id
		if vis[v] == false {
			dfs_helper(g, v, vis)
		}
	}
}

func dfs_undirected_graph(g *Graph[int]) {
	g.printGraphAdjList()
	vis := make([]bool, g.V)
	for i := 0; i < g.V; i++ {
		if vis[i] == false {
			dfs_helper(g, i, vis)
		}
	}
}

func dfs_directed_graph(g *Graph[int]) {
	g.printGraphAdjList()
	dfs_undirected_graph(g) // same logic
}

func dfs_helper_detect_cycle(g *Graph[int], curr int, vis []bool, path []bool, parent int) bool {
	vis[curr] = true  // mark visited
	path[curr] = true // mark path check
	defer func() {
		path[curr] = false
	}() // to execute at end of this function.

	for i := 0; i < len(g.adj[curr]); i++ {
		v := g.adj[curr][i].Id
		if parent == v {
			continue
		}
		if path[v] == true {
			// alredy seen this node => its a cycle
			return true
		}
		if vis[v] == true {
			continue // already visited, might be in other dfs_helper run
		}
		if dfs_helper_detect_cycle(g, v, vis, path, curr) {
			return true
		}
	}
	return false
}

func detect_cycle_undirected_graph(g *Graph[int]) {
	g.printGraphAdjList()
	// you can detect cycle in undirected graph using dfs only.
	// you keep track of path traveersed till now
	// if you encounter a node seen on that path =>
	// its a CYCLE as its possible to reach the same node again and again using same path

	path := make([]bool, g.V)
	vis := make([]bool, g.V)
	for i := 0; i < g.V; i++ {
		if vis[i] == false {
			if dfs_helper_detect_cycle(g, i, vis, path, -1) == true {
				fmt.Println("Cycle detected")
				return
			}
		}
	}
}

// its not essentially bfs, as we are not adding nodes that are next
// only the ones that are having indegree as zero
func detect_cycle_directed_graph_using_bfs(g *Graph[int]) bool {
	g.printGraphAdjList()
	indegree := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		for u := 0; u < len(g.adj[i]); u++ {
			v := g.adj[i][u].Id
			indegree[v]++
		}
	}

	q := []int{}
	for i := 0; i < g.V; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	terminalNodeCount := 0

	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		// fmt.Println(u)
		terminalNodeCount++

		for i := 0; i < len(g.adj[u]); i++ {
			v := g.adj[u][i].Id
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if terminalNodeCount != g.V {
		return true // its a cycle as terminal node count not reaching all nodes.
	}
	return false // no cycle
}

func dfs_helper_directed_graph_check_cycle(g *Graph[int], u int, vis []bool, path []bool) bool {
	vis[u] = true
	path[u] = true
	defer func() {
		path[u] = false
	}()

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id
		if path[v] == true {
			return true // you have reached a node that you have visite in path
		}
		if vis[v] == true {
			// already visited node, no need to continue for this
			continue
		}
		if dfs_helper_directed_graph_check_cycle(g, v, vis, path) {
			return true
		}
	}
	return false
}

func detect_cycle_directed_graph_using_dfs(g *Graph[int]) {
	g.printGraphAdjList()

	vis := make([]bool, g.V)
	path := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if vis[i] == false {
			if dfs_helper_directed_graph_check_cycle(g, i, vis, path) == true {
				fmt.Println("Cycle detected")
				return
			}
		}
	}
	fmt.Println("No Cycle")
}

func topological_sort_bfs(g *Graph[int]) {
	g.printGraphAdjList()
	indegree := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		for u := 0; u < len(g.adj[i]); u++ {
			v := g.adj[i][u].Id
			indegree[v]++
		}
	}
	q := []int{}
	for i := 0; i < g.V; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	res := []int{}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		res = append(res, u)

		for i := 0; i < len(g.adj[u]); i++ {
			v := g.adj[u][i].Id
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	fmt.Println(res)
}

func dfs_helper_topological_sort(g *Graph[int], u int, vis []bool, stack *[]int) {
	vis[u] = true

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id
		if vis[v] == false {
			dfs_helper_topological_sort(g, v, vis, res)
		}
	}
	// POST ORDER addition : we add all children before current node
	*stack = append(*stack, u)
}

func topological_sort_dfs(g *Graph[int]) {
	// POST order addition :
	// we add all children add before curr node
	// consider 0->1 and 2->1
	// right answer : [0,2,1]
	// in first go(2->1), we add [1,2]
	// in next round (0->1) : [1,2,0], 1 is already visited, not visited twice
	// reverse it.

	g.printGraphAdjList()
	stack := []int{}
	vis := make([]bool, g.V)

	for i := 0; i < g.V; i++ {
		if vis[i] == false {
			dfs_helper_topological_sort(g, i, vis, &stack)
		}
	}
	slices.Reverse(stack)
	fmt.Println(stack)
}

func dfs_helper_detect_cycle_topological_sort(g *Graph[int], u int, vis []int, stack *[]int) bool {
	// vis[i] = 0 => not yet visited
	// vis[i] = 1 => visited, its in path
	// vis[i] = 2 => completed processing
	vis[u] = 1 // mark visited

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id
		if vis[v] == 1 {
			// back edge detected => cycle exists
			return true
		}
		if vis[v] == 0 && !dfs_helper_detect_cycle_topological_sort(g, v, vis, stack) {
			return true
		}
	}
	vis[u] = 2
	*stack = append(*stack, u)
	return false
}

func topological_sort_detect_cycle_dfs(g *Graph[int]) bool {
	vis := make([]int, g.V)
	stack := []int{}

	for i := 0; i < g.V; i++ {
		if vis[i] == 0 && dfs_helper_detect_cycle_topological_sort(g, i, vis, &stack) {
			fmt.Println("Cycle Detected")
			return true
		}
	}
	slices.Reverse(stack)
	fmt.Println(stack)
	fmt.Println("No Cycle detected")
	return false
}

type CompleteEdge struct {
	U  int
	V  int
	Wt int
}

func NewCompleteEdge(u, v, wt int) *CompleteEdge {
	return &CompleteEdge{
		U:  u,
		V:  v,
		Wt: wt,
	}
}

func find(u int, par []int) int {
	if par[u] == u {
		return u
	}
	par[u] = find(par[u], par)
	return par[u]
}
func union(u, v int, par []int) {
	uf := find(u, par)
	vf := find(v, par)
	if uf != vf {
		par[uf] = vf
	}
}

func minimum_spaning_tree_kruskal(g *Graph[int]) {
	g.printGraphAdjList()
	minHeap := priorityqueue.New[*CompleteEdge](func(i *CompleteEdge, j *CompleteEdge) int {
		return cmp.Compare(i.Wt, j.Wt)
	})
	for i := 0; i < g.V; i++ {
		for _, edge := range g.adj[i] {
			minHeap.Push(NewCompleteEdge(i, edge.Id, edge.Wt))
		}
	}

	par := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		par[i] = i
	}

	resWt := 0
	resEdges := []CompleteEdge{}
	for !minHeap.Empty() {
		top := minHeap.Pop()
		uf := find(top.U, par)
		vf := find(top.V, par)

		if uf != vf {
			union(uf, vf, par)
			resWt += top.Wt
			resEdges = append(resEdges, *top)
		}
	}

	fmt.Println("Minimum spanning Tree Wt : ", resWt)
	for _, edge := range resEdges {
		fmt.Println(edge.U, edge.V)
	}
}

func minimum_spanning_tree_prim(g *Graph[int]) {
	g.printGraphAdjList()

	// at each step, you are adding a node that is not visited
	// and while you do that,
	// you only consider edges with minimum weight
	minHeap := priorityqueue.New[CompleteEdge](func(a, b CompleteEdge) int {
		return cmp.Compare(a.Wt, b.Wt)
	})
	vis := make([]bool, g.V)
	resWt := 0
	resEdges := []CompleteEdge{}
	addEdgesToPriorityQueue := func(u int) {
		vis[u] = true
		for _, edge := range g.adj[u] {
			minHeap.Push(*NewCompleteEdge(u, edge.Id, edge.Wt))
		}
	}

	source := 0 // any vertex
	addEdgesToPriorityQueue(source)
	for !minHeap.Empty() && len(resEdges) < g.V-1 {
		minEdge := minHeap.Pop()

		if vis[minEdge.V] {
			// already added and visited
			continue
		}

		resWt += minEdge.Wt
		resEdges = append(resEdges, minEdge)

		addEdgesToPriorityQueue(minEdge.V)
	}

	fmt.Print(resWt)
	fmt.Println(resEdges)
}

func minimum_distance_from_source_dijkstra(g *Graph[int]) {
	g.printGraphAdjList()
	type Item struct {
		u    int
		dist int
	}
	NewItem := func(u int, d int) *Item {
		return &Item{
			u:    u,
			dist: d,
		}
	}
	vis := make([]bool, g.V)
	distance := make([]int, g.V)
	parent := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = -1 // represents infinity
		parent[i] = -1   // we dont know parent[i] yet
	}

	minHeap := priorityqueue.New[Item](func(a, b Item) int {
		return cmp.Compare(a.dist, b.dist)
	})

	source := 0

	distance[source] = 0
	minHeap.Push(*NewItem(source, 0))

	visited := 0

	for !minHeap.Empty() {
		curr := minHeap.Pop()
		u := curr.u

		if vis[u] {
			continue
		}

		vis[u] = true
		visited++

		if visited == g.V {
			break
		}

		for _, edge := range g.adj[u] {
			v := edge.Id
			newDist := edge.Wt + distance[u]

			if distance[v] == -1 || newDist < distance[v] {
				distance[v] = newDist
				parent[v] = u
				minHeap.Push(*NewItem(v, newDist))
			}
		}
	}

	fmt.Println("distance : ", distance)
	fmt.Println("parent : ", parent)
}

func getPathDijkastraShortestParentGiven(source int, dest int, parent []int) []int {
	path := []int{}
	if parent[dest] == -1 {
		// cant be reached => empty path
		return path
	}

	curr := dest
	for curr != -1 {
		path = append(path, curr)
		curr = parent[curr]
	}

	slices.Reverse(path)
	fmt.Println(path)
	return path
}

func minimum_distance_from_source_01_bfs(g *Graph[int]) {
	g.printGraphAdjList()
	// type Item struct {
	// 	u int
	// 	wt int
	// }

	dq := deque.New[int]()
	distance := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = -1
	}

	source := 0 // can be any vertex
	distance[source] = 0
	dq.PushBack(source)

	for !dq.Empty() {
		top := dq.PopFront()

		for _, edge := range g.adj[top] {
			if distance[edge.Id] > distance[top]+edge.Wt {
				distance[edge.Id] = distance[top] + edge.Wt
				if edge.Wt&1 == 1 {
					dq.PushBack(edge.Id)
				} else {
					dq.PushFront(edge.Id)
				}
			}
		}
	}

}

func minimum_distance_from_source_bellman_ford(g *Graph[int]) {
	g.printGraphAdjList()
	// relaing edges for V - 1 times to get to min-distance from source
	// if further iteration also lead to further relaxation => negtive cycle

	distance := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = -1 // infinity
	}

	source := 0 // can be any vertex (0 indexed)
	distance[source] = 0

	for i := 0; i < g.V-1; i++ { // if you let the loop for "i < g.V", then add check in relaxation for "i == g.V-1"
		for uu := 0; uu < g.V; uu++ {
			if distance[uu] == -1 {
				continue
			}
			for _, edge := range g.adj[uu] {
				v := edge.Id
				wt := edge.Wt

				if distance[v] == -1 || distance[uu]+wt < distance[v] {
					// if i == g.V-1 { // you can take this negtive cycle check outside as shown down further.
					// 	fmt.Println("Negtive cycle detected") // this wont occur if you run till (i < g.V-1)
					// }
					distance[v] = distance[uu] + wt
				}
			}
		}
	}

	// checking negtive cycle by continuing for one more run
	for u := 0; u < g.V; u++ {
		for _, edge := range g.adj[u] {
			v := edge.Id
			wt := edge.Wt
			if distance[v] == -1 || distance[u]+wt < distance[v] {
				fmt.Printf("Negtive Cycle in the graph detected")
			}
		}
	}
}

func minimum_distance_any_two_nodes(g *Graph[int]) {
	g.printGraphAdjList()

	distance := make([][]int, g.V)

	for i := 0; i < g.V; i++ {
		distance[i] = make([]int, g.V)
		for j := 0; j < g.V; j++ {
			if i == j {
				distance[i][j] = 0
			} else {
				distance[i][j] = -1
			}
		}
	}

	for u := 0; u < g.V; u++ {
		for _, v := range g.adj[u] {
			distance[u][v.Id] = v.Wt
		}
	}

	for u := 0; u < g.V; u++ {
		for v := 0; v < g.V; v++ {
			for k := 0; k < g.V; k++ {
				// if path exists through k from u to v
				if distance[u][k] != -1 && distance[k][v] != -1 {
					if distance[u][v] == -1 || distance[u][k]+distance[k][v] < distance[u][v] {
						distance[u][v] = distance[u][k] + distance[k][v]
					}
				}
			}
		}
	}
}

func main() {
	// bfs_undirected_graph(undirected_graph())
	// bfs_directed_graph(directed_graph())
	// dfs_undirected_graph(undirected_graph())
	// dfs_directed_graph(directed_graph())
	// detect_cycle_undirected_graph(undirected_graph_with_cycle())
	// detect_cycle_directed_graph_using_bfs(directed_graph_with_cycle())
	// detect_cycle_directed_graph_using_dfs(directed_graph_with_cycle())
	// topological_sort_bfs(directed_graph())
	// topological_sort_dfs(directed_graph())
	// topological_sort_detect_cycle_dfs(directed_graph_with_cycle(g))
	// minimum_spaning_tree_kruskal(undirected_graph())
	// minimum_spanning_tree_prim(undirected_graph())
	// minimum_distance_from_source_dijkstra(directed_graph())
	// getPathDijkastraShortestParentGiven() // best
	// minimum_distance_from_source_01_bfs(directed_01_graph())
	// minimum_distance_from_source_bellman_ford(directed_graph())
	// minimum_distance_any_two_nodes(directed_graph())

	// directed_01_graph().printGraphAdjList()
}
