package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n int
		a, b, c []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{3, []int{1,2,2}, []int{3,1,2}, []int{2,3,2}},
			4,
		},
		{
			args{4, []int{1,1,1,1}, []int{1,1,1,1}, []int{1,2,3,4}},
			16,
		},
		{
			args{3, []int{2,3,3}, []int{1,3,3}, []int{1,1,1}},
			0,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
