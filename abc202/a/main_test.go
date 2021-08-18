package main

import (
	"fmt"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		abc []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{[]int{1, 4, 3}},
			13,
		},
		{
			args{[]int{5, 6, 4}},
			6,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			if got := run(tt.args.abc); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
