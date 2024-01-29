package main

import (
	"fmt"
	"log"

	"nipun.io/coding/dsa"
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

type GraphNode[T any] struct {
	Id int
	Wt T
}

func CloneGraph(g *dsa.Graph[int]) *dsa.Graph[int] {
	// we run through the graph to get to all nodes using DFS, while DFS, we keep a map of all nodes and for each node, we create a clone/duplicte node.
	// now for each node we have a mapping, so we iterate over each node and for the neighbour of that node in original graph we get the cloned node and add a edge to that

	//Duplicate vertices :  make a map in original graph nodes and new duplicate nodes.
	// duplicate edges : for each node in original graph's neighbour create or add a duplicate neighbour to the respective cloned/duplicate node.
	return nil
}

func main() {
	g := dsa.NewDirectedGraphWithCycle()
	g.DFSTraversal()

	clone := CloneGraph(g)
	clone.DFSTraversal()
}
