package main

import (
	"fmt"
)

// Given a set P of n points in the plane, find the closest pair of points (p, q) using Euclidean distance

type Point struct {
	x, y float32
}

// the qsort function is extracted and adapted from "parallel-quicksort"

// partitions around s[p_index], returns new index of pivot
func partition(s []Point, p_index int) int {
	pivot := s[p_index].x
	s[p_index], s[len(s)-1] = s[len(s)-1], s[p_index]
	var l, r Point
	for l, r = 0, len(s)-2; l <= r; {
		if s[l].x <= pivot {
			l++
		} else {
			s[l], s[r] = s[r], s[l]
			r--
		}
	}
	s[len(s)-1], s[l] = s[l], s[len(s)-1]
	return l
}

// get_p_index is function that takes in a slice and returns the index of the pivot
// if left nil, pivot default to mid-value
func qsort(s []Point, quit chan int, get_p_index func([]int) int) {
	if len(s) <= 1 {
		quit <- 0
		return
	}
	var p_index int
	if get_p_index == nil {
		p_index = (len(s) - 1)/2
	} else {
		p_index = get_p_index(s)
	}

	l := partition(s, p_index)

	q := make(chan int)
	go qsort(s[:l], q, get_p_index)
	go qsort(s[l+1:], q, get_p_index)
	<-q
	<-q
	quit <- 0
}



// returns closest pair and the distance between them
// if len(P) < 2 returns -1 for dist
func ClosestPair(P []Point) (p, q Point, dist float32) {
	if len(P) < 2 {
		return Point{0, 0}, Point{0, 0}, -1
	} else if len(P) == 2 {
		return P[0], P[1], nil
	}
	q := make(chan int)
	qsort(s, q, nil)
	<-q
	(pl, ql, dl), (pr, qr, dr) := ClosestPair(P[:len(P)/2]), ClosestPair(P[len(P)/2:])
	if dl < 0 {
		return pr, qr, dr
	} else if dr < 0 {
		return pl, ql, dl
	}
	var pmin, qmin Point, dmin float32
	if dl < dr {
		pmin, qmin, dmin = pl, ql, dl
	} else {
		pmin, qmin, dmin = pr, qr, dr
	}
	line := (P[len(P)/2-1] + P[len(P)/2])/2

	//FIXME: complete this algorithm
}
