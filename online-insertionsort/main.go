package main

import (
	"fmt"
)

//TODO: also implement using a skip list which reduces runtime to O(n log(n)), see wikipedia page
// alternatively, leaving gaps in the array (library sort/gapped insertion sort) reduces runtime to O(n*log(n))
// with high probability, but this interferes with the "online" aspect. Another function to "oversee" the InsSort would be necessary to maintain "online" behavior while these changes happen

// builds a sorted array as it receives values from c
// returns the current array when caller reads from quit and ends routine
func InsSort(c chan int, quit chan []int) {
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
