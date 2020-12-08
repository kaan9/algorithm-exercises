package main

import "fmt"

/*
 * Union-find data structure with path compression
 * Make-Set O(1), Find O(log*(n)), Union O(log*(n))
 */

// contains several sets with no common elements
type DisjointSet struct {
	parent map[uint]uint
	rank   map[uint]uint
}

// make a new set in d with only element x
// x must not already be in d
func (d *DisjointSet) MakeSet(x uint) {
	d.parent[x] = x
	d.rank[x] = 0
}

// combines the two sets that contain x and y in d
// assumes x and y exist in d
// does nothing if x and y are in the same set
func (d *DisjointSet) Union(x, y uint) {
	r_x, r_y := d.Find(x), d.Find(y)
	if r_x == r_y {
		return
	}
	if d.rank[r_x] > d.rank[r_y] {
		d.parent[r_y] = r_x
	} else {
		d.parent[r_x] = r_y
		if d.rank[r_x] == d.rank[r_y] {
			d.rank[r_y]++
		}
	}
}

// give the representative element of the set containing x
func (d *DisjointSet) Find(x uint) uint {
	y := x
	for d.parent[y] != y {
		y = d.parent[y]
	}
	for d.parent[x] != x {
		x, d.parent[x] = d.parent[x], y
	}
	return y
}

func main() {
	d := DisjointSet{
		parent: map[uint]uint{},
		rank:   map[uint]uint{},
	}
	var i uint
	for i = 1; i <= 7; i++ {
		d.MakeSet(i)
	}
	fmt.Println(d)

	d.Union(1, 4)
	d.Union(2, 5)
	d.Union(3, 6)

	fmt.Println(d)

	d.Union(3, 7)
	d.Union(5, 1)

	fmt.Println(d)

	d.Union(2, 7)

	fmt.Println(d)

	d.Find(7)
	d.Find(3)

	fmt.Println(d)

}
