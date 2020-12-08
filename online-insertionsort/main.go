package main

import (
	"fmt"
)

// builds a sorted array as it receives values from c
// returns the current array when caller reads from quit and ends routine
func InsSort(c chan int, quit chan []int) {
	a := []int{}
	var j int
	for {
		select {
		case x := <-c:
			for j = 0; j < len(a) && a[j] < x; j++ {
			}
			a = append(a, 0)
			copy(a[j+1:], a[j:len(a)-1])
			a[j] = x
		case quit <- a:
			return
		}

	}
}

func main() {
	c, q := make(chan int), make(chan []int)
	var x int
	go InsSort(c, q)
	for {
		if _, err := fmt.Scanf("%d", &x); err == nil {
			c <- x
		} else {
			break
		}
	}
	//	a := []int{6, 3, 12, 45, 92, 54, 21, -5, 98}
	//	c, q := make(chan int), make(chan []int)
	//	go InsSort(c, q)
	//	for _, x := range a {
	//		c <- x
	//	}
	b := <-q
	fmt.Println(b)
}
