package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		n, k int
		ab IntHeap
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{2, 3, []Friend{
				{2, 1},
				{5, 10},
			}},
			4,
		},
		{
			args{5, 1000000000, []Friend{
				{1, 1000000000},
				{2, 1000000000},
				{3, 1000000000},
				{4, 1000000000},
				{5, 1000000000},
			}},
			6000000000,
		},
		{
			args{3, 2, []Friend{
				{5, 5},
				{2, 1},
				{2, 2},
			}},
			10,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.n, tt.args.k, tt.args.ab); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
