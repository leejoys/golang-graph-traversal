package main

import (
	"fmt"
	"testing"
)

func TestDfs(t *testing.T) {
	// graph
	g := NewDirectedGraph()

	// ErrNotAllVertExist test
	err := g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	// create vertices
	for i := 0; i < 8; i++ {
		err := g.AddVertex(i + 1)
		if err != nil {
			fmt.Println(err)
		}
	}

	// ErrVertExist test
	err = g.AddVertex(5)
	if err != nil {
		fmt.Println(err)
	}

	// begin create edges
	err = g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 4)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 1)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 5)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 6)
	if err != nil {
		fmt.Println(err)
	}
	err = g.AddEdge(6, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 7)
	if err != nil {
		fmt.Println(err)
	}
	// end of edges

	path, err := ShortestBFS(1, 6, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	fmt.Println("")

	path, err = ShortestBFS(6, 3, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	// fmt.Println(g.Vertices[1])
	// fmt.Println(g.Vertices[2])
	// fmt.Println(g.Vertices[3])
	// fmt.Println(g.Vertices[4])
	// fmt.Println(g.Vertices[5])
	// fmt.Println(g.Vertices[6])
	// fmt.Println(g.Vertices[7])
	// fmt.Println(g.Vertices[8])
}

func TestBfs(t *testing.T) {
	// graph
	g := NewUndirectedGraph()

	// ErrNotAllVertExist test
	err := g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	// create vertices
	for i := 0; i < 8; i++ {
		err := g.AddVertex(i + 1)
		if err != nil {
			fmt.Println(err)
		}
	}

	// ErrVertExist test
	err = g.AddVertex(5)
	if err != nil {
		fmt.Println(err)
	}

	// begin create edges
	err = g.AddEdge(1, 2)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 3)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(2, 4)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 1)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 5)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 6)
	if err != nil {
		fmt.Println(err)
	}
	err = g.AddEdge(6, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(5, 7)
	if err != nil {
		fmt.Println(err)
	}

	err = g.AddEdge(4, 7)
	if err != nil {
		fmt.Println(err)
	}
	// end of edges

	path, err := ShortestBFS(1, 6, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	fmt.Println("")

	path, err = ShortestBFS(6, 3, g)
	if err != nil {
		fmt.Println(err)
	} else {
		for path.tail != nil {
			fmt.Println(path.dequeue().Key)
		}
	}

	// fmt.Println(g.Vertices[1])
	// fmt.Println(g.Vertices[2])
	// fmt.Println(g.Vertices[3])
	// fmt.Println(g.Vertices[4])
	// fmt.Println(g.Vertices[5])
	// fmt.Println(g.Vertices[6])
	// fmt.Println(g.Vertices[7])
	// fmt.Println(g.Vertices[8])
}
