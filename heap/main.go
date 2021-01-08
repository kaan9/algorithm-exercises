package main

import "fmt"

type Comparer interface {
	Compare(c interface{}) int
}

type MaxHeap []int

func parent(i int) int { return (i-1)/2 }

func left(i int) int { return 2*i+1 }

func right(i int) int { return 2*i+2 }

func NewMaxHeap(a []int) MaxHeap {
	var h MaxHeap
	h = make([]int, len(a))
	copy(h, a)
	for i := len(h)/2+1; i >= 0; i-- {
		h.maxheapify(i)
	}
	return h
}

// Assumes that left(i) and right(i) are max-heaps but h[i] is possibly out-of-place
// transforms the subtree on h[i] to a maxheap based on this assumption
func (h MaxHeap) maxheapify(i int) {
	var largest int
	l, r := left(i), right(i)
	if l < len(h) && h[l] > h[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(h) && h[r] > h[largest] {
		largest = r
	}
	if largest != i {
		h[i], h[largest] = h[largest], h[i]
		h.maxheapify(largest)
	}
}

//consumes the heap
func HeapSort(h MaxHeap) []int {
	n := len(h)
	for i := n-1; i > 0; i-- {
		h[0], h[i] = h[i], h[0]
		h[:i].maxheapify(0)
	}
	return h
}

func (h MaxHeap) Insert(key int) {
	h = append(h, 0)
	h[len(h)-1] = h[parent(len(h)-1)-1]
	h.IncreaseKey(len(h)-1, key)
}

func (h MaxHeap) ExtractMax() (x int, ok bool) {
	if len(h) < 1 {
		return 0, false
	}
	x = h[0]
	h[0], h = h[len(h)-1], h[:len(h)-1]
	if len(h) > 0 {
		h.maxheapify(0)
	}
	return x, true
}

// increase the key of h[i] to key, does nothing if new key is smaller
func (h MaxHeap) IncreaseKey(i int, key int) {
	if key < h[i] {
		return
	}
	h[i] = key
	for ; i > 0 && h[parent(i)] < h[i]; i = parent(i) {
		h[i], h[parent(i)] = h[parent(i)], h[i]
	}
}

func (h MaxHeap) HeapMaximum() (x int, ok bool) {
	if len(h) < 1 {
		return 0, false
	}
	return h[0], true
}


func main() {
	s := []int{9, 2, 1, 4, 2, 4, 0, 8, 23, 645, 1, 34, -12, 3, 120, 998, 34, 71, 23, 76, 35}
	h := NewMaxHeap(s)
	fmt.Println(HeapSort(h))
	for len(h) > 0 {
		fmt.Println(h.HeapMaximum())
		fmt.Println(h.ExtractMax()) //FIXME: There is an error here
	}
}
