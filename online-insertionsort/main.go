package main

import (
	"fmt"
)

//TODO: also implement using a skip list which reduces runtime to O(n log(n)), see wikipedia page
// alternatively, leaving gaps in the array (library sort/gapped insertion sort) reduces runtime to O(n*log(n))
// with high probability, but this interferes with the "online" aspect. Another function to "oversee" the InsSort would be necessary to maintain "online" behavior while these changes happen

// builds a sorted array as it receives values from c
// returns the current array when caller reads from quit and ends routine
// when a value is read from request, a copy of the array is sent through reply
// request, reply, and quit may be nil
func InsSort(c <-chan int, request <-chan int, reply chan<- []int, quit chan<- []int) {
	a := []int{}
	for {
		select {
		case x := <-c:
			var lo, hi, j int
			for lo, hi = 0, len(a); lo < hi-1; {
				if a[(lo+hi)/2] > x {
					hi = (lo+hi)/2
				} else {
					lo = (lo+hi)/2
				}
			}
			j = hi
			if len(a) > 0 && a[0] > x { //edge-case x smaller than all elements
				j = lo
			}
			a = append(a, 0)
			copy(a[j+1:], a[j:len(a)-1])
			a[j] = x
		case quit <- a:
			return
		case <-request:
			b := make([]int, len(a))
			copy(b, a)
			reply <- b
		}

	}
}

func main() {
	c, request, reply, q := make(chan int), make(chan int), make(chan []int), make(chan []int)
	var x int
	go InsSort(c, request, reply, q)
	for {
		if _, err := fmt.Scanf("%d\n", &x); err == nil {
			c <- x
			request <- 0
			fmt.Println(<-reply)
		} else {
			break
		}
	}
//	a := []int{6, 3, 12, 45, 92, 54, 21, -5, 98}
//	go InsSort(c,request, reply, q)
//	for _, x := range a {
//		c <- x
//		request <- 0
//		fmt.Println(<-reply)
//	}
	b := <-q
	fmt.Println(b)
}
