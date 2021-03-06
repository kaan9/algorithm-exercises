package main

import "fmt"

func issorted(s []int) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}
	return true
}

func bubblesort(s []int) {
	n := len(s)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

// see Knuth volume 3
func bubblesort_bound(s []int) {
	bound := len(s)
	for bound > 0 {
		max := bound
		bound = 0
		for j := 0; j < max-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				bound = j + 1
			}
		}
	}
}

func cocktailshake(s []int) {
	bottom, top := 0, len(s)
	for bottom < top {
		min, max := bottom, top
		bottom, top = top, 0
		for j := min; j < max-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				top = j + 1
			}
		}
		for j := top - 2; j >= 0; j-- {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				bottom = 0
			}
		}

	}
}

func combsort(s []int) {
	const k = 1.3 // shrink factor, empirically determined
	n := len(s)
	for gap := int(float32(n) / k); gap >= 1; gap = int(float32(gap) / k) {
		for i, j := 0, gap; j < n; i, j = i+1, j+1 {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}

func main() {
	mylist := []int{46, 28, 73, 42, 38, 49, 32, 27, 8, 65, 58, 72, 9, 93, 13, 17, 44, 22, 58, 6, 17, 58, 33, 82, 6, 33, 18, 70, 55, 83, 71, 2, 61, 3, 25, 88, 1, 92, 66, 62, 96, 97, 68, 35, 99, 23, 54, 71, 99, 12, 18, 27, 14, 47, 2, 24, 74, 51, 68, 26, 45, 76, 60, 52, 63, 93, 69, 10, 100, 26, 26, 14, 53, 41, 66, 78, 75, 19, 36, 3, 78, 4, 40, 80, 25, 14, 6, 37, 20, 57, 59, 79, 58, 17, 89, 21, 85, 18, 83, 50, 70, 67, 54, 83, 92, 63, 38, 57, 30, 31, 6, 88, 67, 85, 100, 46, 84, 53, 34, 66, 32, 26, 68, 6, 32, 89, 87, 72, 95, 17, 23, 36, 59, 53, 50, 48, 44, 13, 67, 79, 55, 87, 82, 12, 69, 55, 23, 11, 7, 12, 70, 90, 42, 44, 46, 93, 92, 87, 84, 17, 46, 79, 13, 9, 51, 79, 43, 68, 64, 32, 84, 8, 27, 10, 87, 100, 67, 11, 67, 93, 3, 54, 52, 47, 45, 23, 75, 57, 29, 8, 70, 15, 3, 100, 24, 5, 83, 75, 2, 82, 11, 40, 54, 31, 7, 94, 68, 81, 86, 98, 12, 40, 71, 88, 67, 94, 31, 53, 36, 98, 77, 47, 24, 35, 11, 68, 1}

	//	bubblesort(mylist)
	//	bubblesort_bound(mylist)
	//	cocktailshake(mylist)
	combsort(mylist)
	fmt.Println(mylist)
	fmt.Println(issorted(mylist))
}
