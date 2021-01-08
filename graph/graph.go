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
// when c is done returning values or quit reads a value, the program terminates 
// if parent isn't nil, BFS stores the parent hierarchy in it 
// assumes graph doesn't change while BFS executes
func (g *Graph) BFS(s uint, c chan uint, quit chan uint, parent map[uint]uint) {
	defer close(c)

	discovered := make(map[uint]bool)
	if parent == nil {
		parent = make(map[uint]uint)
	}
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

// runs BFS from s and sends each layer to l
// terminates when BFS is complete or quit is read from
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

// performs DFS on g and sends encountered nodes through c
// terminates when all vertices are traversed and sent through c or when a value is read from quit
// stored the discovery, finishing times and parent hierarchy in d, f, parent
// c, q, d, f, parent may be nil in which case no value is stored/sent/read
func (g *Graph) DFS(c, quit chan uint, d, f map[uint]int, parent map[uint]uint)  {
	const (
		White = 0
		Gray = 1
		Black = 2
	)
	color := make(map[uint]byte) // nodes start as White , Gray on discovery, Black on completion
	if d == nil {
		d = make(map[uint]int)
	}
	if f == nil {
		f = make(map[uint]int)
	}
	if parent == nil {
		parent = make(map[uint]uint)
	}
	if c == nil {
		c = make(chan uint)
		go func() { for _ = range c {} }()
	}
	defer close(c)
	time := 0

	var dfs_visit func(u uint)
	dfs_visit = func(u uint) {
		select {
		case c <- u:
			time++
			d[u] = time
			color[u] = Gray
		case <-quit:
			return
		}

		for v := range g.E[u] {
			if color[v] == White {
				parent[v] = u
				dfs_visit(v)
			}
		}
		time++
		f[u] = time
		color[u] = Black
	}

	for u, _ := range g.E {
		if color[u] == White {
			dfs_visit(u)
		}
	}
}

func (g *Graph) Connected(s, t uint) bool {
	c, quit := make(chan uint), make(chan uint)
	parent := make(map[uint]uint)
	go g.BFS(s, c, quit, parent)
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
	if !g.Connected(1, 230) {
		fmt.Println("failed!")
	}

//	l, q := make(chan []uint), make(chan uint)
//	go g.BFS_layer(1, l, q)
//	for layer := range l {
//		fmt.Println(layer)
//	}

	c, q := make(chan uint), make(chan uint)
	go g.DFS(c, q, nil, nil, nil)
	for v := range c {
		println(v)
	}

	//TODO: Implement other functions using BFS/DFS
}
