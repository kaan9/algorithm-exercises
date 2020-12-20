package main

import (
	"fmt"
	"math/rand"
	"time"
)

// mostly the same as quicksort, but without the concurrency

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

// returns the k'th 0-indexed smallest element (if s was sorted, this would be s[k])
// get_p_index is function that takes in a slice and returns the index of the pivot
// if left nil, pivot default to mid-value
func qselect(s []int, k int, get_p_index func([]int) int) int {
	if len(s) <= 0 {
		return -1
	} else if len(s) == 1 {
		return s[0]
	}
	var p_index int
	if get_p_index == nil {
		p_index = (len(s) - 1)/2
	} else {
		p_index = get_p_index(s)
	}

	l := partition(s, p_index)
	if k == l {
		return s[l]
	} else if k < l {
		return qselect(s[:l], k, get_p_index)
	} else {
		return qselect(s[l+1:], k-l-1, get_p_index)
	}

}

func main() {
	s := []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 35}
	for i := 0; i < len(s); i++ {
		fmt.Println(qselect(s, i, nil))
	}

	// randomized quickselect:
	rand.Seed(time.Now().UnixNano())
	s = []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 35}
	for i := 0; i < len(s); i++ {
		fmt.Println(qselect(s, i, func(s []int) int { return rand.Intn(len(s)) }))
	}

}
