package main

import "fmt"

// preferences of n = len(M) = len(W) men and women where M[i] is the preference list of man i
// that is M[i][0] is the top preference, M[i][1] is the second choice, etc, same for W
// returns a stable matching where j = matching[i] corresponds to the pair (i, j)
func GaleShapley(M [][]int, W [][]int) []int{
	// use 0 to indicate an unmatched, n to indicate a match to the 0th man/woman
	n := len(M)
	matching, inv_matching := make([]int, n), make([]int, n)
	proposed := make([][]bool, n, n) // (i, j) == has man i proposed to woman j
	free_queue := make([]int, n) //queue of men who aren't engaged
	for i := 0; i < n; i++ {
		free_queue[i] = i
	}
	for len(free_queue) > 0 {
		m := free_queue[0]
		free_queue = free_queue[1:]
		w := -1
		for _, w_high := range M[m] {
			if !proposed[m][w_high] {
				w = w_high
				proposed[m][w] = true
				break
			}
		}
		if w == -1 {
			continue
		}
		if inv_matching[w] == 0 {
			if w == 0 {
				matching[m] = n
			} else {
				matching[m] = w
			}
			if m == 0 {
				inv_matching[w] = n
			} else {
				inv_matching[w] = m
			}
		} else {
			m_other := inv_matching[w]
			var pref int
			for pref = range W[w] {
				if pref == m || pref == m_other {
					break
				}
			}
			if pref == m {
				matching[m_other] = 0
				inv_matching[w] = m
				matching[m] = w
				free_queue = append(free_queue, m_other)
			} else {
				free_queue = append(free_queue, m)
			}
		}
	}
	return matching
}

//TODO: write some tests and verify
