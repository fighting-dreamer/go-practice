package dsa

// package main

import (
	"fmt"
)

type Node[T any] struct {
	Id int
	Wt T
}

type Graph[T any] struct {
	V   int
	E   int
	adj [][]*Node[T]
}

func (g *Graph[T]) Adj() [][]*Node[T] {
	return g.adj
}

func NewGraph[T any](v int, e int) *Graph[T] {
	return &Graph[T]{
		V:   v,
		E:   e,
		adj: make([][]*Node[T], v),
	}
}

func (g *Graph[T]) AddEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Node[T]{
		Id: v,
		Wt: wt,
	})
	g.adj[v] = append(g.adj[v], &Node[T]{
		Id: u,
		Wt: wt,
	})
}

func (g *Graph[T]) AddDirectedEdge(u, v int, wt T) {
	if u < 0 || u >= g.V || v < 0 || v >= g.V {
		return
	}
	g.adj[u] = append(g.adj[u], &Node[T]{
		Id: v,
		Wt: wt,
	})
}

func (g *Graph[T]) printGraphAdjList() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			n := g.adj[i][j]
			fmt.Println(i, n.Id, n.Wt)
		}
	}
}

func (g *Graph[T]) BFSTraversal() {
	var q []int
	vis := make([]bool, g.V)

	q = append(q, 0)
	vis[0] = true
	for len(q) > 0 {
		u := q[0]
		q = q[1:]

		fmt.Println(u)

		for i := 0; i < len(g.adj[u]); i++ {
			v := g.adj[u][i].Id
			if vis[v] == false {
				q = append(q, v)
				vis[v] = true
			}
		}
	}
}

func dfsHelper[T any](vis []int, u int, g *Graph[T]) {
	vis[u] = 1

	fmt.Println(u)

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id
		if vis[v] == 0 {
			dfsHelper(vis, v, g)
		}
	}
}

func (g *Graph[T]) DFSTraversal() {
	vis := make([]int, g.V)

	for i := 0; i < g.V; i++ {
		if vis[i] == 1 {
			continue
		}
		dfsHelper(vis, i, g)
	}
}

func NewUndirectedGraphWithCycle() *Graph[int] {
	g := NewGraph[int](5, 6)
	g.AddEdge(0, 1, 1)
	g.AddEdge(0, 4, 1)
	g.AddEdge(1, 3, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	return g
}

func NewUndirectedGraphWithOutCycle() *Graph[int] {
	g := NewGraph[int](5, 4)
	g.AddEdge(0, 1, 1)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 1)
	g.AddEdge(3, 4, 1)

	return g
}

