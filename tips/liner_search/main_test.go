package main

import (
	"fmt"
	"testing"
)

func Test_findIndex(t *testing.T) {
	type args struct {
		n []int
		v int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				 []int{1,2,3,4,5,6},
				4,
			},
			3,
		},
		{
			args{
				[]int{100,19,5,26,301,77},
				77,
			},
			5,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			if got := findIndex(tt.args.n, tt.args.v); got != tt.want {
				t.Errorf("findIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}