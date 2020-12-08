package main

import (
	"fmt"
)

// adjacency list representation of a graph data structure
// vertices are denoted with positive integers (not zero!)
// the actual structure uses hash maps for efficiency and ease of use
// V is the set of keys E has
type Graph struct {
	E map[uint]map[uint]bool
}

// runs BFS starting from s and writes each discovered vertex to c
// terminates when quit reads a value or when c no longer returns values`
// assumes graph doesn't change while BFS executes
func (g *Graph) BFS(s uint, c, quit chan uint) {
	defer close(c)

	discovered, parent := make(map[uint]bool), make(map[uint]uint)
	queue := make([]uint, 0)
	queue = append(queue, s)
	discovered[s] = true

	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		for u, b := range g.E[v] {
			if b && !discovered[u] {
				select {
				case c <- u:
					discovered[u] = true
					queue = append(queue, u)
					parent[u] = v
				case <-quit:
					return
				}
			}
		}
	}
}

func (g *Graph) BFS_layer(s uint, l chan []uint, quit chan uint) {
	defer close(l)

	discovered := make(map[uint]bool)
	discovered[s] = true

	L_prev := []uint{s}

	for len(L_prev) > 0 {
		L_new := make([]uint, 0)
		for _, v := range L_prev {
			for u, b := range g.E[v] {
				if b && !discovered[u] {
					discovered[u] = true
					L_new = append(L_new, u)
				}
			}
		}
		select {
		case l <- L_prev:
			L_prev = L_new
		case <-quit:
			return
		}
	}
}

func Check_connected(g *Graph, s, t uint) bool {
	c, quit := make(chan uint), make(chan uint)
	go g.BFS(s, c, quit)
	for v := range c {
		fmt.Println(v)
		if v == t {
			return true
		}
	}
	return false
}

func main() {
	g := new(Graph)
	g.E = map[uint]map[uint]bool{
		1: map[uint]bool{
			2: true,
			3: true,
		},
		2: map[uint]bool{
			1: true,
			3: true,
		},
		3: map[uint]bool{
			1: true,
			2: true,
			5: true,
		},
		5: map[uint]bool{
			3:   true,
			100: true,
		},
		100: map[uint]bool{
			5:   true,
			230: true,
		},
		230: map[uint]bool{
			100: true,
		},
	}
	if !Check_connected(g, 1, 230) {
		fmt.Println("failed!")
	}

	l, q := make(chan []uint), make(chan uint)
	go g.BFS_layer(1, l, q)
	for layer := range l {
		fmt.Println(layer)
	}
}