func NewDirectedGraphWithCycle() *Graph[int] {
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

func NewDirectedGraphNoCycle() *Graph[int] {
	g := NewGraph[int](5, 4)

	/*
		1 -> 0 -> 3
		\u2B06\uFE0F   \u2197\uFE0F  \u2196\uFE0F   \u2B06\uFE0F
		2       4

	*/

	g.AddDirectedEdge(1, 0, 1)
	g.AddDirectedEdge(0, 3, 1)
	g.AddDirectedEdge(2, 1, 1)
	// g.AddDirectedEdge(0, 2, 1)
	// g.AddDirectedEdge(4, 0, 1)
	g.AddDirectedEdge(3, 4, 1)

	return g
}

func DetectCycleInDirectedGraphUsingDFSHelper(g *Graph[int], u int, vis []int, path map[int]int) bool {
	vis[u] = 1
	path[u] = 1
	defer func() {
		delete(path, u)
	}()

	for i := 0; i < len(g.adj[u]); i++ {
		v := g.adj[u][i].Id

		if path[v] == 1 {
			return true
		}

		if vis[v] == 1 {
			continue
		}

		if DetectCycleInDirectedGraphUsingDFSHelper(g, v, vis, path) {
			return true
		}
	}

	return false
}

func DetectCycleInDirectedGraphUsingDFS(g *Graph[int]) bool {
	vis := make([]int, g.V)
	path := map[int]int{}

	for i := 0; i < g.V; i++ {
		if vis[i] == 0 && DetectCycleInDirectedGraphUsingDFSHelper(g, i, vis, path) {
			return true
		}
	}
	return false
}

func DetectCycleInDirectedGraphUsingBFS(g *Graph[int]) bool {
	q := []int{}
	indegree := make([]int, g.V)

	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			v := g.adj[i][j].Id
			indegree[v]++
		}
	}

	for i := 0; i < g.V; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	terminalNodesCount := 0

	for len(q) > 0 {
		front := q[0]
		terminalNodesCount++
		q = q[1:]

		for i := 0; i < len(g.adj[front]); i++ {
			v := g.adj[front][i].Id
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return !(terminalNodesCount == g.V) // we were able to reach all nodes in graph => no cycle ie terminalNodesCount == total vertices in graph => no cycle so opposite of that tells is we detected a cycle.
}

func DetectCycleInDirectedGraph(g *Graph[int]) bool {
	// using DFS
	res1 := DetectCycleInDirectedGraphUsingDFS(g)
	// using BFS : topological sorting
	res2 := DetectCycleInDirectedGraphUsingBFS(g)
	fmt.Println("cycle detection in directed graph using DFS : ", res1)
	fmt.Println("cycle detection in directed graph using BFS : ", res2)

	return res2
}

func TopologicalSortDirectedGraph(g *Graph[int]) []int {
	q := []int{}
	indegree := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			v := g.adj[i][j].Id
			indegree[v]++
		}
	}
	for i := 0; i < g.V; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	res := []int{}

	for len(q) > 0 {
		front := q[0]
		q = q[1:]

		// process the node ie
		res = append(res, front)
		// for every edge from front to X, decrement the indegree for X
		for i := 0; i < len(g.adj[front]); i++ {
			v := g.adj[front][i].Id
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	return res
}

func DetectCycleInUndirectedGraphUsingDFSHelper(g *Graph[int], curr int, parent int, vis []int) bool {
	vis[curr] = 1

	for i := 0; i < len(g.adj[curr]); i++ {
		v := g.adj[curr][i].Id
		if v == parent {
			continue
		}
		if vis[v] == 1 {
			return true
		}
		if DetectCycleInUndirectedGraphUsingDFSHelper(g, v, curr, vis) {
			return true
		}
	}

	return false
}

func DetectCycleInUndirectedGraphUsingDFS(g *Graph[int]) bool {
	vis := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		if vis[i] == 1 {
			continue
		}
		if DetectCycleInUndirectedGraphUsingDFSHelper(g, i, -1, vis) {
			return true
		}
	}
	return false
}

func DetectCycleInUndirectedGraph(g *Graph[int]) bool {
	res1 := DetectCycleInUndirectedGraphUsingDFS(g)
	return res1
}

func main() {
	g := NewUndirectedGraphWithCycle()

	g.printGraphAdjList()
	fmt.Println()
	g.BFSTraversal()
	fmt.Println()
	g.DFSTraversal()

	fmt.Println("-------------------------------------")
	g = NewDirectedGraphWithCycle()
	fmt.Println("Cycle in directed graph : ", DetectCycleInDirectedGraph(g))
	fmt.Println("Topological sort directed graph with cycle : ", TopologicalSortDirectedGraph(g))
	g = NewDirectedGraphNoCycle()
	fmt.Println("Cycle in directed graph : ", DetectCycleInDirectedGraph(g))
	fmt.Println("Topological sort directed graph with no cycle : ", TopologicalSortDirectedGraph(g))

	g = NewUndirectedGraphWithCycle()
	fmt.Println("Cycle in undirected graph : ", DetectCycleInUndirectedGraph(g))
	g = NewUndirectedGraphWithOutCycle()
	fmt.Println("Cycle in undirected graph : ", DetectCycleInUndirectedGraph(g))
}
