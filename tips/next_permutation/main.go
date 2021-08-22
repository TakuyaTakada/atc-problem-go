package main

import (
	"sort"
)

func nextPermutation(s sort.Interface) bool {
	n := s.Len()
	if n < 2 {
		return false
	}
	i := n - 2
	for ; !s.Less(i, i+1); i-- {
		if i == 0 {
			return false
		}
	}
	j := n - 1
	for !s.Less(i, j) {
		j--
	}
	s.Swap(i, j)
	for k, l := i+1, n-1; k < l; {
		s.Swap(k, l)
		k++
		l--
	}
	return true
}

type Chars []string
func (s Chars) Len() int           { return len(s) }
func (s Chars) Less(i, j int) bool { return s[i] < s[j] }
func (s Chars) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }