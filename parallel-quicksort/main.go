package main

import (
	"fmt"
	"math/rand"
	"time"
)

// partitions around s[p_index], returns new index of pivot
func partition(s []int, p_index int) int {
	pivot := s[p_index]
	s[p_index], s[len(s)-1] = s[len(s)-1], s[p_index]
	var l, r int
	for l, r = 0, len(s)-2; l <= r; {
		if s[l] <= pivot {
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
func qsort(s []int, quit chan int, get_p_index func([]int) int) {
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

func main() {
	s := []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 35}
	q := make(chan int)
	go qsort(s, q, nil)
	<-q
	fmt.Println(s)

	// random quicksort:
	s = []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 35}
	q = make(chan int)
	rand.Seed(time.Now().UnixNano())
	go qsort(s, q, func(s []int) int { return rand.Intn(len(s)) })
	<-q
	fmt.Println(s)

}
