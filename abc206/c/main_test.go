package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{3, []int{1,7,1}},
			2,
		},
		{
			args{10,[]int{1,10,100,1000,10000,100000,1000000,10000000,100000000,1000000000}},
			45,
		},
		{
			args{20,[]int{7,8,1,1,4,9,9,6,8,2,4,1,1,9,5,5,5,3,6,4}},
			173,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
