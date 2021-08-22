package main

import "math"

func findIndex(n []int, v int) int {
	ans := -1
	for i, val := range n {
		if val == v {
			ans = i
			break
		}
	}
	return ans
}

func min(n []int) int {
	ans := math.MaxInt64
	for _, v := range n {
		if ans > v {
			ans = v
		}
	}
	return ans
}
