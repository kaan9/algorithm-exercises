package main

import "fmt"

func merge(s []int, quit chan int) {
	if len(s) <= 1 {
		quit <- 0
		return
	} else if len(s) == 2 {
		if s[0] > s[1] {
			s[0], s[1] = s[1], s[0]
		}
		quit <- 0
		return
	}
	l, r := make([]int, len(s)/2), make([]int, len(s)-len(s)/2)
	copy(l, s[:len(s)/2])
	copy(r, s[len(s)/2:])
	q := make(chan int)
	go merge(l, q)
	go merge(r, q)
	<-q
	<-q
	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			s[i+j] = l[i]
			i++
		} else {
			s[i+j] = r[j]
			j++
		}
	}
	if i < len(l) {
		copy(s[i+j:], l[i:])
	}
	if j < len(r) {
		copy(s[i+j:], r[j:])
	}
	quit <- 0
}

func main() {
	a := []int{6, 3, 12, 45, 92, 54, 21, -5, 98}
	q := make(chan int)
	go merge(a, q)
	<-q
	fmt.Println(a)
}
