package main

import (
	"fmt"
	"log"

	"github.com/liyue201/gostl/ds/pair"
	"github.com/liyue201/gostl/ds/priorityqueue"
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
	g.adj[u] = append(g.adj[u], &Node[T]{Id: v, Wt: wt})
}

func (g *Graph[T]) printGraphAdjList() {
	for i := 0; i < g.V; i++ {
		for j := 0; j < len(g.adj[i]); j++ {
			n := g.adj[i][j]
			fmt.Println(i, n.Id, n.Wt)
		}
	}
}

func dijkstra(g *Graph[int], source int) map[int]int {
	distanceMapRes := map[int]int{}
	minHeap := priorityqueue.New[pair.Pair](func(a, b pair.Pair) int {
		if a.Back == b.Back {
			return 0
		} else {
			if a.Back.(int) < b.Back.(int) {
				return -1
			} else {
				return 1
			}
		}
	})

	distance := make([]int, g.V)
	vis := make([]int, g.V)
	for i := 0; i < g.V; i++ {
		distance[i] = -1 // -1 => infinity
	}

	distance[source] = 0
	minHeap.Push(*pair.MakePair(source, 0))
	count := 0
	for !minHeap.Empty() { // must loop till minHeap is empty coz there is possibility of some node coming in queue twice and have min distances or getting processed twice, we can eliminate that using "vis" visited flag filter
		top := minHeap.Pop()
		u := top.Front.(int)
		count++
		fmt.Println(count, distance, u, vis[u])
		if vis[u] == 1 {
			continue
		}
		vis[u] = 1 // we have visited the node u and it is now having the shortest distance possible.

		// for each adjoining vertex of u
		for i := 0; i < len(g.Adj()[u]); i++ {
			v := g.Adj()[u][i]

			if distance[v.Id] == -1 || distance[v.Id] > distance[u]+v.Wt {
				distance[v.Id] = distance[u] + v.Wt
				minHeap.Push(*pair.MakePair(v.Id, distance[v.Id]))
			}
		}
	}

	fmt.Println(count)
	for i := 0; i < g.V; i++ {
		distanceMapRes[i] = distance[i]
	}
	return distanceMapRes
}

/*
V = 9, E = 14
u v wt
0 1 4
0 7 8
1 7 11
1 2 8
7 8 7
7 6 1
2 8 2
8 6 6
2 3 7
2 5 4
6 5 2
3 5 14
3 4 9
5 4 10

result :
0 4 12 19 21 11 9 8 14
*/
func main() {
	v, e := readInt(), readInt()
	g := NewGraph[int](v, e)
	for i := 0; i < e; i++ {
		a, b, wt := readInt(), readInt(), readInt()
		g.AddEdge(a, b, wt)
	}
	distanceMap := dijkstra(g, 0)

	// fmt.Println(distanceMap)
	for i := 0; i < g.V; i++ {
		fmt.Print(distanceMap[i], " ")
	}
}
