package main

import "fmt"

func qsort(s []int, quit chan int) {
	if len(s) <= 1 {
		quit <- 0
		return
	}
	pivot := s[(len(s) - 1) / 2]
	s[(len(s) - 1) / 2], s[len(s) - 1] = s[len(s) - 1], s[(len(s) - 1) / 2]
	var l, r int
	for l, r = 0, len(s) - 2; l <= r; {
		if s[l] <= pivot {
			l++
		} else {
			s[l], s[r] = s[r], s[l]
			r--
		}
	}
	s[len(s) - 1], s[l] = s[l], s[len(s) - 1]
	q := make(chan int)
	go qsort(s[:l], q)
	go qsort(s[l + 1:], q)
	<-q
	<-q
	quit <- 0
}

func main() {
	s := []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 34}
//	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	q := make(chan int)
	go qsort(s, q)
	<-q
	fmt.Println(s)
}
